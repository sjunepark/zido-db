package main

import (
	"github.com/joho/godotenv"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"github.com/sjunepark/go-gis/internal/parser"
	"log"
	"runtime"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	files, err := fileprocessor.GetFilesWithExt("data/input/location_202401", ".txt")
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

	parser.ParseFiles(files, gobDir, runtime.NumCPU())
}
