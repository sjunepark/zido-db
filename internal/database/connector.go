package database

import (
	"database/sql"
	"github.com/libsql/go-libsql"
	"os"
)

func InitTursoDB() (*sql.DB, *libsql.Connector) {

	switch os.Getenv("TURSO_ENV") {
	case "local":
		dbName := os.Getenv("TURSO_LOCAL_LOCATION_DB_PATH")
		db, err := sql.Open("libsql", "file:"+dbName)
		if err != nil {
			panic(err)
		}
		return db, nil
	case "embedded":
		dbName := os.Getenv("TURSO_EMBEDDED_LOCATION_DB_PATH")
		primaryUrl := os.Getenv("TURSO_EMBEDDED_PRIMARY_URL")
		authToken := os.Getenv("TURSO_EMBEDDED_AUTH_TOKEN")

		connector, err := libsql.NewEmbeddedReplicaConnector("file:"+dbName, primaryUrl, authToken)
		if err != nil {
			panic(err)
		}
		db := sql.OpenDB(connector)
		return db, connector
	default:
		panic("Unknown environment")
	}
}
