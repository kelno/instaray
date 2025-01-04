// Package main is the entry point of the application. It initializes the
// configuration, sets up logging, and starts the Instaray bot.
package main

import (
	"fmt"

	"github.com/Madh93/instaray/internal/config"
	"github.com/Madh93/instaray/internal/instaray"
	"github.com/Madh93/instaray/internal/logging"
)

// main initializes the configuration, sets up logging, and starts the
// Instaray bot.
func main() {
	// Load configuration
	config := config.New()

	// Setup logger
	logger := logging.New(&config.Logging)
	if config.Path != "" {
		logger.Debug(fmt.Sprintf("Loaded configuration from %s", config.Path))
	}

	// Setup instaray
	instaray := instaray.New(logger, &instaray.Config{
		Telegram: &config.Telegram,
	})

	// Let's go
	logger.Info("3, 2, 1...  Launching Instaray... 🚀")
	if err := instaray.Run(); err != nil {
		logger.Fatal("💥 Something went wrong.", "error", err)
	}
}
