package parser

import (
	"github.com/dawitel/schemadiff/internal/core/domain"
	"github.com/xwb1989/sqlparser"
)

func (p *SQLParser) handleAlter(ddl *sqlparser.DDL, schema *domain.Schema) error {
	table, exists := schema.Tables[ddl.Table.Name.String()]
	if !exists {
		return nil
	}

	for _, spec := range ddl.TableSpec.Columns {
		switch spec.Type.Action {
		case "add":
			p.handleAddColumn(spec, &table)
		case "modify":
			p.handleModifyColumn(spec, &table)
		case "drop":
			p.handleDropColumn(spec, &table)
		}
	}

	schema.Tables[table.Name] = table
	return nil
}

func (p *SQLParser) handleAddColumn(spec *sqlparser.ColumnDefinition, table *domain.Table) {
	column := domain.Column{
		Name:     spec.Name.String(),
		Type:     spec.Type.Type,
		Nullable: !spec.Type.NotNull,
	}

	if spec.Type.Default != nil {
		defaultVal := spec.Type.Default.String()
		column.Default = &defaultVal
	}

	table.Columns[column.Name] = column
}

func (p *SQLParser) handleModifyColumn(spec *sqlparser.ColumnDefinition, table *domain.Table) {
	// Implementation for modifying columns
}

func (p *SQLParser) handleDropColumn(spec *sqlparser.ColumnDefinition, table *domain.Table) {
	delete(table.Columns, spec.Name.String())
}
