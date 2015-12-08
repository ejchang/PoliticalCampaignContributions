package routes

import (
	"encoding/json"
	"finalproject/globals"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GetCongressPerson ...
//   returns representative based on state and party/*
func GetCongressPerson() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		memberID := vars["memberID"]
		out := map[string]interface{}{}

		db := globals.DB
		rows, err := db.Query("SELECT name, party, state, chamber FROM congressmembers WHERE congress_id = $1", memberID)
		if err != nil {
			log.Panic(err)
		}
		var repname, party, state, chamber string
		for rows.Next() {
			err = rows.Scan(&repname, &party, &state, &chamber)
			if err != nil {
				log.Fatal(err)
			}
			out["name"] = repname
			out["party"] = party
			out["state"] = state
			out["chamber"] = chamber
		}

		lastName := strings.Split(repname, " ")
		// get bills congress person has voted on
		rows, err = db.Query("SELECT DISTINCT b.bill_id, b.name, v.vote, b.description FROM Bill b, Voted v WHERE v.state = $1 AND v.name = $2 AND v.bill_id = b.bill_id", state, lastName[1])

		if err != nil {
			log.Panic(err)
		}

		bills := []interface{}{}
		var billID, billName, vote, description string

		for rows.Next() {
			temp := make(map[string]string)
			err := rows.Scan(&billID, &billName, &vote, &description)
			if err != nil {
				log.Fatal(err)
			}
			temp["billid"] = billID
			temp["vote"] = vote
			temp["description"] = description
			temp["billname"] = billName
			bills = append(bills, temp)
		}
		out["bills"] = bills

		// get top industries and how much they donated to this member
		rows, err = db.Query("SELECT pd.industry, sum(pd.amount) FROM CongressMembers cg, Pac_Donations pd WHERE cg.congress_id = $1 and pd.congress_id = cg.congress_id GROUP BY cg.congress_id, pd.industry ORDER BY sum(pd.amount) DESC LIMIT 20", memberID)

		if err != nil {
			log.Panic(err)
		}
		var industry, name string
		var amount int
		donors := []interface{}{}

		for rows.Next() {
			err := rows.Scan(&industry, &amount)
			if err != nil {
				log.Fatal(err)
			}
			temp := make(map[string]interface{})
			temp["industry"] = industry
			temp["amount"] = amount
			pacrows, err := db.Query("SELECT p.name, pd.amount FROM PAC p, PAC_Donations pd WHERE p.industry = $1 AND pd.congress_id = $2 AND pd.pac_id = p.pacID", industry, memberID)
			if err != nil {
				log.Panic(err)
			}
			paclist := make(map[string]int)
			for pacrows.Next() {
				err = pacrows.Scan(&name, &amount)
				if err != nil {
					log.Fatal(err)
				}
				if _, ok := paclist[name]; !ok {
					paclist[name] = 0
				}
				paclist[name] = paclist[name] + amount
			}
			temp["pacs"] = paclist
			donors = append(donors, temp)
		}
		out["donors"] = donors

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(out)
	}
}
