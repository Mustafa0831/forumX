package server

import (
	"forum/internal/database"
)

// Config ...
type Config struct {
	WebPort  string
	Database *database.Config
}

// NewConfig generates configurations for the Server
func NewConfig(port *string) *Config {
	return &Config{
		WebPort:  *port,
		Database: database.NewConfig(),
	}
}
