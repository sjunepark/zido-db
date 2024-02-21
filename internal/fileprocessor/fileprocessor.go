package fileprocessor

import (
	"errors"
	"github.com/sjunepark/go-gis/internal/types"
	"io/fs"
	"log"
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

func GetOutputDirectory(inputPath string) (string, error) {
	inputDir := filepath.Dir(inputPath)
	outputDir := strings.Replace(inputDir, "data/input", "data/output", 1)
	//	Create output directory if not exists
	err := CreateDirIfNotExists(outputDir)
	if err != nil {
		return "", err
	}
	return outputDir, nil
}

func GetOutputPath(inputPath string, outputExt string) (string, error) {
	outputDir, err := GetOutputDirectory(inputPath)
	if err != nil {
		return "", err

	}
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

func GetFilesWithExt(dirPath, ext string) (filePaths []string, err error) {
	if !strings.HasPrefix(ext, ".") {
		return nil, errors.New("ext must start with '.'")
	}

	err = filepath.WalkDir(dirPath, func(relPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(relPath, ext) {
			filePaths = append(filePaths, relPath)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filePaths, nil
}

func CreateDirIfNotExists(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveDir(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return err
	}
	log.Printf("Removed directory: %s\n", dir)
	return nil
}
