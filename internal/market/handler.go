package market

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Application struct {
	database Repository
}

func NewProduct(r Repository) *Application {
	return &Application{database: r}
}

func (a *Application) Create(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	product.ID = uuid.New()
	fmt.Println(product)
	a.database.Upsert(&product)
	w.WriteHeader(http.StatusCreated)
}

func (a *Application) ReadAll(w http.ResponseWriter, r *http.Request) {
	products, err := a.database.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(products)
	productsJSON, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(productsJSON)
	w.WriteHeader(http.StatusOK)
}

func (a *Application) Read(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	key, err := uuid.Parse(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	product, err := a.database.GetByID(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(productJSON)
	w.WriteHeader(http.StatusOK)
}

func (a *Application) Update(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	key, err := uuid.Parse(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}
	var product Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	product.ID = key
	a.database.Upsert(&product)
	w.WriteHeader(http.StatusCreated)
}

func (a *Application) Delete(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	key, err := uuid.Parse(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = a.database.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
