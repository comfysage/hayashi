package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func PathConfig() string {
	return filepath.Join(HAYASHI_ROOT, CONFIG_NAME)
}

func PathStoreFile() string {
	return filepath.Join(HAYASHI_ROOT, STORE_NAME)
}

func PathCl(name string) string {
	if len(name) == 0 {
		panic("argument for collection name was nil")
	}
	return filepath.Join(PKG_ROOT, name)
}

func pkgName(name string) string {
	if len(name) == 0 {
		panic("argument for pkg name was nil")
	}
	return name + ".yaml"
}

func RemoveExtension(name string) string {
	if len(name) == 0 {
		panic("argument for pkg name was nil")
	}
	lastdot := 0
	for i, s := range name {
		if s == '.' {
			lastdot = i
		}
	}
	return name[:lastdot]
}

func PathPkg(collection string, name string) string {
	cl := PathCl(collection)
	pkg_name := pkgName(name)
	return filepath.Join(cl, pkg_name)
}

func PathExists(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}

func PkgExists(cl string, name string) bool {
	p := PathPkg(cl, name)
	return PathExists(p)
}

func PkgSearch(name string) (string, error) {
	if PkgExists("core", name) {
		return PathPkg("core", name), nil
	}

	p := filepath.Join(PathCl("core"), name+".ini")
	if PathExists(p) {
		return p, nil
	}

	fd, err := ioutil.ReadDir(PKG_ROOT)
	for _, d := range fd {
		if !d.IsDir() {
			continue
		}
		clName := d.Name()

		if PkgExists(clName, name) {
			return PathPkg(clName, name), nil
		}
	}
	if err != nil {
		return "", err
	}

	return "", fmt.Errorf("pkg not found.")
}

func PathRepo(name string) string {
	if len(name) == 0 {
		panic("argument for repo name was nil")
	}
	return filepath.Join(REPO_ROOT, name)
}

func PathRepoFile(name string, path string) string {
	if len(path) == 0 {
		panic("argument for file name was nil")
	}
	return filepath.Join(PathRepo(name), path)
}

func PathPackDir(dir string) string {
	if len(dir) == 0 {
		panic("argument for type dir was nil")
	}
	return filepath.Join(PACK_ROOT, dir)
}

func PathPackFile(dir string, path string) string {
	if len(path) == 0 {
		panic("argument for path was nil")
	}
	file := filepath.Base(path)
	if file == "." {
		panic("argument for file was invalid")
	}
	return filepath.Join(PathPackDir(dir), file)
}
