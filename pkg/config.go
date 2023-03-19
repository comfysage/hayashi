package pkg

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DeepClone bool `yaml:"deep_clone"`
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
