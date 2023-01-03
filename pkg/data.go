package pkg

import (
	"bufio"
	"fmt"
	"os"

	"github.com/crispybaccoon/hayashi/util"
)

func GetPkg(name string) (Pkg, error) {
	path, err := util.PkgSearch(name)
	if err != nil {
		return Pkg{}, err
	}

	fh, err := os.Open(path)
	if err != nil {
		return Pkg{}, err
	}

	var pkg Pkg
	// fmt.Println("[ " + path[len(path)-4:] + " file ]")
	if path[len(path)-4:] == ".ini" {
		scanner := bufio.NewScanner(fh)
		pkg.IniFromString(*scanner)
	} else {
		pkg.FromString(fh)
	}

	// infer information
	pkg.InferInfo()

	return pkg, nil
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

	var config Config
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
