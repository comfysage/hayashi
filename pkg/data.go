package pkg

import (
	"bufio"
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/util"
)

func GetPkgFromPath(path string) (Pkg, error) {
	fullpath, err := util.ExtendPath(path)
	if err != nil {
		return Pkg{}, err
	}
	fh, err := os.Open(fullpath)
	if err != nil {
		return Pkg{}, err
	}

	var pkg Pkg

	_, err = util.HasFileExtension(fullpath)
	if err != nil {
		return Pkg{}, err
	}

	pkgRead := 0
	isYaml, _ := util.HasCorrectFileExtension(fullpath, ".yaml")
	if isYaml {
		err = pkg.FromString(fh)
		if err != nil {
			return Pkg{}, err
		}
		pkgRead++
	}
	isIni, _ := util.HasCorrectFileExtension(fullpath, ".ini")
	if isIni {
		scanner := bufio.NewScanner(fh)
		pkg.IniFromString(*scanner)
		pkgRead++
	}

	if pkgRead < 1 {
		return Pkg{}, fmt.Errorf("no correct file extension; requires `ini` or `yaml` extension")
	}

	err = pkg.SafeGuard()
	if err != nil {
		return Pkg{}, err
	}

	// infer information
	err = pkg.InferInfo(fullpath)
	if err != nil {
		return Pkg{}, err
	}

	return pkg, nil
}

func GetPkg(name string) (Pkg, error) {
	path, err := util.PkgSearch(name)
	if err != nil {
		return Pkg{}, err
	}

	pkg, err := GetPkgFromPath(path)
	return pkg, err
}

func SavePkg(pkg Pkg) error {
	str, err := pkg.String()
	if err != nil {
		return err
	}

	fh, err := os.OpenFile(util.PathPkg("custom", pkg.Name),
		os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	defer fh.Close()

	if _, err = fh.WriteString(str); err != nil {
		panic(err)
	}

	return nil
}

func GetConfig() (Config, error) {
	bool := util.PathExists(util.PathConfig())
	if bool == false {
		return Config{}, fmt.Errorf("config does not exist. try running hayashi config init.")
	}

	fh, err := os.Open(util.PathConfig())
	if err != nil {
		return Config{}, err
	}

	var config Config = DefaultConfig()
	config.FromString(fh)
	return config, nil
}

func SaveConfig(config Config) error {
	str, err := config.String()
	if err != nil {
		return err
	}

	fh, err := os.OpenFile(util.PathConfig(),
		os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	defer fh.Close()

	if _, err = fh.WriteString(str); err != nil {
		panic(err)
	}

	return nil
}

func GetStoreFile() (StoreFile, error) {
	bool := util.PathExists(util.PathStoreFile())
	if bool == false {
		store := DefaultStoreFile()
		SaveStoreFile(store)
		return GetStoreFile()
	}

	fh, err := os.Open(util.PathStoreFile())
	if err != nil {
		return StoreFile{}, err
	}

	store := DefaultStoreFile()
	store.FromString(fh)
	return store, nil
}

func SaveStoreFile(store StoreFile) error {
	str, err := store.String()
	if err != nil {
		return err
	}

	fh, err := os.OpenFile(util.PathStoreFile(),
		os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	defer fh.Close()

	if _, err = fh.WriteString(str); err != nil {
		panic(err)
	}

	return nil
}
