package parser

import (
	"bufio"
	"github.com/sjunepark/go-gis/internal/types"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"log"
	"os"
	fp "path/filepath"
	"strconv"
	"strings"
	"time"
)

func ParseTxt(filepath string, baseDate time.Time) ([]types.Location, error) {
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

	limit, err := strconv.Atoi(os.Getenv("INPUT_LIMIT"))
	if err != nil {
		limit = 0
	}
	dev := os.Getenv("ENV") == "dev"

	var locationCount int
	for scanner.Scan() {
		if dev && limit > 0 && locationCount >= limit {
			break
		}
		line := scanner.Text()
		fields := strings.Split(line, "|")
		if len(fields) != 18 {
			log.Printf("Invalid number of fields: expected 18, got %d. Line begins with: %s...\n", len(fields), getLineSnippet(line))
			continue
		}

		location, err := fieldsToLocation(fields, baseDate)
		if err != nil {
			log.Printf("Error processing line: %s. Error: %s\n", getLineSnippet(line), err)
			continue
		}
		locations = append(locations, location)

		locationCount++
		if locationCount%10000 == 0 {
			log.Printf("Number of fields appended to locations for file %s: %d\n", fp.Base(filepath), locationCount)
		}
	}

	log.Printf("Number of fields appended to locations for file %s: %d\n", fp.Base(filepath), locationCount)
	return locations, nil
}

func fieldsToLocation(fields []string, baseDate time.Time) (types.Location, error) {
	location, err := types.NewLocation(
		fields[0],  // 시군구코드
		fields[1],  // 입구번호
		fields[2],  // 법정동코드
		fields[3],  // 시도명
		fields[4],  // 시군구명
		fields[5],  // 읍면동명
		fields[6],  // 도로명코드(시군구코드(5) + 도로명번호(7))
		fields[7],  // 도로명
		fields[8],  // 지하여부
		fields[9],  // 건물본번
		fields[10], // 건물부번
		fields[11], // 건물명
		fields[12], // 우편번호
		fields[13], // 건물용도
		fields[14], // 건물구분
		fields[15], // 관할행정동
		fields[16], // X좌표
		fields[17], // Y좌표
		"EPSG:5179",
		baseDate,
	)
	if err != nil {
		return types.Location{}, err
	}

	return location, nil
}

func getLineSnippet(line string) string {
	const maxSnippetLength = 50 // Define a maximum snippet length
	if len(line) <= maxSnippetLength {
		return line // Return the entire line if it's short enough
	}
	// Return a substring of the line, adding an ellipsis to indicate it's been truncated
	return line[:maxSnippetLength] + "..."
}
