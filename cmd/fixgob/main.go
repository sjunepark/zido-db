package main

import (
	"context"
	"encoding/gob"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"github.com/sjunepark/go-gis/internal/parser"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func main() {
	gobDir := "data/gob/location_202401"
	files, err := fileprocessor.GetFilesWithExt(gobDir, ".gob")
	if err != nil {
		panic(err)
	}

	outputDir := "data/gob/location_202401_fix"
	err = fileprocessor.RemoveDir(outputDir)
	if err != nil {
		panic(err)
	}
	err = fileprocessor.CreateDirIfNotExists(outputDir)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, runtime.NumCPU())
	errChan := make(chan error, 1)

	for _, file := range files {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(file string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			locations, err := parser.GetLocations(file)
			if err != nil {
				cancel()
				select {
				case errChan <- err:
				default:
				}
				return
			}

			var count int
			for _, location := range locations {
				select {
				case <-ctx.Done():
					return
				default:
					location.AddGroupInfo()
					count++
					if count%1000000 == 0 {
						log.Printf("Added group info for %d locations\n", count)
					}
				}
			}
			log.Printf("Added group info for %d locations from %s\n", count, file)

			// Create file path from outputDir and file. join outputDir and file name
			fixFilePath := filepath.Join(outputDir, filepath.Base(file))
			fixFile, err := os.Create(fixFilePath)
			if err != nil {
				cancel()
				select {
				case errChan <- err:
				default:
				}
				return
			}
			enc := gob.NewEncoder(fixFile)
			if err := enc.Encode(locations); err != nil {
				cancel()
				select {
				case errChan <- err:
				default:
				}
				return
			}
			log.Printf("gob file created: %s\n", fixFile.Name())
		}(file)
	}

	wg.Wait()
	close(errChan)

	select {
	case err := <-errChan:
		panic(err)
	default:
		log.Println("All locations fixed")
	}
}
