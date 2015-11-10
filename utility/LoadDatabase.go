package utility

import (
	"database/sql"
	"finalproject/globals"
	"fmt"
	"log"

	//Driver package for db
	_ "github.com/lib/pq"
)

// LoadDatabase ...
// returns database to be used
func LoadDatabase() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=donations host=donationsapp.cp7op9t1la4t.us-east-1.rds.amazonaws.com port=5432", globals.DBuser, globals.DBpw))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
