package parser

import (
	"github.com/dawitel/schemadiff/internal/core/domain"
	"github.com/xwb1989/sqlparser"
)

type TableParser struct{}

func NewTableParser() *TableParser {
	return &TableParser{}
}

func (p *TableParser) ParseCreateTable(ddl *sqlparser.DDL) (*domain.Table, error) {
	table := &domain.Table{
		Name:    ddl.Table.Name.String(),
		Columns: make(map[string]domain.Column),
	}

	for _, colDef := range ddl.TableSpec.Columns {
		col := p.parseColumn(colDef)
		table.Columns[col.Name] = col
	}

	return table, nil
}

func (p *TableParser) parseColumn(colDef *sqlparser.ColumnDefinition) domain.Column {
	col := domain.Column{
		Name:     colDef.Name.String(),
		Type:     colDef.Type.Type,
		Nullable: !colDef.Type.NotNull,
	}

	if colDef.Type.Default != nil {
		defaultVal := colDef.Type.Default.String()
		col.Default = &defaultVal
	}

	return col
}
