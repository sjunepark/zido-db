package database

import (
	"database/sql"
	"fmt"
	"github.com/sjunepark/go-gis/internal/types"
	"log"
	"time"
)

func PersistToDb(db *sql.DB, locations []types.Location) error {
	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Prepare the statement for inserting data
	stmt, err := tx.Prepare(`
	INSERT INTO locations (
		BJDNumber, SGGNumber, EMDNumber, RoadNumber, UndergroundFlag, BuildingMainNumber, BuildingSubNumber, 
		SDName, SGGName, EMDName, RoadName, BuildingName, PostalNumber, Long, Lat, Crs, X, Y, ValidPosition, 
		BaseDate, DatetimeAdded
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(SGGNumber, EMDNumber, RoadNumber, UndergroundFlag, BuildingMainNumber, BuildingSubNumber) DO UPDATE SET
		BJDNumber=excluded.BJDNumber, SDName=excluded.SDName, SGGName=excluded.SGGName, EMDName=excluded.EMDName, RoadName=excluded.RoadName, 
		BuildingName=excluded.BuildingName, PostalNumber=excluded.PostalNumber, Long=excluded.Long, Lat=excluded.Lat, 
		Crs=excluded.Crs, X=excluded.X, Y=excluded.Y, ValidPosition=excluded.ValidPosition, 
		BaseDate=excluded.BaseDate, DatetimeAdded=excluded.DatetimeAdded
`)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return fmt.Errorf("failed to prepare statement, failed to rollback transaction: %w", rollbackErr)
		}
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Print("failed to close the statement")
			panic(err)
		}
	}(stmt)

	// Iterate over the slice of locations and insert each one
	var locationCount int
	for _, loc := range locations {
		_, err = stmt.Exec(
			loc.BJDNumber, loc.SGGNumber, loc.EMDNumber, loc.RoadNumber, loc.UndergroundFlag,
			loc.BuildingMainNumber, loc.BuildingSubNumber, loc.SDName, loc.SGGName, loc.EMDName, loc.RoadName,
			loc.BuildingName, loc.PostalNumber, loc.Long, loc.Lat, loc.Crs, loc.X, loc.Y, loc.ValidPosition,
			loc.BaseDate.Format(time.RFC3339), loc.DatetimeAdded.Format(time.RFC3339),
		)
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return fmt.Errorf("failed to insert data, failed to rollback transaction: %w", rollbackErr)
			}
			return fmt.Errorf("failed to insert data: %w", err)
		}

		locationCount++
		if locationCount%1000 == 0 {
			log.Printf("Number of fields inserted to the database: %d\n", locationCount)
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("Successfully persisted %d locations to the database\n", locationCount)
	return nil
}

// todo: delete after fix
func PersistFirstToDb(db *sql.DB, l types.Location) error {
	query := `INSERT INTO locations (
		BJDNumber, SGGNumber, EMDNumber, RoadNumber, UndergroundFlag, BuildingMainNumber, BuildingSubNumber, 
		SDName, SGGName, EMDName, RoadName, BuildingName, PostalNumber, Long, Lat, Crs, X, Y, ValidPosition, 
		BaseDate, DatetimeAdded
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := db.Exec(
		query,
		l.BJDNumber, l.SGGNumber, l.EMDNumber, l.RoadNumber, l.UndergroundFlag,
		l.BuildingMainNumber, l.BuildingSubNumber, l.SDName, l.SGGName, l.EMDName, l.RoadName,
		l.BuildingName, l.PostalNumber, l.Long, l.Lat, l.Crs, l.X, l.Y, l.ValidPosition,
		l.BaseDate.Format(time.RFC3339), l.DatetimeAdded.Format(time.RFC3339),
	)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get the last inserted ID: %w", err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get the number of rows affected: %w", err)
	}

	// Both returns 0
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
	return nil
}
