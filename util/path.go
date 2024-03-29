package util

import (
	"fmt"
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

func PathDetermine(path string) (string, bool, error) {
	is_dir := false
	p := ""

	if path[0] == '.' || path[0] == '/' {
		// path is either relative or absolute
		p, err := ExtendPath(path)
		if err != nil {
			return p, is_dir, err
		}

		// path points to either a file or a directory
		ph, err := os.Stat(p)
		if err != nil {
			return p, is_dir, err
		}
		if ph.IsDir() {
			is_dir = true
		}

		return p, is_dir, nil
	}

	// path is a name
	p, err := PkgSearch(path)
	if err != nil {
		return p, is_dir, err
	}

	return p, is_dir, nil
}

func PkgSearch(name string) (string, error) {
	if PkgExists("core", name) {
		return PathPkg("core", name), nil
	}

	p := filepath.Join(PathCl("core"), name+".ini")
	if PathExists(p) {
		return p, nil
	}

	fd, err := os.ReadDir(PKG_ROOT)
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

func PathRepoPkg(path string) string {
	if len(path) == 0 {
		panic("argument for path was nil")
	}
	return filepath.Join(path, REPO_PKG)
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
