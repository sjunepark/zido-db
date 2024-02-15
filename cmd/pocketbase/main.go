package main

import (
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	_ "github.com/sjunepark/go-gis/pb/pb_migrations"
	"log"
	"os"
)

func main() {
	app := pocketbase.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("ENV")

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Dir:         "pb/pb_migrations",
		Automigrate: env == "dev",
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
