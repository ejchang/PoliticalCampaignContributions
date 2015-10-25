package utility

import (
	"database/sql"
	"log"

	//Driver package for db
	_ "github.com/lib/pq"
)

// LoadDatabase ...
// returns database to be used
func LoadDatabase() *sql.DB {
	db, err := sql.Open("postgres", "user=ethanchang dbname=donations sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
