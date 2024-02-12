package main

import (
	"github.com/sjunepark/go-gis/internal/location/txtparser"
	"github.com/sjunepark/go-gis/internal/validation"
)

func main() {
	validation.Init()

	err := txtparser.ReadTxtAndSaveToDb()
	if err != nil {
		panic(err)
	}
}
