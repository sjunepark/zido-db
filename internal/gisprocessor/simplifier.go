package gisprocessor

import (
	"errors"
	"fmt"
	"github.com/sjunepark/go-gis/internal/fileprocessor"
	"os"
	"os/exec"
	"path"
	"strings"
)

func SimplifyTopojson(topojsonPath string, planarValue float64) (outputPath string, err error) {
	formattedPlanarValue := fmt.Sprintf("%.1g", planarValue)

	outputPath = path.Dir(topojsonPath)
	if !strings.HasSuffix(topojsonPath, ".topo.json") {
		return "", errors.New("input file must end with '.topo.json'")
	}
	outputName := strings.TrimSuffix(path.Base(topojsonPath), ".topo.json")
	outputName = fmt.Sprintf("%s_%s.topo.json", outputName, formattedPlanarValue)
	outputPath = path.Join(outputPath, outputName)

	err = fileprocessor.RemoveFileIfExists(outputPath)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("toposimplify", topojsonPath, "-p", fmt.Sprintf("%f", planarValue), "-o", outputPath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return outputPath, nil
}

func SimplifyGeojson(geojsonPath string, dp float64) (outputPath string, err error) {
	// format as dp as example: 20%
	formattedDp := fmt.Sprintf("%.0f%%", dp*100)

	outputPath = path.Dir(geojsonPath)
	if !strings.HasSuffix(geojsonPath, ".geo.json") {
		return "", errors.New("input file must end with '.geo.json'")
	}
	outputName := strings.TrimSuffix(path.Base(geojsonPath), ".geo.json")
	outputName = fmt.Sprintf("%s_%s.geo.json", outputName, formattedDp)
	outputPath = path.Join(outputPath, outputName)

	err = fileprocessor.RemoveFileIfExists(outputPath)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("mapshaper", geojsonPath, "-simplify", fmt.Sprintf("%f", dp), "-o", outputPath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
