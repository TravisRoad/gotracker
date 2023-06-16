package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port          string `yaml:"port"`
	DbPath        string `yaml:"dbPath"`
	PolicyPath    string `yaml:"policyPath"`
	TokenLifeSpan int    `yaml:"tokenLifespan"`
	JwtSecret     string `yaml:"jwtSecret"`
	TmdbKey       string `yaml:"tmdbKey"`
}

func ConfigFromFile(configPath string) (*Config, error) {
	conf := &Config{}

	file, err := os.Open(configPath)
	if err == nil {
		defer file.Close()

		d := yaml.NewDecoder(file)

		if err := d.Decode(&conf); err != nil {
			return nil, err
		}
	} else {
		log.Println("failed to open config file: ", err.Error())
		log.Println("use default config")
	}
	log.Println(conf.TmdbKey)
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
		c.TokenLifeSpan = 3_600_000 // us 1 hour
	}
	if c.JwtSecret == "" {
		c.JwtSecret = "secret" // default
	}
}
