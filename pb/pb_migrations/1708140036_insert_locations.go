package pb_migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/sjunepark/go-gis/internal/database"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"github.com/sjunepark/go-gis/internal/parser"
	"log"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		gobDir := "data/gob/location_202401"
		err := fileprocessor.CreateDirIfNotExists(gobDir)
		if err != nil {
			return err
		}

		files, err := fileprocessor.GetFilesWithExt(gobDir, ".gob")
		if err != nil {
			return err
		}

		for _, file := range files {
			locations, err := parser.GetLocations(file)
			if err != nil {
				return err
			}

			for _, location := range locations {
				err := database.InsertLocation(db, &location)
				if err != nil {
					return err
				}
			}
			log.Printf("Inserted %d locations from %s\n", len(locations), file)
		}

		return nil
	}, func(db dbx.Builder) error {
		//goland:noinspection SqlWithoutWhere,SqlResolve
		q := db.NewQuery("DELETE FROM locations")
		execute, err := q.Execute()
		if err != nil {
			return err
		}

		rowsAffected, err := execute.RowsAffected()
		if err != nil {
			return err
		}

		log.Printf("Deleted %d rows from location table\n", rowsAffected)
		return nil
	})
}
