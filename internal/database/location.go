package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sjunepark/go-gis/internal/sqlc"
	"github.com/sjunepark/go-gis/internal/types"
	"log"
	"time"
)

func PersistLocations(db *sql.DB, locations []types.Location) error {
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

func PersistLocation(q *sqlc.Queries, ctx context.Context, l types.Location) error {
	lToInsert := sqlc.InsertLocationParams{
		BJDNumber:          l.BJDNumber,
		SGGNumber:          l.SGGNumber,
		EMDNumber:          l.EMDNumber,
		RoadNumber:         l.RoadNumber,
		UndergroundFlag:    l.UndergroundFlag,
		BuildingMainNumber: l.BuildingMainNumber,
		BuildingSubNumber:  l.BuildingSubNumber,
		SDName:             l.SDName,
		SGGName:            sql.NullString{String: l.SGGName, Valid: true},
		EMDName:            l.EMDName,
		RoadName:           l.RoadName,
		BuildingName:       sql.NullString{String: l.BuildingName, Valid: true},
		PostalNumber:       l.PostalNumber,
		Long:               sql.NullFloat64{Float64: l.Long, Valid: true},
		Lat:                sql.NullFloat64{Float64: l.Lat, Valid: true},
		Crs:                l.Crs,
		X:                  sql.NullFloat64{Float64: l.X, Valid: true},
		Y:                  sql.NullFloat64{Float64: l.Y, Valid: true},
		ValidPosition:      l.ValidPosition,
		BaseDate:           l.BaseDate,
		DatetimeAdded:      l.DatetimeAdded,
	}

	// bug: there's a bug with turso or go-libsql that database level constraint checks do not emit errors even checks fail
	res, err := q.InsertLocation(ctx, lToInsert)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get the number of rows affected: %w", err)
	}
	// bug: this is to substitute for the bug mentioned above, about go-libsql not emitting errors when constraint checks do not emit errors
	if rowCnt == 0 {
		return fmt.Errorf("the insert resulted in 0 rows affected")
	}

	return nil
}
