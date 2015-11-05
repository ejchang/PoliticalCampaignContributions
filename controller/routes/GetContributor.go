package routes

import (
	"FinalProject/globals"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetContributor ...
//   returns information on bill that a donor is interested in
func GetContributor() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		contributor := vars["contributor"]
		db := globals.DB
		query := "SELECT DISTINCT bill_id, description, name FROM donors d, bill b WHERE d.industry = b.industry and d.id=" + contributor
		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		out := map[string]interface{}{}
		donorInfo := []string{}
		billInfo := []map[string]string{}
		var billID int
		var description, name string
		for rows.Next() {
			err := rows.Scan(&billID, &description, &name)
			if err != nil {
				log.Fatal(err)
			}
			temp := make(map[string]string)
			temp["id"] = strconv.Itoa(billID)
			temp["description"] = description
			billInfo = append(billInfo, temp)
		}
		donorInfo = append(donorInfo, name)
		out["donor"] = donorInfo
		out["bills"] = billInfo
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(out)
	}
}
