package util

import (
	"os"
	"path/filepath"
)

func GetCwd() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	wdp, err := filepath.EvalSymlinks(wd)
	if err != nil {
		return "", err
	}
	return wdp, nil
}
}
