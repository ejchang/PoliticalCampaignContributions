package routes

import (
	"FinalProject/globals"
	"fmt"
	"log"
	"net/http"
)

// GetContributor ...
//   returns information on bill/*
func GetContributor() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db := globals.DB
		rows, err := db.Query("SELECT DISTINCT bill_id, description FROM donors d, bill b WHERE d.industry = b.industry")
		if err != nil {
			log.Fatal(err)
		}
		var billID int
		var description string
		for rows.Next() {
			err := rows.Scan(&billID, &description)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(billID, description)
		}
	}
}
