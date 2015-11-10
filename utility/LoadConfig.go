package utility

import (
	"encoding/json"
	"finalproject/globals"
	"fmt"
	"log"
	"os"
)

func LoadConfig() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Panic(err)
	}
	decoder := json.NewDecoder(file)
	configuration := make(map[string]map[string]string)
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	globals.DBuser = configuration["Database"]["username"]
	globals.DBpw = configuration["Database"]["password"]
}
