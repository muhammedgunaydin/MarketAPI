package market

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Repository interface {
	GetByID(ID uuid.UUID) (*Product, error)
	GetAll() ([]Product, error)
	Upsert(product *Product)
	Delete(ID uuid.UUID) error
}

type repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{DB: db}
}

func (db *repository) GetByID(ID uuid.UUID) (*Product, error) {
	sqlQuery := `SELECT * FROM product WHERE id =$1`

	row := db.DB.QueryRow(sqlQuery, ID)

	product := Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "product not found")
		}
		return nil, errors.Wrap(err, "failed to scan")
	}
	return &product, nil
}

func (db *repository) GetAll() ([]Product, error) {
	sqlQuery := `SELECT * FROM product`

	rows, err := db.DB.Query(sqlQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "products not found")
		}
		return nil, errors.Wrap(err, "failed to scan")
	}
	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan")
		}
		products = append(products, product)
	}
	return products, nil
}

func (db *repository) Upsert(product *Product) {
	sqlQuery := `INSERT INTO product (id, name, price) values ($1,$2,$3)
	ON CONFLICT(id)
	DO UPDATE SET name=$2 price=$3`

	args := []interface{}{product.ID, product.Name, product.Price}
	db.DB.Exec(sqlQuery, args...)
}

func (db *repository) Delete(ID uuid.UUID) error {
	sqlQuery := `DELETE FROM product WHERE ID=$1`

	_, err := db.DB.Exec(sqlQuery, ID)
	if err != nil {
		return errors.Wrap(err, "product not found")
	}
	return nil
}
