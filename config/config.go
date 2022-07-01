package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiKey   string `yaml:"apiKey"`
	OuputDir string `yaml:"outputDir"`
}

var conf *Config
var confOnce sync.Once

func DefaultConfig() *Config {
	confOnce.Do(func() {
		conf = &Config{
			ApiKey:   "example",
			OuputDir: "../output/",
		}
	})
	return conf
}

func (c *Config) WithApiKey(apiKey string) *Config {
	c.ApiKey = apiKey
	return c
}

func (c *Config) WithOutputDir(outputDir string) *Config {
	c.OuputDir = outputDir
	return c
}

func LoadConfig(configFile string) {
	f, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		panic(err)
	}
}
