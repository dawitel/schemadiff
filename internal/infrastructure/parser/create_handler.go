package parser

import (
	"github.com/dawitel/schemadiff/internal/core/domain"
	"github.com/xwb1989/sqlparser"
)

func (p *SQLParser) handleCreate(ddl *sqlparser.DDL, schema *domain.Schema) error {
	switch ddl.TableSpec.Type {
	case "table":
		return p.handleCreateTable(ddl, schema)
	case "sequence":
		return p.handleCreateSequence(ddl, schema)
	case "function":
		return p.handleCreateFunction(ddl, schema)
	}
	return nil
}

func (p *SQLParser) handleCreateTable(ddl *sqlparser.DDL, schema *domain.Schema) error {
	table := domain.Table{
		Name:    ddl.Table.Name.String(),
		Columns: make(map[string]domain.Column),
	}

	for _, col := range ddl.TableSpec.Columns {
		column := domain.Column{
			Name:     col.Name.String(),
			Type:     col.Type.Type,
			Nullable: !col.Type.NotNull,
		}

		if col.Type.Default != nil {
			defaultVal := col.Type.Default.String()
			column.Default = &defaultVal
		}

		table.Columns[column.Name] = column
	}

	schema.Tables[table.Name] = table
	return nil
}

func (p *SQLParser) handleCreateSequence(ddl *sqlparser.DDL, schema *domain.Schema) error {
	// Implementation for creating sequences
	return nil
}

func (p *SQLParser) handleCreateFunction(ddl *sqlparser.DDL, schema *domain.Schema) error {
	// Implementation for creating functions
	return nil
}
