package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dawitel/schemadiff/internal/core/domain"
)

func TestParseDirectory(t *testing.T) {
	tests := []struct {
		name    string
		sqlFile string
		want    *domain.Schema
		wantErr bool
	}{
		{
			name: "parse create table",
			sqlFile: `CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255) NOT NULL,
				email VARCHAR(255) UNIQUE
			);`,
			want: &domain.Schema{
				Tables: map[string]domain.Table{
					"users": {
						Name: "users",
						Columns: map[string]domain.Column{
							"id": {
								Name:     "id",
								Type:     "SERIAL",
								Nullable: false,
							},
							"name": {
								Name:     "name",
								Type:     "VARCHAR(255)",
								Nullable: false,
							},
							"email": {
								Name:     "email",
								Type:     "VARCHAR(255)",
								Nullable: true,
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "parse create sequence",
			sqlFile: `CREATE SEQUENCE user_id_seq
				START WITH 1
				INCREMENT BY 1;`,
			want: &domain.Schema{
				Sequences: map[string]domain.Sequence{
					"user_id_seq": {
						Name:      "user_id_seq",
						StartWith: 1,
						Increment: 1,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary directory
			tmpDir, err := os.MkdirTemp("", "schema_test")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			// Create test SQL file
			testFile := filepath.Join(tmpDir, "test.sql")
			if err := os.WriteFile(testFile, []byte(tt.sqlFile), 0644); err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}

			parser := NewSQLParser()
			got, err := parser.ParseDirectory(tmpDir)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil {
				compareSchemas(t, got, tt.want)
			}
		})
	}
}

func compareSchemas(t *testing.T, got, want *domain.Schema) {
	// Compare tables
	if len(got.Tables) != len(want.Tables) {
		t.Errorf("Tables count mismatch: got %d, want %d", len(got.Tables), len(want.Tables))
	}

	for tableName, wantTable := range want.Tables {
		gotTable, exists := got.Tables[tableName]
		if !exists {
			t.Errorf("Missing table: %s", tableName)
			continue
		}

		compareTables(t, gotTable, wantTable)
	}

	// Compare sequences
	if len(got.Sequences) != len(want.Sequences) {
		t.Errorf("Sequences count mismatch: got %d, want %d", len(got.Sequences), len(want.Sequences))
	}

	for seqName, wantSeq := range want.Sequences {
		gotSeq, exists := got.Sequences[seqName]
		if !exists {
			t.Errorf("Missing sequence: %s", seqName)
			continue
		}

		compareSequences(t, gotSeq, wantSeq)
	}
}

func compareTables(t *testing.T, got, want domain.Table) {
	if got.Name != want.Name {
		t.Errorf("Table name mismatch: got %s, want %s", got.Name, want.Name)
	}

	if len(got.Columns) != len(want.Columns) {
		t.Errorf("Columns count mismatch for table %s: got %d, want %d", got.Name, len(got.Columns), len(want.Columns))
	}

	for colName, wantCol := range want.Columns {
		gotCol, exists := got.Columns[colName]
		if !exists {
			t.Errorf("Missing column %s in table %s", colName, got.Name)
			continue
		}

		compareColumns(t, gotCol, wantCol)
	}
}

func compareColumns(t *testing.T, got, want domain.Column) {
	if got.Name != want.Name {
		t.Errorf("Column name mismatch: got %s, want %s", got.Name, want.Name)
	}
	if got.Type != want.Type {
		t.Errorf("Column type mismatch for %s: got %s, want %s", got.Name, got.Type, want.Type)
	}
	if got.Nullable != want.Nullable {
		t.Errorf("Column nullable mismatch for %s: got %v, want %v", got.Name, got.Nullable, want.Nullable)
	}
}

func compareSequences(t *testing.T, got, want domain.Sequence) {
	if got.Name != want.Name {
		t.Errorf("Sequence name mismatch: got %s, want %s", got.Name, want.Name)
	}
	if got.StartWith != want.StartWith {
		t.Errorf("Sequence start mismatch for %s: got %d, want %d", got.Name, got.StartWith, want.StartWith)
	}
	if got.Increment != want.Increment {
		t.Errorf("Sequence increment mismatch for %s: got %d, want %d", got.Name, got.Increment, want.Increment)
	}
}
