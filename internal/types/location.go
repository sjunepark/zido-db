package types

import "time"

// LocationDbRecord type is from 위치정보요약DB(https://business.juso.go.kr/addrlink/elctrnMapProvd/geoDBDwldList.do?menu=%EC%9C%84%EC%B9%98%EC%A0%95%EB%B3%B4%EC%9A%94%EC%95%BDDB)
// X, Y are represented in GRS80 UTM-K
type LocationDbRecord struct {
	SGGNumber           string    `validate:"required,len=5"`
	EntranceNumber      string    `validate:"len=10"`
	BJDNumber           string    `validate:"required,len=10"`
	SDName              string    `validate:"required,max=40"`
	SGGName             string    `validate:"max=40"`
	EMDName             string    `validate:"required,max=40"`
	RoadNumber          string    `validate:"required,len=12"`
	RoadName            string    `validate:"required,max=80"`
	UndergroundFlag     string    `validate:"len=1"`
	BuildingMainNumber  int       `validate:"min=0"`
	BuildingSubNumber   int       `validate:"min=0"`
	BuildingName        string    `validate:"max=40"`
	PostalNumber        string    `validate:"required,len=5"`
	BuildingUseCategory string    `validate:"max=100"`
	BuildingGroupFlag   string    `validate:"len=1"`
	JurisdictionHJD     string    `validate:"max=8"`
	X                   float64   `validate:"required"`
	Y                   float64   `validate:"required"`
	BaseDate            time.Time `validate:"required"`
	DatetimeAdded       time.Time `validate:"required"`
}

// Location type is a struct that represents a location, which is a part of the LocationDbRecord type
type Location struct {
	SGGNumber      string    `validate:"required,len=5"`
	EntranceNumber string    `validate:"len=10"`
	BJDNumber      string    `validate:"required,len=10"`
	SDName         string    `validate:"required,max=40"`
	SGGName        string    `validate:"required,max=40"`
	EMDName        string    `validate:"required,max=40"`
	RoadNumber     string    `validate:"required,len=12"`
	RoadName       string    `validate:"required,max=80"`
	BuildingName   string    `validate:"max=40"`
	PostalNumber   string    `validate:"required,len=5"`
	X              float64   `validate:"required"`
	Y              float64   `validate:"required"`
	BaseDate       time.Time `validate:"required"`
	DatetimeAdded  time.Time `validate:"required"`
}

func (ldr LocationDbRecord) ToLocation() Location {
	return Location{
		SGGNumber:      ldr.SGGNumber,
		EntranceNumber: ldr.EntranceNumber,
		BJDNumber:      ldr.BJDNumber,
		SDName:         ldr.SDName,
		SGGName:        ldr.SGGName,
		EMDName:        ldr.EMDName,
		RoadNumber:     ldr.RoadNumber,
		RoadName:       ldr.RoadName,
		BuildingName:   ldr.BuildingName,
		PostalNumber:   ldr.PostalNumber,
		X:              ldr.X,
		Y:              ldr.Y,
		BaseDate:       ldr.BaseDate,
		DatetimeAdded:  ldr.DatetimeAdded,
	}
}
