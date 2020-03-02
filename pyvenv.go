package bddgo

import (
	"fmt"
	"os"
	"path"
)

const venvDirectory string = ".bddenv"

type VEnv struct {
	Path   string
	Python string
}

func GetVenvRoot(directory string) string {
	return path.Join(directory, venvDirectory)
}

func VirtualEnv(directory string, python string) *VEnv {
	return &VEnv{
		GetVenvRoot(directory),
		python,
	}
}

func (this *VEnv) Exists() bool {
	_, err := os.Stat(this.Path)
	return os.IsNotExist(err)
}

func (this *VEnv) Create() error {
	return fmt.Errorf("Not implemented")
}

func (this *VEnv) Delete() error {
	return fmt.Errorf("Not implemented")
}
