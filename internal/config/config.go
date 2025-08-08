
package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port    int
	GinMode string
}

func Load() Config {
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		port = 8080
	}
	return Config{
		Port:    port,
		GinMode: getEnv("GIN_MODE", "release"),
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}
