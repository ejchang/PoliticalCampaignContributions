package main

import (
	"FinalProject/controller/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Test())
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
