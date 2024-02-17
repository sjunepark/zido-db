package types

import (
	"fmt"
	"github.com/sjunepark/go-gis/internal/validation"
	"github.com/twpayne/go-proj/v10"
	"strconv"
	"strings"
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
	Pk                 string    `db:"pk" validate:"required,len=26"`
	BJDNumber          string    `db:"bjdNumber" validate:"required,len=10"` // 법정동코드: 시군구코드(5) + 읍면동코드(3) + 00
	SGGNumber          string    `db:"sggNumber" validate:"required,len=5"`  // 시군구코드
	EMDNumber          string    `db:"emdNumber" validate:"required,len=3"`  // 읍면동코드
	RoadNumber         string    `db:"roadNumber" validate:"required,len=7"`
	UndergroundFlag    string    `db:"undergroundFlag" validate:"required,len=1"`
	BuildingMainNumber string    `db:"buildingMainNumber" validate:"required,max=5"`
	BuildingSubNumber  string    `db:"buildingSubNumber" validate:"max=5"`
	SDName             string    `db:"sdName" validate:"required,max=40"`
	SGGName            string    `db:"sggName" validate:"max=40"`
	EMDName            string    `db:"emdName" validate:"required,max=40"`
	RoadName           string    `db:"roadName" validate:"required,max=80"`
	BuildingName       string    `db:"buildingName" validate:"max=40"`
	PostalNumber       string    `db:"postalNumber" validate:"required,len=5"`
	Long               float64   `db:"long" validate:"max=180,min=-180"`
	Lat                float64   `db:"lat" validate:"max=90,min=-90"`
	Crs                string    `db:"crs" validate:"required"`
	X                  float64   `db:"x"`
	Y                  float64   `db:"y"`
	ValidPosition      bool      `db:"validPosition"`
	BaseDate           time.Time `db:"baseDate" validate:"required"`
	Address            string    `db:"address" validate:"required,max=100"` // 시도 + 시군구 + 읍면동 + 도로명 + 건물본번 + 건물부번
}

func NewLocation(sggNumber, entranceNumber, bjdNumber, sdName, sggName, emdName, roadNumber, roadName, undergroundFlag, buildingMainNumber, buildingSubNumber, buildingName, postalNumber, buildingUseCategory, buildingGroupFlag, jurisdictionHJD, x, y, crs string, baseDate time.Time) (Location, error) {
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
		coord := proj.NewCoord(floatY, floatX, 0, 0) // The api uses lat, long, which is the opposite
		pjCoord, err = pj.Forward(coord)
		if err != nil {
			panic(err)
		}
		long = pjCoord.Y()
		lat = pjCoord.X()
		validPosition = true
	} else {
		long = 0
		lat = 0
		validPosition = false
	}

	emdNumber := bjdNumber[5:8]
	roadNumber = roadNumber[5:12]

	address := buildAddress(sdName, sggName, emdName, roadName, buildingMainNumber, buildingSubNumber)
	if strings.Contains(address, "  ") {
		return Location{}, fmt.Errorf("formatted address contains double spaces: %s", address)
	}

	location := Location{
		Pk:                 sggNumber + emdNumber + roadNumber + undergroundFlag + fmt.Sprintf("%05s", buildingMainNumber) + fmt.Sprintf("%05s", buildingSubNumber),
		BJDNumber:          bjdNumber,
		SGGNumber:          sggNumber,
		EMDNumber:          emdNumber,
		RoadNumber:         roadNumber,
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
		Address:            address,
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

func buildAddress(sdName, sggName, emdName, roadName, buildingMainNumber, buildingSubNumber string) string {
	var buildingNumber string
	if buildingSubNumber != "" && buildingSubNumber != "0" {
		buildingNumber = buildingMainNumber + "-" + buildingSubNumber
	} else {
		buildingNumber = buildingMainNumber
	}

	emName := func() string {
		if strings.HasSuffix(emdName, "동") {
			return ""
		}
		return emdName
	}()

	var addressParts []string
	for _, s := range []string{sdName, sggName, emName, roadName, buildingNumber} {
		if strings.TrimSpace(s) != "" {
			addressParts = append(addressParts, s)
		}
	}

	return strings.Join(addressParts, " ")
}
