package main

import (
	"fmt"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
)

func main() {
	path := "data/input/jscode20240208/KIKcd_H.20240208.xlsx"

	hjdCodes, err := fileprocessor.GetHjdData(path)
	if err != nil {
		panic(err)
	}

	for _, code := range hjdCodes {
		fmt.Printf("%+v\n", code)
	}
}
