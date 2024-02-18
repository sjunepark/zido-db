package pb_migrations

import (
	"context"
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/sjunepark/go-gis/internal/types"
	"log"
	"sync"
)

func init() {

	m.Register(func(db dbx.Builder) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		rowChan := make(chan types.AddressGroup)
		errChan := make(chan error, 1)
		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			log.Println("Running goroutine for inserting scanning rows from locations_summary table")
			defer wg.Done()
			q := db.NewQuery(`SELECT DISTINCT sdSggEm, addrDetail FROM locations_summary`)

			rows, err := q.Rows()
			if err != nil {
				cancel()
				select {
				case errChan <- err:
				default:
				}
				close(rowChan)
				return
			}

			var count int
			for rows.Next() {
				select {
				case <-ctx.Done():
					close(rowChan)
					return
				default:
					var addressGroup types.AddressGroup
					err := rows.ScanStruct(&addressGroup)
					if err != nil {
						cancel()
						select {
						case errChan <- err:
						default:
						}
						close(rowChan)
						return
					}
					rowChan <- addressGroup
					count++
					if count%100000 == 0 {
						log.Printf("Scanned %d rows from locations_summary table\n", count)
					}
				}
			}
			log.Printf("Scanned %d rows from locations_summary table\n", count)
			close(rowChan)
		}()

		wg.Add(1)
		go func() {
			log.Println("Running goroutine for inserting rows into address_group table")
			defer wg.Done()
			var count int
			for addressGroup := range rowChan {
				select {
				case <-ctx.Done():
					return
				default:
					if err := db.Model(&addressGroup).Insert(); err != nil {
						cancel()
						select {
						case errChan <- err:
						default:
						}
						return
					}
					count++
					if count%100000 == 0 {
						log.Printf("Inserted %d rows into address_group table\n", count)
					}
				}
			}
			log.Printf("Inserted %d rows into address_group table\n", count)
		}()

		wg.Wait()
		close(errChan)

		select {
		case err := <-errChan:
			return err
		default:
			return nil
		}
	}, func(db dbx.Builder) error {
		//goland:noinspection SqlResolve,SqlWithoutWhere
		q := db.NewQuery("DELETE FROM address_group")
		execute, err := q.Execute()
		if err != nil {
			return err
		}

		rowsAffected, err := execute.RowsAffected()
		if err != nil {
			return err
		}

		log.Printf("Deleted %d rows from address_group table\n", rowsAffected)
		return nil
	})
}
