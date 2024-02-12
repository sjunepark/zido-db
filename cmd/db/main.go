package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/libsql/go-libsql"
	"github.com/sjunepark/go-gis/internal/database"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tursoDB, connector := database.InitTursoDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(tursoDB)
	if connector != nil {
		defer func(connector *libsql.Connector) {
			err := connector.Close()
			if err != nil {
				panic(err)
			}
		}(connector)
	}

	database.CreateUserTable(tursoDB)
	database.InsertUser(tursoDB, "sjunepark", "password", "junepark202012@gmail.com")
	database.SelectUsers(tursoDB)

}
