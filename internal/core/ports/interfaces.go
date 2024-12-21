package ports

import "github.com/dawitel/schemadiff/internal/core/domain"

type SchemaParser interface {
	ParseDirectory(path string) (*domain.Schema, error)
}

type SchemaStorage interface {
	SaveMigration(path string, content string) error
	EnsureDirectory(path string) error
}

type SchemaDiffService interface {
	Execute() error
}
