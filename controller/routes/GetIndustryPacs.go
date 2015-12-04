package routes

import (
	"encoding/json"
	"finalproject/globals"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetIndustryPacs ...
//   returns PACs from a specific industry that supported a politician
func GetIndustryPacs() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		memberID := vars["memberID"]
		industry := vars["industry"]
		db := globals.DB
		// get bills congress person has voted on
		query := fmt.Sprintf("SELECT DISTINCT p.name, pd.amount FROM PAC p, PAC_Donations pd WHERE p.industry = '%s' AND pd.congress_id = '%s' AND pd.pac_id = p.pacID", industry, memberID)
		rows, err := db.Query(query)

		if err != nil {
			log.Panic(err)
		}

		pacs := []interface{}{}
		var pac string
		var amount int

		for rows.Next() {
			temp := make(map[string]string)
			err := rows.Scan(&pac, &amount)
			if err != nil {
				log.Fatal(err)
			}
			temp["pac"] = pac
			temp["amount"] = strconv.Itoa(amount)
			pacs = append(pacs, temp)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(pacs)
	}
}
