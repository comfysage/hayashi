package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func pathCl(name string) string {
	return PKG_ROOT + "/" + name
}

func pkgName(name string) string {
	return name + ".yaml"
}

func PathPkg(cl string, name string) string {
	return pathCl(cl) + "/" + pkgName(name)
}

func PathExists(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

func PkgExists(cl string, name string) bool {
	return PathExists(PathPkg(cl, name))
}

func PkgSearch(name string) (string, error) {
	if PkgExists("core", name) {
		return PathPkg("core", name), nil
	}
	if PathExists(pathCl("core") + "/" + name + ".ini") {
		return pathCl("core") + "/" + name + ".ini", nil
	}

	fd, err := ioutil.ReadDir(PKG_ROOT)
	for _, d := range(fd) {
		if !d.IsDir() {
			continue
		}
		if PkgExists(d.Name(), name) {
			return PathPkg(d.Name(), name), nil
		}
	}
	if err != nil {
		return "", err
	}

	return "", fmt.Errorf("pkg not found.")
}

func PathRepo(name string) string {
	return REPO_ROOT + "/" + name
}
