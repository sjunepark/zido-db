package main

import (
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"os"

	_ "github.com/sjunepark/go-gis/pb/pb_migrations"
	"log"
)

func main() {
	config := pocketbase.Config{
		DefaultDataDir: "pb/pb_data",
	}

	app := pocketbase.NewWithConfig(config)

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
