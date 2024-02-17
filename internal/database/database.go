package database

import (
	"github.com/pocketbase/dbx"
	"github.com/sjunepark/go-gis/internal/types"
	"log"
)

func InsertLocation(db dbx.Builder, location *types.Location) error {
	err := db.Model(location).Insert()
	if err != nil {
		return err
	}
	return nil
}

func CheckForDuplicateAddr(db dbx.Builder) error {
	// add up queries...
	q := db.NewQuery("SELECT address, COUNT(address) AS address_count FROM (SELECT address FROM location WHERE validPosition = 1 GROUP BY address) AS derived_table GROUP BY address HAVING address_count > 1;")

	type AddressCount struct {
		Address      string `db:"address"`
		AddressCount int
	}

	var addressCounts []AddressCount
	if err := q.All(&addressCounts); err != nil {
		return err
	}

	countQ := db.NewQuery("SELECT COUNT(*) AS count FROM location WHERE validPosition = 1;")
	type Count struct {
		Count int
	}
	var count Count
	err := countQ.One(&count)
	if err != nil {
		return err
	}

	if len(addressCounts) > 0 {
		printLimit := 10
		for i, addressCount := range addressCounts {
			println(addressCount.Address, addressCount.AddressCount)
			if i > printLimit {
				break
			}
		}
		log.Fatalf("There are %d duplicate addresses out of %d locations", len(addressCounts), count.Count)
	}

	log.Printf("No duplicate addresses found out of %d locations", count.Count)
	return nil
}
