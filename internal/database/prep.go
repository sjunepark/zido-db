package database

import "database/sql"

func CreateLocationsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS locations (
    Id                 TEXT PRIMARY KEY,
    BJDNumber          TEXT NOT NULL,
    SGGNumber          TEXT NOT NULL,
    EMDNumber          TEXT NOT NULL,
    RoadNumber         TEXT NOT NULL,
    UndergroundFlag    TEXT,
    BuildingMainNumber TEXT NOT NULL,
    BuildingSubNumber  TEXT,
    SDName             TEXT NOT NULL,
    SGGName            TEXT,
    EMDName            TEXT NOT NULL,
    RoadName           TEXT NOT NULL,
    BuildingName       TEXT,
    PostalNumber       TEXT NOT NULL,
    Long               REAL,
    Lat                REAL,
    Crs                TEXT NOT NULL,
    X                  REAL,
    Y                  REAL,
    ValidPosition      BOOLEAN,
    BaseDate           DATETIME NOT NULL,
    DatetimeAdded      DATETIME NOT NULL
);`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func DropLocationsTable(db *sql.DB) {
	query := `DROP TABLE IF EXISTS locations`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func queryLocationsTable(db *sql.DB) {
	query := `SELECT * FROM locations LIMIT 10`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	// todo: check code
	for i := 0; i < 10; i++ {
		rows.Next()
		var id, bjdNumber, sggNumber, emdNumber, roadNumber, undergroundFlag, buildingMainNumber, buildingSubNumber, sdName, sggName, emdName, roadName, buildingName, postalNumber, crs, baseDate, datetimeAdded string
		var long, lat, x, y float64
		var validPosition bool
		err = rows.Scan(&id, &bjdNumber, &sggNumber, &emdNumber, &roadNumber, &undergroundFlag, &buildingMainNumber, &buildingSubNumber, &sdName, &sggName, &emdName, &roadName, &buildingName, &postalNumber, &long, &lat, &crs, &x, &y, &validPosition, &baseDate, &datetimeAdded)
		if err != nil {
			panic(err)
		}
	}
}
