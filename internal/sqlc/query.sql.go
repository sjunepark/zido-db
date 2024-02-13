// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const getLocations = `-- name: GetLocations :one
SELECT bjdnumber, sggnumber, emdnumber, roadnumber, undergroundflag, buildingmainnumber, buildingsubnumber, sdname, sggname, emdname, roadname, buildingname, postalnumber, long, lat, crs, x, y, validposition, basedate, datetimeadded
FROM locations
`

func (q *Queries) GetLocations(ctx context.Context) (Location, error) {
	row := q.db.QueryRowContext(ctx, getLocations)
	var i Location
	err := row.Scan(
		&i.Bjdnumber,
		&i.Sggnumber,
		&i.Emdnumber,
		&i.Roadnumber,
		&i.Undergroundflag,
		&i.Buildingmainnumber,
		&i.Buildingsubnumber,
		&i.Sdname,
		&i.Sggname,
		&i.Emdname,
		&i.Roadname,
		&i.Buildingname,
		&i.Postalnumber,
		&i.Long,
		&i.Lat,
		&i.Crs,
		&i.X,
		&i.Y,
		&i.Validposition,
		&i.Basedate,
		&i.Datetimeadded,
	)
	return i, err
}

const insertLocation = `-- name: InsertLocation :execresult
INSERT INTO locations (BJDNumber, SGGNumber, EMDNumber, RoadNumber, UndergroundFlag, BuildingMainNumber,
                       BuildingSubNumber, SDName, SGGName, EMDName, RoadName, BuildingName, PostalNumber, Long, Lat,
                       Crs, X, Y, ValidPosition, BaseDate, DatetimeAdded)
VALUES (?1, ?2, ?3, ?4, ?5, ?6,
        ?7, ?8, ?9, ?10, ?11, ?12, ?13, ?14, ?15,
        ?16, ?17, ?18, ?19, ?20, ?21)
`

type InsertLocationParams struct {
	BJDNumber          string
	SGGNumber          string
	EMDNumber          string
	RoadNumber         string
	UndergroundFlag    int64
	BuildingMainNumber int64
	BuildingSubNumber  int64
	SDName             string
	SGGName            sql.NullString
	EMDName            string
	RoadName           string
	BuildingName       sql.NullString
	PostalNumber       string
	Long               sql.NullFloat64
	Lat                sql.NullFloat64
	Crs                string
	X                  sql.NullFloat64
	Y                  sql.NullFloat64
	ValidPosition      bool
	BaseDate           time.Time
	DatetimeAdded      time.Time
}

func (q *Queries) InsertLocation(ctx context.Context, arg InsertLocationParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertLocation,
		arg.BJDNumber,
		arg.SGGNumber,
		arg.EMDNumber,
		arg.RoadNumber,
		arg.UndergroundFlag,
		arg.BuildingMainNumber,
		arg.BuildingSubNumber,
		arg.SDName,
		arg.SGGName,
		arg.EMDName,
		arg.RoadName,
		arg.BuildingName,
		arg.PostalNumber,
		arg.Long,
		arg.Lat,
		arg.Crs,
		arg.X,
		arg.Y,
		arg.ValidPosition,
		arg.BaseDate,
		arg.DatetimeAdded,
	)
}