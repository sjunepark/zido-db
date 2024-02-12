package database

import (
	"database/sql"
	"fmt"
	"github.com/sjunepark/go-gis/internal/types"
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
		INSERT INTO locations (Id, BJDNumber, SGGNumber, EMDNumber, RoadNumber, UndergroundFlag, BuildingMainNumber, BuildingSubNumber, SDName, SGGName, EMDName, RoadName, BuildingName, PostalNumber, Long, Lat, Crs, X, Y, ValidPosition, BaseDate, DatetimeAdded) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
			fmt.Printf("failed to close statement: %v", err)
		}
	}(stmt)

	// Iterate over the slice of locations and insert each one
	for _, loc := range locations {
		_, err = stmt.Exec(
			loc.Id, loc.BJDNumber, loc.SGGNumber, loc.EMDNumber, loc.RoadNumber, loc.UndergroundFlag,
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
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
