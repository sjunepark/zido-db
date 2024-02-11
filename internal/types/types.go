package types

import "os"

type ShpFile struct {
	absPath string
	relPath string
	name    string
}

func NewShpFile(relPath string, name string) ShpFile {
	absPath := getAbsPath(relPath)
	return ShpFile{absPath: absPath, relPath: relPath, name: name}
}

func (s ShpFile) AbsPath() string {
	return s.absPath
}

func (s ShpFile) RelPath() string {
	return s.relPath
}

func (s ShpFile) Name() string {
	return s.name
}

func getAbsPath(relPath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd + "/" + relPath
}
