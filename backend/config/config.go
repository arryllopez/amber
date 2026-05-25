// Loads environment configuration for the backend.
package config

import "os"

type Config struct {
	Addr string
}

func Load() Config {
	addr := os.Getenv("AMBER_HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	return Config{
		Addr: addr,
	}
}
