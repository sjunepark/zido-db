package fileprocessor

import (
	"errors"
	"github.com/sjunepark/go-gis/internal/types"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type ShpFile = types.ShpFile

func CollectShpFiles(dir string) ([]ShpFile, error) {
	var files []ShpFile
	err := filepath.WalkDir(dir, func(relPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(relPath, ".shp") {
			name := strings.TrimSuffix(d.Name(), ".shp")
			files = append(files, types.NewShpFile(relPath, name))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func GetOutputDirectory(inputPath string) string {
	inputDir := filepath.Dir(inputPath)
	outputDir := strings.Replace(inputDir, "data/input", "data/output", 1)
	//	Create output directory if not exists
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return ""
		}
	}
	return outputDir
}

func GetOutputPath(inputPath string, outputExt string) (string, error) {
	outputDir := GetOutputDirectory(inputPath)
	filenameWithExt := filepath.Base(inputPath)
	filename := strings.TrimSuffix(filenameWithExt, filepath.Ext(inputPath))

	if !strings.HasPrefix(outputExt, ".") {
		return "", errors.New("outputExt must start with '.'")
	}

	return filepath.Join(outputDir, filename+outputExt), nil
}

func RemoveFileIfExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		err := os.Remove(path)
		println("Removed existing file: " + path)
		if err != nil {
			return err
		}
	}
	return nil
}
