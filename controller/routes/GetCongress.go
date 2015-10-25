package routes

import (
	"FinalProject/globals"
	"encoding/json"
	"log"
	"net/http"
)

// GetCongress ...
//   returns information on bill/*
func GetCongress() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db := globals.DB
		out := make(map[string]map[string]string)
		rows, err := db.Query("SELECT name, chamber, party FROM congressmember WHERE state = 'AZ'")
		if err != nil {
			log.Fatal(err)
		}
		var name, chamber, party string

		for rows.Next() {
			err := rows.Scan(&name, &chamber, &party)
			if err != nil {
				log.Fatal(err)
			}
			temp := make(map[string]string)
			temp["chamber"] = chamber
			temp["party"] = party
			out[name] = temp
		}

		formatted, err := json.Marshal(out)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(formatted))
	}
}
