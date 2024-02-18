package pb_migrations

import (
	"context"
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/sjunepark/go-gis/internal/database"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"github.com/sjunepark/go-gis/internal/parser"
	"log"
	"runtime"
	"sync"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		gobDir := "data/gob/location_202401_fix"
		err := fileprocessor.CreateDirIfNotExists(gobDir)
		if err != nil {
			return err
		}

		files, err := fileprocessor.GetFilesWithExt(gobDir, ".gob")
		if err != nil {
			return err
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
						err := database.InsertLocation(db, &location)
						if err != nil {
							cancel()
							select {
							case errChan <- err:
							default:
							}
							return
						}
						count++
						if count%100000 == 0 {
							log.Printf("Inserted %d locations from %s\n", count, file)
						}
					}
				}
				log.Printf("Inserted %d locations from %s\n", len(locations), file)
			}(file)
		}

		wg.Wait()
		close(errChan)

		select {
		case err := <-errChan:
			return err
		default:
			log.Println("All locations inserted")
			return nil
		}
	}, func(db dbx.Builder) error {
		//goland:noinspection SqlWithoutWhere,SqlResolve
		q := db.NewQuery("DELETE FROM locations")
		execute, err := q.Execute()
		if err != nil {
			return err
		}

		rowsAffected, err := execute.RowsAffected()
		if err != nil {
			return err
		}

		log.Printf("Deleted %d rows from locations table\n", rowsAffected)
		return nil
	})
}
