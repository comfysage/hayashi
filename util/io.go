package util

import "os"

// mkdir `path` if directory does not already exists.
// if path is created return true else return false
func Mkdir(path string) (bool, error) {
	exists := PathExists(path)
	if exists {
		return false, nil
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return false, err
	}
	return true, nil
}

