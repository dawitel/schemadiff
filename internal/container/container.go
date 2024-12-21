package container

import (
	"github.com/dawitel/schemadiff/internal/config"
	"github.com/dawitel/schemadiff/internal/core/ports"
	"github.com/dawitel/schemadiff/internal/core/services"
	"github.com/dawitel/schemadiff/internal/infrastructure/parser"
	"github.com/dawitel/schemadiff/internal/infrastructure/storage"
)

type Container struct {
	SchemaDiffService ports.SchemaDiffService
}

func NewContainer(cfg *config.Config) *Container {
	// Initialize infrastructure
	schemaParser := parser.NewSQLParser()
	fileStorage := storage.NewFileStorage()

	// Initialize services
	schemaDiffService := services.NewSchemaDiffService(
		cfg,
		schemaParser,
		fileStorage,
	)

	return &Container{
		SchemaDiffService: schemaDiffService,
	}
}
