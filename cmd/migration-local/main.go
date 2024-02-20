// This is custom goose binary with sqlite3 support only.

package main

import (
	"context"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	_ "github.com/sjunepark/go-gis/db/supabase/migrations"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("goose: failed to parse flags: %v\n", err)
	}
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	err = godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbString := os.Getenv("SB_DB_LOCAL")
	dir := "db/supabase/migrations"

	db, err := goose.OpenDBWithDriver("pgx", dbString)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	var arguments []string
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	ctx := context.Background()
	command := args[0]
	if err := goose.RunContext(ctx, command, db, dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
