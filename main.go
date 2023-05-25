package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rajpatil7322/MicroService-in-Golang/handlers"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/transactions", handlers.GetProductHandler()).Methods("GET")
	router.Handle("/transactions", handlers.CreateProductHandler()).Methods("POST")
	server := http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	fmt.Println("Staring Product Catalog server on Port 9090")
	server.ListenAndServe()
}
