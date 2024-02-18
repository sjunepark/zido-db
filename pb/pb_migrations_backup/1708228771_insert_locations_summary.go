package pb_migrations_backup

import (
	"context"
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
	"log"
	"sync"
)

func init() {

	m.Register(func(db dbx.Builder) error {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		rowChan := make(chan LocationSummary)
		errChan := make(chan error, 1)
		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			log.Println("Running goroutine for inserting scanning rows from locations table")
			defer wg.Done()
			q := db.NewQuery(`
			SELECT
				(ROW_NUMBER() OVER (ORDER BY address)) AS id,
				address,
				addressGroup,
				roadNameGroup,
				AVG(lat) AS lat,
				AVG(long) AS long,
				AVG(x) AS x,
				AVG(y) AS y
			FROM
				locations
			WHERE
				validPosition = 1
			GROUP BY
				address;`)

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
					var locationSummary LocationSummary
					err := rows.ScanStruct(&locationSummary)
					if err != nil {
						cancel()
						select {
						case errChan <- err:
						default:
						}
						close(rowChan)
						return
					}
					rowChan <- locationSummary
					count++
					if count%100000 == 0 {
						log.Printf("Scanned %d rows from locations table\n", count)
					}
				}
			}
			log.Printf("Scanned %d rows from locations table\n", count)
			close(rowChan)
		}()

		wg.Add(1)
		go func() {
			log.Println("Running goroutine for inserting rows into locations_summary table")
			defer wg.Done()
			var count int
			for locationSummary := range rowChan {
				select {
				case <-ctx.Done():
					return
				default:
					if err := db.Model(&locationSummary).Insert(); err != nil {
						cancel()
						select {
						case errChan <- err:
						default:
						}
						return
					}
					count++
					if count%100000 == 0 {
						log.Printf("Inserted %d rows into locations_summary table\n", count)
					}
				}
			}
			log.Printf("Inserted %d rows into locations_summary table\n", count)
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
		q := db.NewQuery("DELETE FROM locations_summary")
		execute, err := q.Execute()
		if err != nil {
			return err
		}

		rowsAffected, err := execute.RowsAffected()
		if err != nil {
			return err
		}

		log.Printf("Deleted %d rows from locations_summary table\n", rowsAffected)
		return nil
	})
}

type LocationSummary struct {
	ID            int     `db:"id"`
	Address       string  `db:"address"`
	AddressGroup  string  `db:"addressGroup"`
	RoadNameGroup string  `db:"roadNameGroup"`
	Lat           float64 `db:"lat"`
	Long          float64 `db:"long"`
	X             float64 `db:"x"`
	Y             float64 `db:"y"`
}

func (ls LocationSummary) TableName() string {
	return "locations_summary"
}
