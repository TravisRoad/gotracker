package config

type Config struct {
	Port    string   `yaml:"port"`
	Db      Database `yaml:"db"`
	Redis   Redis    `yaml:"redis"`
	TmdbKey string   `yaml:"tmdbkey"`
}
