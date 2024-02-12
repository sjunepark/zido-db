package txtparser

import (
	"bufio"
	"fmt"
	"github.com/sjunepark/go-gis/internal/types"
	"github.com/sjunepark/go-gis/internal/validation"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadTxtAndSaveToDb() error {
	// todo: fix path
	locations, err := processTxt("data/input/location_202401/entrc_sejong.txt")
	if err != nil {
		return err
	}
	err = persistToDb(locations)
	if err != nil {
		return err
	}
	return nil
}

func processTxt(filepath string) ([]types.Location, error) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := korean.EUCKR.NewDecoder()
	reader := transform.NewReader(file, decoder)

	var locations []types.Location

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "|")
		if len(fields) != 18 {
			fmt.Printf("Invalid number of fields: expected 18, got %d. Line begins with: %s...\n", len(fields), getLineSnippet(line))
			continue
		}

		locationDBRecord, err := fieldsToLocationDBRecord(fields, time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("KST", 9*60*60)))
		if err != nil {
			fmt.Printf("Error processing line: %s. Error: %s\n", getLineSnippet(line), err)
			continue
		}
		location := locationDBRecord.ToLocation()
		locations = append(locations, location)
	}

	return locations, nil
}

func persistToDb(locations []types.Location) error {
	// Add logic to persist to db, in a single transaction.
	return nil
}

func fieldsToLocationDBRecord(fields []string, baseDatetime time.Time) (types.LocationDbRecord, error) {
	var locationDBRecord types.LocationDbRecord
	var err error
	locationDBRecord.SGGNumber = fields[0]
	locationDBRecord.EntranceNumber = fields[1]
	locationDBRecord.BJDNumber = fields[2]
	locationDBRecord.SDName = fields[3]
	locationDBRecord.SGGName = fields[4]
	locationDBRecord.EMDName = fields[5]
	locationDBRecord.RoadNumber = fields[6]
	locationDBRecord.RoadName = fields[7]
	locationDBRecord.UndergroundFlag = fields[8]
	locationDBRecord.BuildingMainNumber, err = strconv.Atoi(fields[9])
	if err != nil {
		locationDBRecord.BuildingMainNumber = 0
	}
	locationDBRecord.BuildingSubNumber, err = strconv.Atoi(fields[10])
	if err != nil {
		locationDBRecord.BuildingSubNumber = 0
	}
	locationDBRecord.BuildingName = fields[11]
	locationDBRecord.PostalNumber = fields[12]
	locationDBRecord.BuildingUseCategory = fields[13]
	locationDBRecord.BuildingGroupFlag = fields[14]
	locationDBRecord.JurisdictionHJD = fields[15]
	locationDBRecord.X, err = strconv.ParseFloat(fields[16], 64)
	if err != nil {
		locationDBRecord.X = 0
	}
	locationDBRecord.Y, err = strconv.ParseFloat(fields[17], 64)
	if err != nil {
		locationDBRecord.Y = 0
	}
	locationDBRecord.BaseDate = baseDatetime
	locationDBRecord.DatetimeAdded = time.Now()

	err = validation.ValidateStruct(locationDBRecord)
	if err != nil {
		return types.LocationDbRecord{}, err
	}

	return locationDBRecord, nil
}

func getLineSnippet(line string) string {
	const maxSnippetLength = 50 // Define a maximum snippet length
	if len(line) <= maxSnippetLength {
		return line // Return the entire line if it's short enough
	}
	// Return a substring of the line, adding an ellipsis to indicate it's been truncated
	return line[:maxSnippetLength] + "..."
}
