package vectorprocessor

import (
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"github.com/sjunepark/go-gis/internal/types"
	"os"
	"os/exec"
	"strings"
)

type ShpFile = types.ShpFile

func Shp2Geojson(shp ShpFile) (outputPath string) {

	geojsonPath, err := fileprocessor.GetOutputPath(shp.AbsPath(), ".geo.json")
	if err != nil {
		panic(err)
	}
	err = fileprocessor.RemoveFileIfExists(geojsonPath)
	if err != nil {
		return ""
	}

	cmd := exec.Command("ogr2ogr", "-f", "GeoJSON", "-t_srs", "EPSG:4326", geojsonPath, shp.AbsPath(), "-s_srs", "EPSG:5179")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	return geojsonPath
}

func Geojson2Topojson(geojsonPath string) (outputPath string, err error) {

	topojsonPath := strings.Replace(geojsonPath, ".geo.json", ".topo.json", 1)
	err = fileprocessor.RemoveFileIfExists(topojsonPath)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("geo2topo", geojsonPath, "-o", topojsonPath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	return topojsonPath, nil
}
