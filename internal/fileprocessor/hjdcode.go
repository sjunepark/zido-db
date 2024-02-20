package fileprocessor

import (
	"github.com/sjunepark/go-gis/internal/types"
	"github.com/xuri/excelize/v2"
)

func GetHjdData(file string) ([]types.HjdCodeRaw, error) {
	f, err := excelize.OpenFile(file)
	if err != nil {
		return nil, err
	}
	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	sheetName := "KIKcd_H"

	var data []types.HjdCodeRaw

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}

	for _, row := range rows {
		row = ensureRowLength(row, 6)
		hjdCode := types.HjdCodeRaw{
			HjdCode:    row[0],
			SdName:     row[1],
			SggName:    row[2],
			EmdName:    row[3],
			CreateDate: row[4],
			ExpireDate: row[5],
		}
		data = append(data, hjdCode)
	}

	return data, nil
}

func ensureRowLength(row []string, length int) []string {
	// Create a new slice with the desired length
	newRow := make([]string, length)
	// Copy data from the original row to the new row
	copy(newRow, row)
	// If the original row is shorter than the desired length,
	// the remaining elements of newRow will already be initialized to ""
	return newRow
}
