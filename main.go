package main

import (
	"FinalProject/controller/routes"
	"FinalProject/globals"
	"FinalProject/utility"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	globals.DB = utility.LoadDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/", routes.GetBill())
	r.HandleFunc("/congress", routes.GetCongress())
	r.HandleFunc("/member/{memberID}", routes.GetCongressPerson())
	r.HandleFunc("/contributor/{contributor}", routes.GetContributor())
	r.HandleFunc("/bill", routes.GetBill())
	http.Handle("/", r)
	log.Printf("so far so good")
	log.Fatal(http.ListenAndServe(":8080", r))
}
