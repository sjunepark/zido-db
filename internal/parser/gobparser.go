package parser

import (
	"encoding/gob"
	"github.com/sjunepark/go-gis/internal/types"
	"log"
	"os"
)

func GetLocations(gobFilePath string) ([]types.Location, error) {
	gobFile, err := os.Open(gobFilePath)
	if err != nil {
		return nil, err
	}

	var locations []types.Location
	dec := gob.NewDecoder(gobFile)
	if err := dec.Decode(&locations); err != nil {
		return nil, err
	}

	log.Printf("gob file decoded to locations: %s\n", gobFile.Name())
	return locations, nil
}
