package pb_migrations

import (
	"encoding/gob"
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/sjunepark/go-gis/internal/database"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"github.com/sjunepark/go-gis/internal/txtparser"
	"github.com/sjunepark/go-gis/internal/types"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		files, err := fileprocessor.GetTxtFiles("data/input/location_202401")
		if err != nil {
			return err
		}

		gobDir := "data/gob/location_202401"
		err = fileprocessor.CreateDirIfNotExists(gobDir)
		if err != nil {
			return err
		}

		for _, file := range files {
			fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
			gobFilePath := filepath.Join(gobDir, fileName+".gob")

			var locations []types.Location
			if _, err := os.Stat(gobFilePath); os.IsNotExist(err) {
				log.Printf("gob file not found, parsing txt file: %s\n", file)
				locations, err = txtparser.ParseTxt(file, time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("KST", 9*60*60)))
				if err != nil {
					return err
				}

				gobFile, err := os.Create(gobFilePath)
				if err != nil {
					return err
				}
				enc := gob.NewEncoder(gobFile)
				if err := enc.Encode(locations); err != nil {
					return err
				}
				log.Printf("gob file created: %s\n", gobFile.Name())
			} else {
				log.Printf("gob file found, decoding gob file: %s\n", gobFilePath)
				gobFile, err := os.Open(gobFilePath)
				if err != nil {
					return err
				}
				dec := gob.NewDecoder(gobFile)
				if err := dec.Decode(&locations); err != nil {
					return err
				}
				log.Printf("gob file decoded: %s\n", gobFile.Name())
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
		//goland:noinspection SqlWithoutWhere
		q := db.NewQuery("DELETE FROM location")
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
