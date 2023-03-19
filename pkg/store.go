package pkg

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type StoreFile struct {
	Installed []string `yaml:"installed"`
}

func DefaultStoreFile() StoreFile {
	return StoreFile{
		Installed: []string{},
	}
}

func (c StoreFile) String() (string, error) {
	d, err := yaml.Marshal(c)
	return string(d), err
}

func (c *StoreFile) FromString(r io.Reader) error {
	d := yaml.NewDecoder(r)
	return d.Decode(c)
}

func (c *StoreFile) AddInstalled(pkg Pkg) error {
	if pkg.Collection == "" {
		return fmt.Errorf("no collection defined for pkg")
	}
	if pkg.Name == "" {
		return fmt.Errorf("no name defined for pkg")
	}

	pkgString := pkg.Collection + "/" + pkg.Name
	for _, s := range c.Installed {
		if s == pkgString {
			return nil
		}
	}
	c.Installed = append(c.Installed, pkgString)
	SaveStoreFile(*c)

	return nil
}
