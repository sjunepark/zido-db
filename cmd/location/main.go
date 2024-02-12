package main

import (
	"github.com/sjunepark/go-gis/internal/location/txtparser"
	"github.com/sjunepark/go-gis/internal/validation"
	"time"
)

func main() {
	validation.Init()

	filepath := "data/input/location_202401/entrc_sejong.txt"
	baseDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("KST", 9*60*60))

	locations, err := txtparser.ParseTxt(filepath, baseDate)
	if err != nil {
		panic(err)
	}

	println("Parsed locations: ", len(locations))
}
