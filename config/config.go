package config

import (
	"log"

	"github.com/pelletier/go-toml"
)

var cfg *toml.TomlTree

func init() {
	var err error
	cfg, err = toml.LoadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
}

// Get value from toml tree by key, or default nil
func Get(key string) interface{} {
	return cfg.GetDefault(key, nil)
}

// GetString string from toml tree by key, or default nil
func GetString(key string) string {
	return cfg.GetDefault(key, nil).(string)
}
