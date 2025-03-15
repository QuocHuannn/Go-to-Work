package main

import (
	"github.com/QuocHuannn/Go-to-Work/internal/config"
	"github.com/QuocHuannn/Go-to-Work/internal/initalize"
)

func main() {
	// Load configuration from environment variables and .env files
	config.LoadConfig()

	// Initialize and run application
	initalize.Run()
}
