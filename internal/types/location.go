package types

import (
	"fmt"
	"github.com/sjunepark/go-gis/internal/validation"
	"github.com/twpayne/go-proj/v10"
	"strconv"
	"time"
)

// LocationDbRecord type is from 위치정보요약DB(https://business.juso.go.kr/addrlink/elctrnMapProvd/geoDBDwldList.do?menu=%EC%9C%84%EC%B9%98%EC%A0%95%EB%B3%B4%EC%9A%94%EC%95%BDDB)
// X, Y are represented in GRS80 UTM-K, which is EPSG:5179
// Id: https://business.juso.go.kr/addrlink/tchnlgyNotice/tchnlgyNoticeDetail.do?currentPage=1&keyword=&searchType=searchType%3D&noticeMgtSn=22899&noticeType=TCHNLGYNOTICE
type LocationDbRecord struct {
	Location            Location
	EntranceNumber      string `validate:"max=10"`
	BuildingUseCategory string `validate:"max=100"`
	BuildingGroupFlag   string `validate:"len=1"`
	JurisdictionHJD     string `validate:"max=8"`
}

// Location type is the struct which has relevant fields to persist to the database
type Location struct {
	Id                 string  `validate:"required,len=26"`
	BJDNumber          string  `validate:"required,len=10"` // 법정동코드: 시군구코드(5) + 읍면동코드(3) + 00
	SGGNumber          string  `validate:"required,len=5"`  // 시군구코드
	EMDNumber          string  `validate:"required,len=3"`
	RoadNumber         string  `validate:"required,len=7"`
	UndergroundFlag    string  `validate:"len=1"`
	BuildingMainNumber string  `validate:"required,max=5"`
	BuildingSubNumber  string  `validate:"max=5"`
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
	ValidPosition      bool
	BaseDate           time.Time `validate:"required"`
	DatetimeAdded      time.Time `validate:"required"`
}

func NewLocation(sggNumber, entranceNumber, bjdNumber, sdName, sggName, emdName, roadNumber, roadName, undergroundFlag, buildingMainNumber, buildingSubNumber, buildingName, postalNumber, buildingUseCategory, buildingGroupFlag, jurisdictionHJD, x, y, crs string, baseDate time.Time) (Location, error) {
	datetimeAdded := time.Now()

	// PK: 시군구코드(5) + 읍면동코드(3) + 도로명번호(7) + 지하여부(1) + 건물본번(5) + 건불부번(5)
	id := sggNumber + bjdNumber[5:8] + roadNumber[5:12] + undergroundFlag + fmt.Sprintf("%05s", buildingMainNumber) + fmt.Sprintf("%5s", buildingSubNumber)
	if len(id) != 26 {
		panic(fmt.Sprintf("ID length is not 26. ID: %s", id))
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
	var validPosition bool
	if floatX != 0 && floatY != 0 {
		pj, err := proj.NewCRSToCRS(crs, "EPSG:4326", nil)
		if err != nil {
			panic(err)
		}
		coord := proj.NewCoord(floatY, floatX, 0, 0) // The api uses lat, long
		pjCoord, err = pj.Forward(coord)
		if err != nil {
			panic(err)
		}
		lat = pjCoord.X()
		long = pjCoord.Y()
		validPosition = true
	} else {
		long = 0
		lat = 0
		validPosition = false
	}

	location := Location{
		Id:                 id,
		BJDNumber:          bjdNumber,
		SGGNumber:          sggNumber,
		EMDNumber:          bjdNumber[5:8],
		RoadNumber:         roadNumber[5:12],
		UndergroundFlag:    undergroundFlag,
		BuildingMainNumber: buildingMainNumber,
		BuildingSubNumber:  buildingSubNumber,
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

	locationDbRecord := LocationDbRecord{
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
