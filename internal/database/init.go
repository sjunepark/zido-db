package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func InitTursoDB() *sql.DB {
	switch os.Getenv("TURSO_ENV") {
	case "local":
		dbName := os.Getenv("TURSO_LOCAL_LOCATION_DB_PATH")
		db, err := sql.Open("sqlite3", "file:"+dbName)
		if err != nil {
			panic(err)
		}
		return db
	default:
		panic("Unknown environment")
	}
}
