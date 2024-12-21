package main

import (
	"log"

	"github.com/dawitel/schemadiff/internal/config"
	"github.com/dawitel/schemadiff/internal/container"
)

func main() {
	// Initialize configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	// Initialize dependency container
	c := container.NewContainer(cfg)

	// Run the application
	if err := c.SchemaDiffService.Execute(); err != nil {
		log.Fatalf("Failed to execute schema diff: %v", err)
	}
}
