package market

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(app *Application) http.Handler {

	router := mux.NewRouter()

	router.HandleFunc("/v1/product", app.Create).Methods(http.MethodPost)
	router.HandleFunc("/v1/product", app.ReadAll).Methods(http.MethodGet)
	router.HandleFunc("/v1/product/{id}", app.Read).Methods(http.MethodGet)
	router.HandleFunc("/v1/product/{id}", app.Update).Methods(http.MethodPut)
	router.HandleFunc("/v1/product/{id}", app.Delete).Methods(http.MethodDelete)

	return router
}
