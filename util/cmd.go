package util

import (
	"fmt"
	"os"
)

func CloneCommand(url string, output string) []string {
	return []string{"git", "clone", "--filter=blob:none", url, output}
}

func install(src string, dst string, perm os.FileMode) []string {
	return []string{"install", "-m", perm.String(), src, dst}
}

func PackCommand(name string, path string, prefix string) ([]string, error) {
	var cmd []string
	var err error
	if len(path) == 0 || len(prefix) == 0 {
		return nil, fmt.Errorf("pack struct has missing properties")
	}
	src := PathRepoFile(name, path)
	dst := PathPackFile(prefix, path)

	if _, err = os.Stat(src); err != nil {
		return nil, err
	}

	/* if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	} */

	/* source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close() */

	if _, err = Mkdir(PathPackDir(prefix)); err != nil {
		return nil, err
	}

	var destination *os.File
	if destination, err = os.Create(dst); err != nil {
		return nil, err
	}
	defer destination.Close()

	if prefix == "bin" {
		cmd = install(src, dst, 0755)
	} else {
		cmd = install(src, dst, 0644)
	}
	return cmd, nil
}
