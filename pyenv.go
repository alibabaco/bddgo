package bddgo

import (
	"os"
	"path"
)

const venvDirectory string = ".bddenv"

type VEnv struct {
	Path string
}

func GetVenvRoot(directory string) string {
	return path.Join(directory, venvDirectory)
}

func Exists(directory string) bool {
	_, err := os.Stat(GetVenvRoot(directory))
	return os.IsNotExist(err)
}
