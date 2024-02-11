package gisprocessor

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

	cmd := exec.Command("shp2json", shp.AbsPath(), "--crs-name", "EPSG:5179", "--encoding", "EUC-KR", "-o", geojsonPath)

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
