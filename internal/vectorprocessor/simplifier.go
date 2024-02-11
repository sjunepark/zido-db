package vectorprocessor

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

	cmd := exec.Command("toposimplify", topojsonPath, "-p", formattedPlanarValue, "-o", outputPath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
