package types

import (
	"github.com/sjunepark/go-gis/internal/validation"
	"github.com/twpayne/go-proj/v10"
	"strconv"
	"time"
)

// LocationDbRec type is from 위치정보요약DB(https://business.juso.go.kr/addrlink/elctrnMapProvd/geoDBDwldList.do?menu=%EC%9C%84%EC%B9%98%EC%A0%95%EB%B3%B4%EC%9A%94%EC%95%BDDB)
type LocationDbRec struct {
	Location            Location
	EntranceNumber      string `validate:"max=10"`
	BuildingUseCategory string `validate:"max=100"`
	BuildingGroupFlag   string `validate:"len=1"`
	JurisdictionHJD     string `validate:"max=8"`
}

// Location type is the struct which has relevant fields to persist to the database
// PK: 시군구코드(5) + 읍면동코드(3) + 도로명번호(7) + 지하여부(1) + 건물본번(5) + 건불부번(5)
// https://business.juso.go.kr/addrlink/tchnlgyNotice/tchnlgyNoticeDetail.do?currentPage=1&keyword=&searchType=searchType%3D&noticeMgtSn=22899&noticeType=TCHNLGYNOTICE
// X, Y are represented in GRS80 UTM-K, which is EPSG:5179
type Location struct {
	BJDNumber          string  `validate:"required,len=10"` // 법정동코드: 시군구코드(5) + 읍면동코드(3) + 00
	SGGNumber          string  `validate:"required,len=5"`  // 시군구코드
	EMDNumber          string  `validate:"required,len=3"`
	RoadNumber         string  `validate:"required,len=7"`
	UndergroundFlag    int64   `validate:"max=2,min=0"`
	BuildingMainNumber int64   `validate:"required,max=99999"`
	BuildingSubNumber  int64   `validate:"max=99999"`
	SDName             string  `validate:"required,max=40"`
	SGGName            string  `validate:"max=40"`
	EMDName            string  `validate:"required,max=40"`
	RoadName           string  `validate:"required,max=80"`
	BuildingName       string  `validate:"max=40"`
	PostalNumber       string  `validate:"required,len=5"`
	Long               float64 `validate:"max=180,min=-180"`
	Lat                float64 `validate:"max=90,min=-90"`
	Crs                string  `validate:"required"`
	X                  float64
	Y                  float64
	ValidPosition      int64     `validate:"max=1,min=0"`
	BaseDate           time.Time `validate:"required"`
	DatetimeAdded      time.Time `validate:"required"`
}

func NewLocation(sggNumber, entranceNumber, bjdNumber, sdName, sggName, emdName, roadNumber, roadName, undergroundFlag, buildingMainNumber, buildingSubNumber, buildingName, postalNumber, buildingUseCategory, buildingGroupFlag, jurisdictionHJD, x, y, crs string, baseDate time.Time) (Location, error) {
	datetimeAdded := time.Now()

	undergroundFlagInt, err := strconv.ParseInt(undergroundFlag, 10, 64)
	if err != nil {
		undergroundFlagInt = 0
	}
	buildingMainNumberInt, err := strconv.ParseInt(buildingMainNumber, 10, 64)
	if err != nil {
		return Location{}, err
	}
	buildingSubNumberInt, err := strconv.ParseInt(buildingSubNumber, 10, 64)
	if err != nil {
		buildingSubNumberInt = 0
	}

	floatX, err := strconv.ParseFloat(x, 64)
	if err != nil {
		floatX = 0
	}
	floatY, err := strconv.ParseFloat(y, 64)
	if err != nil {
		floatY = 0
	}

	var pjCoord proj.Coord
	var long float64
	var lat float64
	var validPosition int64
	if floatX != 0 && floatY != 0 {
		pj, err := proj.NewCRSToCRS(crs, "EPSG:4326", nil)
		if err != nil {
			panic(err)
		}
		coord := proj.NewCoord(floatY, floatX, 0, 0) // The api uses lat, long, which is the opposite
		pjCoord, err = pj.Forward(coord)
		if err != nil {
			panic(err)
		}
		long = pjCoord.Y()
		lat = pjCoord.X()
		validPosition = 1
	} else {
		long = 0
		lat = 0
		validPosition = 0
	}

	location := Location{
		BJDNumber:          bjdNumber,
		SGGNumber:          sggNumber,
		EMDNumber:          bjdNumber[5:8],
		RoadNumber:         roadNumber[5:12],
		UndergroundFlag:    undergroundFlagInt,
		BuildingMainNumber: buildingMainNumberInt,
		BuildingSubNumber:  buildingSubNumberInt,
		SDName:             sdName,
		SGGName:            sggName,
		EMDName:            emdName,
		RoadName:           roadName,
		BuildingName:       buildingName,
		PostalNumber:       postalNumber,
		Long:               long,
		Lat:                lat,
		Crs:                crs,
		X:                  floatX,
		Y:                  floatY,
		ValidPosition:      validPosition,
		BaseDate:           baseDate,
		DatetimeAdded:      datetimeAdded,
	}

	err = validation.ValidateStruct(location)
	if err != nil {
		panic(err)
	}

	locationDbRecord := LocationDbRec{
		Location:            location,
		EntranceNumber:      entranceNumber,
		BuildingUseCategory: buildingUseCategory,
		BuildingGroupFlag:   buildingGroupFlag,
		JurisdictionHJD:     jurisdictionHJD,
	}

	err = validation.ValidateStruct(locationDbRecord)
	if err != nil {
		panic(err)
	}

	return location, nil
}
