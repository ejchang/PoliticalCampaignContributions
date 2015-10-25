package routes

import (
	"FinalProject/globals"
	"fmt"
	"log"
	"net/http"
)

// GetCongressPerson ...
//   returns representative based on state and party/*
func GetCongressPerson() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db := globals.DB
		// out := make(map[string]map[string]string)
		rows, err := db.Query("SELECT DISTINCT b.bill_id, v.vote, b.description FROM CongressMember cg, Bill b, Voted v, Donors d WHERE cg.name = 'Nancy Pelosi' AND v.congress_id = cg.congress_id AND v.bill_id = b.bill_id and v.congress_id = cg.congress_id")

		if err != nil {
			log.Panic(err)
		}
		var billID, vote, description string

		for rows.Next() {
			err := rows.Scan(&billID, &vote, &description)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(billID, vote, description)
		}

		rows, err = db.Query("SELECT d.name, pd.amount FROM CongressMember cg, Donors d, PAC_Donations pd, Support s WHERE cg.name = 'Nancy Pelosi' AND s.pac = pd.pac_id AND cg.congress_id = s.supported AND pd.donor_id = d.id ORDER BY pd.amount DESC LIMIT 10;")

		if err != nil {
			log.Panic(err)
		}
		var name string
		var amount int

		for rows.Next() {
			err := rows.Scan(&name, &amount)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(name, amount)
		}
	}
}
