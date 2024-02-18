package main

import (
	"encoding/gob"
	"github.com/sjunepark/go-gis/internal/types"
	"log"
	"os"
)

func main() {
	gobFilePath := "data/gob/location_202401/entrc_sejong.gob"
	gobFile, err := os.Open(gobFilePath)
	if err != nil {
		panic(err)
	}

	var locations []types.Location
	dec := gob.NewDecoder(gobFile)
	if err := dec.Decode(&locations); err != nil {
		panic(err)
	}

	limit := 10
	var count int
	for _, location := range locations {
		count++
		if count >= limit {
			break
		}
		log.Printf("%+v\n", location)
	}
}
