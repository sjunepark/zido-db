package txtparser

import (
	"bufio"
	"fmt"
	"github.com/sjunepark/go-gis/internal/types"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"os"
	"strings"
	"time"
)

func ReadTxtAndSaveToDb(filepath string, baseDate time.Time) error {
	locations, err := parseTxt(filepath, baseDate)
	if err != nil {
		return err
	}
	err = persistToDb(locations)
	if err != nil {
		return err
	}
	return nil
}

func parseTxt(filepath string, baseDate time.Time) ([]types.Location, error) {
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

		location, err := fieldsToLocation(fields, baseDate)
		if err != nil {
			fmt.Printf("Error processing line: %s. Error: %s\n", getLineSnippet(line), err)
			continue
		}
		locations = append(locations, location)
	}

	return locations, nil
}

func persistToDb(locations []types.Location) error {
	// Add logic to persist to db, in a single transaction.
	return nil
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
