package main

import (
	"github.com/joho/godotenv"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"github.com/sjunepark/go-gis/internal/parser"
	"log"
	"runtime"
	"testing"
)

func benchmark(goroutinesLimit int, b *testing.B) {
	log.Printf("Using %d goroutines, number of CPUs: %d\n", goroutinesLimit, runtime.NumCPU())

	files, err := fileprocessor.GetFilesWithExt("data/input/location_202401", ".txt")
	if err != nil {
		panic(err)
	}
	gobDir := "data/gob/location_202401"

	for i := 0; i < b.N; i++ {
		setupEnv()
		parser.ParseFiles(files, gobDir, goroutinesLimit)
	}
}

func BenchmarkParseFiles1(b *testing.B) {
	benchmark(1, b)
}
func BenchmarkParseFiles2(b *testing.B) {
	benchmark(2, b)
}
func BenchmarkParseFiles4(b *testing.B) {
	benchmark(4, b)
}
func BenchmarkParseFiles8(b *testing.B) {
	benchmark(8, b)
}
func BenchmarkParseFiles16(b *testing.B) {
	benchmark(16, b)
}

func setupEnv() {
	err := godotenv.Load(".env")
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
}
