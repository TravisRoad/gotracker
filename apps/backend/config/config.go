package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var Conf *Config

type Config struct {
	Port          string `yaml:"port"`
	DbPath        string `yaml:"dbPath"`
	PolicyPath    string `yaml:"policyPath"`
	TokenLifeSpan int    `yaml:"tokenLifespan"`
	JwtSecret     string `yaml:"jwtSecret"`
}

func ConfigFromFile(path string) (*Config, error) {
	conf := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&conf); err != nil {
		return nil, err
	}

	conf.setDefaults()
	return conf, nil
}

func (c *Config) setDefaults() {
	if c.Port == "" {
		c.Port = "8080"
	}
	if c.PolicyPath == "" {
		c.PolicyPath = "./policy.csv"
	}
	if c.TokenLifeSpan == 0 {
		// c.TokenLifeSpan = 3_600_000 // us 1 hour
		c.TokenLifeSpan = 6_000
	}
	if c.JwtSecret == "" {
		c.JwtSecret = "secret" // default
	}

}
