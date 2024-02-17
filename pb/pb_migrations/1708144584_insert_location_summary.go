package pb_migrations

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
		rowChan := make(chan LocationSummary)
		errChan := make(chan error, 1)
		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			q := db.NewQuery(`
			SELECT
				(ROW_NUMBER() OVER (ORDER BY address)) AS id,
				address,
				sggName,
				emdName,
				roadName,
				AVG(lat) AS avg_lat,
				AVG(long) AS avg_long,
				AVG(x) AS avg_x,
				AVG(y) AS avg_y,
				GROUP_CONCAT(DISTINCT postalNumber) AS postalNumbers
			FROM
				location
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
				return
			}

			for rows.Next() {
				select {
				case <-ctx.Done():
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
						return
					}
					rowChan <- locationSummary
				}

			}
			close(rowChan)
		}()

		wg.Add(1)
		go func() {
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
					if count%1000000 == 0 {
						log.Printf("Inserted %d rows into location_summary table\n", count)
					}
				}
			}
			log.Printf("Inserted %d rows into location_summary table\n", count)
		}()

		wg.Wait()
		select {
		case err := <-errChan:
			return err
		default:
			return nil
		}
	}, func(db dbx.Builder) error {
		//goland:noinspection SqlResolve,SqlWithoutWhere
		q := db.NewQuery("DELETE FROM location_summary")
		execute, err := q.Execute()
		if err != nil {
			return err
		}

		rowsAffected, err := execute.RowsAffected()
		if err != nil {
			return err
		}

		log.Printf("Deleted %d rows from location_summary table\n", rowsAffected)
		return nil
	})
}

type LocationSummary struct {
	ID            int     `db:"id"`
	Address       string  `db:"address"`
	SggName       string  `db:"sggName"`
	EmdName       string  `db:"emdName"`
	RoadName      string  `db:"roadName"`
	AvgLat        float64 `db:"avg_lat"`
	AvgLong       float64 `db:"avg_long"`
	AvgX          float64 `db:"avg_x"`
	AvgY          float64 `db:"avg_y"`
	PostalNumbers string  `db:"postalNumbers"`
}

func (ls LocationSummary) TableName() string {
	return "location_summary"
}
