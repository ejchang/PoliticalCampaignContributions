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

// GetCongressPerson ...
//   returns representative based on state and party/*
func GetCongressPerson() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		memberID := vars["memberID"]
		db := globals.DB
		// get bills congress person has voted on
		// query := fmt.Sprintf("SELECT DISTINCT b.bill_id, v.vote, b.description FROM CongressMembers cg, Bill b, Voted v WHERE v.congress_id = cg.congress_id AND v.bill_id = b.bill_id and v.congress_id = cg.congress_id and cg.congress_id = '%s'", memberID)
		// rows, err := db.Query(query)
		//
		// if err != nil {
		// 	log.Panic(err)
		// }

		out := map[string]interface{}{}
		// bills := []interface{}{}
		// var billID, vote, description string
		//
		// for rows.Next() {
		// 	temp := make(map[string]string)
		// 	err := rows.Scan(&billID, &vote, &description)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	temp["billid"] = billID
		// 	temp["vote"] = vote
		// 	temp["description"] = description
		// 	bills = append(bills, temp)
		// }
		// out["bills"] = bills

		// get top industries and how much they donated to this member
		query := fmt.Sprintf("SELECT pd.industry, sum(pd.amount) FROM CongressMembers cg, Pac_Donations pd WHERE cg.congress_id = '%s' and pd.congress_id = cg.congress_id GROUP BY cg.congress_id, pd.industry ORDER BY sum(pd.amount) DESC LIMIT 20", memberID)
		rows, err := db.Query(query)

		if err != nil {
			log.Panic(err)
		}
		var industry string
		var amount int
		donors := []interface{}{}

		for rows.Next() {
			err := rows.Scan(&industry, &amount)
			if err != nil {
				log.Fatal(err)
			}
			temp := make(map[string]string)
			temp["industry"] = industry
			temp["amount"] = strconv.Itoa(amount)
			donors = append(donors, temp)
		}
		out["donors"] = donors

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(out)
	}
}
