package routes

import (
	"FinalProject/globals"
	"log"
	"net/http"
)

// GetBill ...
//   returns information on bill/*
func GetBill() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// var out interface{}
		db := globals.DB

		rows, err := db.Query("SELECT cm.name, vote FROM voted v, congressmember cm WHERE bill_id = 1003")
		if err != nil {
			log.Fatal(err)
		}
		var name, vote string

		for rows.Next() {
			err := rows.Scan(&name, &vote)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(name, vote)
		}
	}
}
