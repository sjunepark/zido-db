package pb_migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/sjunepark/go-gis/internal/database"
	"github.com/sjunepark/go-gis/internal/txtparser"
	"time"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		// add up queries...
		filepath := "data/input/location_202401/entrc_sejong.txt"
		// Date in Korea time
		locations, err := txtparser.ParseTxt(filepath, time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("KST", 9*60*60)))
		if err != nil {
			return err
		}

		// todo: concatenate errors for each file and then return it
		for _, location := range locations {
			err := database.InsertLocation(db, &location)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(db dbx.Builder) error {
		// add down queries...

		return nil
	})
}
