package util

import (
	"os"
	"path/filepath"
)

func GetCwd() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	exep, err := filepath.EvalSymlinks(exe)
	if err != nil {
		return "", err
	}
	cwd := filepath.Dir(exep)
	return cwd, nil
}
