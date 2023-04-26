package config

import (
	"os"
)

// Config godoc
type Config struct {
	VERSION float32
	MYSQL   string
}

// GetConfig godoc
func GetConfig() Config {
	return Config{
		VERSION: 0.1,
		MYSQL:   os.Getenv("MYSQL"),
	}
}
