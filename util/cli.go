package util

import (
	"fmt"
	"os"
	fpath "path"
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

func ExtendPath(path string) (string, error) {
	if len(path) < 1 {
		return "", fmt.Errorf("path not long enough")
	}

	// absolute path
	if string(path[0]) == "/" {
		return path, nil
	}

	// relative path
	cwd, err := GetCwd()
	if err != nil {
		return "", err
	}
	extended := fpath.Join(cwd, path)
	abspath, err := filepath.Abs(extended)
	if err != nil {
		return "", err
	}

	return abspath, nil
}

func HasFileExtension(path string) (string, error) {
	if len(path) < 1 {
		return "", fmt.Errorf("path not long enough")
	}

	ext := filepath.Ext(path)
	if len(ext) < 1 {
		return "", fmt.Errorf("no extension specified")
	}

	return ext, nil
}

func HasCorrectFileExtension(path string, extension string) (bool, error) {
	ext, err := HasFileExtension(path)
	if err != nil {
		return false, err
	}

	if ext != extension {
		return false, fmt.Errorf("incorrect extension")
	}

	return true, nil
}
