// crud
// psql connection
// docker
// postman test

package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/muhammedgunaydin/MarketAPI/internal/market"
)

func main() {
	db, err := OpenConn("postgres://postgres:postgres@localhost/product?sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := market.NewRepository(db)
	svc := market.NewProduct(repo)
	router := market.NewRouter(svc)

	http.ListenAndServe(":8080", router)

}

func OpenConn(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "error opening postgres connection")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "ping failed")
	}
	return db, nil
}
