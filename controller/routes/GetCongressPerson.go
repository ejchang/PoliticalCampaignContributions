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
		query := fmt.Sprintf("SELECT DISTINCT b.bill_id, v.vote, b.description FROM CongressMembers cg, Bill b, Voted v, Donors d WHERE v.congress_id = cg.congress_id AND v.bill_id = b.bill_id and v.congress_id = cg.congress_id and cg.congress_id = '%s'", memberID)
		rows, err := db.Query(query)

		if err != nil {
			log.Panic(err)
		}

		out := map[string]interface{}{}
		bills := []interface{}{}
		var billID, vote, description string

		for rows.Next() {
			temp := make(map[string]string)
			err := rows.Scan(&billID, &vote, &description)
			if err != nil {
				log.Fatal(err)
			}
			temp["billid"] = billID
			temp["vote"] = vote
			temp["description"] = description
			bills = append(bills, temp)
		}
		out["bills"] = bills

		// get donors and how much they donated to this member
		query = fmt.Sprintf("SELECT d.name, don.amount, d.id FROM CongressMembers cg, Donors d, Donations don WHERE cg.congress_id = '%s' and don.member_id = cg.congress_id and don.donor_id = d.id", memberID)
		rows, err = db.Query(query)

		if err != nil {
			log.Panic(err)
		}
		var name string
		var amount, id int
		donors := []interface{}{}

		for rows.Next() {
			err := rows.Scan(&name, &amount, &id)
			if err != nil {
				log.Fatal(err)
			}
			temp := make(map[string]string)
			temp["name"] = name
			temp["amount"] = strconv.Itoa(amount)
			temp["id"] = strconv.Itoa(id)
			donors = append(donors, temp)
		}
		out["donors"] = donors

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(out)
	}
}
