package routes

import (
	"encoding/json"
	"finalproject/globals"
	"log"
	"net/http"
)

// GetCongress ...
//   returns information on bill/*
func GetCongress() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db := globals.DB
		out := []interface{}{}
		rows, err := db.Query("SELECT name, chamber, congress_id FROM congressmembers")
		if err != nil {
			log.Fatal(err)
		}

		var name, chamber, congressID string

		for rows.Next() {
			err := rows.Scan(&name, &chamber, &congressID)
			if err != nil {
				log.Fatal(err)
			}
			temp := make(map[string]string)
			temp["name"] = name
			temp["chamber"] = chamber
			temp["id"] = congressID
			out = append(out, temp)
		}

		// log.Println(string(formatted))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(out)
	}
}
