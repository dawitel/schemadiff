package config

import (
	"flag"
	"os"
	"path/filepath"
)

type Config struct {
	DevSchemaPath  string
	ProdSchemaPath string
	OutputDir      string
	Help           bool
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	
	// Get working directory for default paths
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Set up command line flags with defaults
	flag.StringVar(&cfg.DevSchemaPath, "dev", filepath.Join(wd, "schemas", "dev"), "Path to development schema directory")
	flag.StringVar(&cfg.ProdSchemaPath, "prod", filepath.Join(wd, "schemas", "prod"), "Path to production schema directory")
	flag.StringVar(&cfg.OutputDir, "output", filepath.Join(wd, "migrations"), "Output directory for migration files")
	flag.BoolVar(&cfg.Help, "help", false, "Show help message")

	flag.Parse()

	return cfg, nil
}