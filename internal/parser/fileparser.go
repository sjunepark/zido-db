package parser

import (
	"encoding/gob"
	"github.com/sjunepark/go-gis/internal/types"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func ParseFiles(files []string, gobDir string, noOfGoroutines int) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, noOfGoroutines)

	for _, file := range files {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(file string) {
			defer wg.Done()
			parseFile(gobDir, file)
			<-semaphore
		}(file)
	}

	wg.Wait()
}

func parseFile(gobDir, file string) {
	fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
	gobFilePath := filepath.Join(gobDir, fileName+".gob")

	var locations []types.Location
	locations, err := ParseTxt(file, time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("KST", 9*60*60)))
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
