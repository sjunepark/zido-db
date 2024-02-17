package main

import (
	"encoding/gob"
	"github.com/joho/godotenv"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"github.com/sjunepark/go-gis/internal/txtparser"
	"github.com/sjunepark/go-gis/internal/types"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	files, err := fileprocessor.GetTxtFiles("data/input/location_202401")
	if err != nil {
		panic(err)
	}

	gobDir := "data/gob/location_202401"
	err = fileprocessor.RemoveDir(gobDir)
	log.Printf("Removed directory: %s\n", gobDir)
	if err != nil {
		panic(err)
	}
	err = fileprocessor.CreateDirIfNotExists(gobDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
		gobFilePath := filepath.Join(gobDir, fileName+".gob")

		var locations []types.Location
		locations, err = txtparser.ParseTxt(file, time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("KST", 9*60*60)))
		if err != nil {
			panic(err)
		}

		gobFile, err := os.Create(gobFilePath)
		if err != nil {
			panic(err)
		}
		enc := gob.NewEncoder(gobFile)
		if err := enc.Encode(locations); err != nil {
			panic(err)
		}
		log.Printf("gob file created: %s\n", gobFile.Name())
	}
}
