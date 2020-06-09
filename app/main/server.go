package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	log.Println("MATRIX OPERATIONS API")
	log.Println("Starting Server")

	api := router.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "v1")
	})

	api.HandleFunc("/echo", echo).Methods(http.MethodPost)
	api.HandleFunc("/invert", invert).Methods(http.MethodPost)
	api.HandleFunc("/flatten", flatten).Methods(http.MethodPost)
	api.HandleFunc("/sum", sum).Methods(http.MethodPost)
	api.HandleFunc("/multiply", multiply).Methods(http.MethodPost)

	log.Println("Server Started!!!")
	log.Fatal(http.ListenAndServe(":8082", router))
}
