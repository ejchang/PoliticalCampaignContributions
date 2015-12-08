package routes

import (
	"encoding/json"
	"finalproject/globals"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetBill ...
//   returns information on bill/*
func GetBill() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		out := map[string]interface{}{}

		vars := mux.Vars(r)
		billID := vars["billID"]

		db := globals.DB

		var billid, billname, description, voteDate string
		rows, err := db.Query("SELECT * FROM BILL WHERE bill_id= $1", billID)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			err = rows.Scan(&billid, &billname, &description, &voteDate)
			if err != nil {
				log.Fatal(err)
			}
			out["billid"] = billid
			out["billname"] = billname
			out["description"] = description
			out["voteDate"] = voteDate
		}

		rows, err = db.Query("SELECT cm.name, vote, cm.state, cm.party, cm.congress_id FROM voted v, congressmembers cm WHERE bill_id = $1 AND cm.state = v.state AND cm.name LIKE '%' || v.name || '%'", billID)
		if err != nil {
			log.Fatal(err)
		}
		var name, vote, state, party, cmid string
		votes := []interface{}{}
		for rows.Next() {
			temp := make(map[string]string)
			err := rows.Scan(&name, &vote, &state, &party, &cmid)
			if err != nil {
				log.Fatal(err)
			}
			temp["member"] = name
			temp["vote"] = vote
			temp["state"] = state
			temp["party"] = party
			temp["id"] = cmid
			votes = append(votes, temp)
		}
		out["votes"] = votes
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(out)
	}
}
