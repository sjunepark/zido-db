package pb_migrations

import (
	"github.com/pocketbase/dbx"
	m "github.com/pocketbase/pocketbase/migrations"
	"log"
)

func init() {

	var locationSummaries []LocationSummary

	m.Register(func(db dbx.Builder) error {
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

		err := q.All(&locationSummaries)
		if err != nil {
			return err
		}

		var count int
		for _, locationSummary := range locationSummaries {
			err := db.Model(&locationSummary).Insert()
			if err != nil {
				return err
			}
			count++
			if count%1000000 == 0 {
				log.Printf("Inserted %d rows into location_summary table\n", count)
			}

		}

		log.Printf("Inserted %d rows into location_summary table\n", count)
		return nil
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
