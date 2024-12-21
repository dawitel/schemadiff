package services

import (
	"fmt"
	"path/filepath"

	"github.com/dawitel/schemadiff/internal/config"
	"github.com/dawitel/schemadiff/internal/core/ports"
)

type SchemaDiffService struct {
	cfg              *config.Config
	schemaParser     ports.SchemaParser
	storage          ports.SchemaStorage
	diffAnalyzer     *DiffAnalyzer
	migrationGenerator *MigrationGenerator
}

func NewSchemaDiffService(
	cfg *config.Config,
	schemaParser ports.SchemaParser,
	storage ports.SchemaStorage,
) *SchemaDiffService {
	return &SchemaDiffService{
		cfg:                cfg,
		schemaParser:       schemaParser,
		storage:            storage,
		diffAnalyzer:       NewDiffAnalyzer(),
		migrationGenerator: NewMigrationGenerator(),
	}
}

func (s *SchemaDiffService) Execute() error {
	if s.cfg.Help {
		s.printHelp()
		return nil
	}

	// Ensure output directory exists
	if err := s.storage.EnsureDirectory(s.cfg.OutputDir); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Parse schemas
	devSchema, err := s.schemaParser.ParseDirectory(s.cfg.DevSchemaPath)
	if err != nil {
		return fmt.Errorf("failed to parse dev schema: %w", err)
	}

	prodSchema, err := s.schemaParser.ParseDirectory(s.cfg.ProdSchemaPath)
	if err != nil {
		return fmt.Errorf("failed to parse prod schema: %w", err)
	}

	// Generate diff
	diff := s.diffAnalyzer.AnalyzeDiff(devSchema, prodSchema)

	// Generate migration
	migration := s.migrationGenerator.Generate(diff)

	// Save migration
	outputPath := filepath.Join(s.cfg.OutputDir, "schema_migration.sql")
	if err := s.storage.SaveMigration(outputPath, migration); err != nil {
		return fmt.Errorf("failed to save migration: %w", err)
	}

	fmt.Printf("Migration file generated successfully at: %s\n", outputPath)
	return nil
}

func (s *SchemaDiffService) printHelp() {
	fmt.Println(`SQL Schema Diff Tool

Usage:
  schemadiff [options]

Options:
  -dev string    Path to development schema directory (default: "./schemas/dev")
  -prod string   Path to production schema directory (default: "./schemas/prod")
  -output string Output directory for migration files (default: "./migrations")
  -help         Show this help message

Example:
  schemadiff -dev ./schemas/dev -prod ./schemas/prod -output ./migrations`)
}
