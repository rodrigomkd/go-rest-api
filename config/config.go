package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	ServerPort int
	DataSource string
	ApiUri     string
}

func ReadConfig(configfile string) Config {
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
