package database

import (
	"github.com/pocketbase/dbx"
	"github.com/sjunepark/go-gis/internal/types"
)

func InsertLocation(db dbx.Builder, location *types.Location) error {
	err := db.Model(location).Insert()
	if err != nil {
		return err
	}
	return nil
}
