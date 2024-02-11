package main

import (
	"fmt"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"github.com/sjunepark/go-gis/internal/vectorprocessor"
)

func main() {
	files, err := fileprocessor.CollectShpFiles("data/input")
	if err != nil {
		return
	}
	for _, f := range files {
		geojsonPath := vectorprocessor.Shp2Geojson(f)

		//simplifiedGeojsonPath, err := vectorprocessor.SimplifyGeojson(geojsonPath, 0.02)
		//if err != nil {
		//	return
		//}
		//fmt.Printf("Simplified GeoJSON file created: %s\n", simplifiedGeojsonPath)

		topojsonPath, err := vectorprocessor.Geojson2Topojson(geojsonPath)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Printf("TopoJSON file created: %s\n", topojsonPath)

		outputPath, err := vectorprocessor.SimplifyTopojson(topojsonPath, 5e5)
		if err != nil {
			return
		}

		fmt.Printf("Simplified TopoJSON file created: %s\n", outputPath)
	}
}
