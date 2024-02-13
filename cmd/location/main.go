package main

import (
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/libsql/go-libsql"
	"github.com/sjunepark/go-gis/internal/database"
	"github.com/sjunepark/go-gis/internal/location/txtparser"
	"github.com/sjunepark/go-gis/internal/sqlc"
	"github.com/sjunepark/go-gis/internal/validation"
	"log"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	validation.Init()

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

	ctx := context.Background()
	queries := sqlc.New(tursoDB)

	filepath := "data/input/location_202401/entrc_sejong.txt"
	baseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("KST", 9*60*60))

	locations, err := txtparser.ParseTxt(filepath, baseDate)
	if err != nil {
		panic(err)
	}

	err = database.PersistLocation(queries, ctx, locations[0])
	if err != nil {
		log.Fatal(err)
	}
}
