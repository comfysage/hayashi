package pkg

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DeepClone bool `yaml:"deep_clone"`

	Installed []string `yaml:"installed,flow"`
}

func DefaultConfig() Config {
	return Config{
		DeepClone: false,
	}
}

func (c Config) String() (string, error) {
	d, err := yaml.Marshal(c)
	return string(d), err
}

func (c *Config) FromString(r io.Reader) error {
	d := yaml.NewDecoder(r)
	return d.Decode(c)
}

func (c *Config) AddInstalled(pkg Pkg) error {
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
	SaveConfig(*c)

	return nil
}
