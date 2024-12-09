package api

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/products", CreateProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", GetFilteredProductsHandler).Methods("GET")
	r.HandleFunc("/products", GetFilteredProductsHandler).Methods("GET")
	return r
}
