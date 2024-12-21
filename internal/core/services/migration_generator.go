package services

import (
	"fmt"
	"strings"

	"github.com/dawitel/schemadiff/internal/core/domain"
)

type MigrationGenerator struct{}

func NewMigrationGenerator() *MigrationGenerator {
	return &MigrationGenerator{}
}

func (g *MigrationGenerator) Generate(diff *domain.SchemaDiff) string {
	var sb strings.Builder

	sb.WriteString("-- Generated migration\n\n")
	sb.WriteString("BEGIN;\n\n")

	// Generate table changes
	g.generateTableChanges(&sb, diff)

	// Generate sequence changes
	g.generateSequenceChanges(&sb, diff)

	// Generate function changes
	g.generateFunctionChanges(&sb, diff)

	sb.WriteString("\nCOMMIT;\n")

	return sb.String()
}

func (g *MigrationGenerator) generateTableChanges(sb *strings.Builder, diff *domain.SchemaDiff) {
	// Generate CREATE TABLE statements
	for _, table := range diff.AddedTables {
		g.generateCreateTable(sb, table)
	}

	// Generate ALTER TABLE statements for modified tables
	for _, table := range diff.ModifiedTables {
		g.generateAlterTable(sb, table)
	}

	// Generate DROP TABLE statements
	for _, tableName := range diff.RemovedTables {
		sb.WriteString(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE;\n", tableName))
	}
}

func (g *MigrationGenerator) generateCreateTable(sb *strings.Builder, table domain.TableDiff) {
	sb.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", table.Name))

	// Add columns
	var columns []string
	for _, col := range table.AddedColumns {
		columns = append(columns, g.formatColumn(col))
	}

	sb.WriteString(strings.Join(columns, ",\n"))
	sb.WriteString("\n);\n\n")
}

func (g *MigrationGenerator) generateAlterTable(sb *strings.Builder, table domain.TableDiff) {
	// Add columns
	for _, col := range table.AddedColumns {
		sb.WriteString(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s;\n",
			table.Name, g.formatColumn(col)))
	}

	// Modify columns
	for _, col := range table.ModifiedColumns {
		sb.WriteString(fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s TYPE %s;\n",
			table.Name, col.Name, col.NewType))

		if col.NewNullable != col.OldNullable {
			constraint := "SET NOT NULL"
			if col.NewNullable {
				constraint = "DROP NOT NULL"
			}
			sb.WriteString(fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s %s;\n",
				table.Name, col.Name, constraint))
		}
	}

	// Drop columns
	for _, colName := range table.RemovedColumns {
		sb.WriteString(fmt.Sprintf("ALTER TABLE %s DROP COLUMN IF EXISTS %s;\n",
			table.Name, colName))
	}

	sb.WriteString("\n")
}

func (g *MigrationGenerator) formatColumn(col domain.Column) string {
	parts := []string{col.Name, col.Type}

	if !col.Nullable {
		parts = append(parts, "NOT NULL")
	}

	if col.Default != nil {
		parts = append(parts, fmt.Sprintf("DEFAULT %s", *col.Default))
	}

	return strings.Join(parts, " ")
}

func (g *MigrationGenerator) generateSequenceChanges(sb *strings.Builder, diff *domain.SchemaDiff) {
	// Implementation for sequence changes
}

func (g *MigrationGenerator) generateFunctionChanges(sb *strings.Builder, diff *domain.SchemaDiff) {
	// Implementation for function changes
}
