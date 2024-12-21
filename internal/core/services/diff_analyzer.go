package services

import (
	"github.com/dawitel/schemadiff/internal/core/domain"
)

type DiffAnalyzer struct{}

func NewDiffAnalyzer() *DiffAnalyzer {
	return &DiffAnalyzer{}
}

func (a *DiffAnalyzer) AnalyzeDiff(dev, prod *domain.Schema) *domain.SchemaDiff {
	diff := &domain.SchemaDiff{}

	a.analyzeTables(dev, prod, diff)
	a.analyzeSequences(dev, prod, diff)
	a.analyzeFunctions(dev, prod, diff)

	return diff
}

func (a *DiffAnalyzer) analyzeTables(dev, prod *domain.Schema, diff *domain.SchemaDiff) {
	// Find added tables
	for name, table := range dev.Tables {
		if _, exists := prod.Tables[name]; !exists {
			diff.AddedTables = append(diff.AddedTables, domain.TableDiff{
				Name:         name,
				AddedColumns: a.getTableColumns(table),
			})
		}
	}

	// Find removed tables
	for name := range prod.Tables {
		if _, exists := dev.Tables[name]; !exists {
			diff.RemovedTables = append(diff.RemovedTables, name)
		}
	}

	// Find modified tables
	for name, devTable := range dev.Tables {
		if prodTable, exists := prod.Tables[name]; exists {
			if tableDiff := a.compareTable(devTable, prodTable); tableDiff != nil {
				diff.ModifiedTables = append(diff.ModifiedTables, *tableDiff)
			}
		}
	}
}

func (a *DiffAnalyzer) getTableColumns(table domain.Table) []domain.Column {
	columns := make([]domain.Column, 0, len(table.Columns))
	for _, col := range table.Columns {
		columns = append(columns, col)
	}
	return columns
}

func (a *DiffAnalyzer) compareTable(dev, prod domain.Table) *domain.TableDiff {
	diff := &domain.TableDiff{Name: dev.Name}

	// Compare columns
	for name, devCol := range dev.Columns {
		if prodCol, exists := prod.Columns[name]; exists {
			if columnDiff := a.compareColumn(devCol, prodCol); columnDiff != nil {
				diff.ModifiedColumns = append(diff.ModifiedColumns, *columnDiff)
			}
		} else {
			diff.AddedColumns = append(diff.AddedColumns, devCol)
		}
	}

	for name := range prod.Columns {
		if _, exists := dev.Columns[name]; !exists {
			diff.RemovedColumns = append(diff.RemovedColumns, name)
		}
	}

	if len(diff.AddedColumns) == 0 && len(diff.RemovedColumns) == 0 && len(diff.ModifiedColumns) == 0 {
		return nil
	}
	return diff
}

func (a *DiffAnalyzer) compareColumn(dev, prod domain.Column) *domain.ColumnDiff {
	if dev.Type != prod.Type || dev.Nullable != prod.Nullable ||
		!a.compareStringPtr(dev.Default, prod.Default) {
		return &domain.ColumnDiff{
			Name:        dev.Name,
			OldType:     prod.Type,
			NewType:     dev.Type,
			OldNullable: prod.Nullable,
			NewNullable: dev.Nullable,
			OldDefault:  prod.Default,
			NewDefault:  dev.Default,
		}
	}
	return nil
}

func (a *DiffAnalyzer) compareStringPtr(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
