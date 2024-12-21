package services

import (
	"strings"
	"testing"

	"github.com/dawitel/schemadiff/internal/core/domain"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
		diff *domain.SchemaDiff
		want string
	}{
		{
			name: "generate create table",
			diff: &domain.SchemaDiff{
				AddedTables: []domain.TableDiff{
					{
						Name: "users",
						AddedColumns: []domain.Column{
							{Name: "id", Type: "SERIAL", Nullable: false},
							{Name: "name", Type: "VARCHAR(255)", Nullable: false},
						},
					},
				},
			},
			want: strings.TrimSpace(`
BEGIN;

CREATE TABLE users (
id SERIAL NOT NULL,
name VARCHAR(255) NOT NULL
);

COMMIT;
`),
		},
		{
			name: "generate alter table",
			diff: &domain.SchemaDiff{
				ModifiedTables: []domain.TableDiff{
					{
						Name: "users",
						AddedColumns: []domain.Column{
							{Name: "email", Type: "VARCHAR(255)", Nullable: true},
						},
						RemovedColumns: []string{"phone"},
					},
				},
			},
			want: strings.TrimSpace(`
BEGIN;

ALTER TABLE users
ADD COLUMN email VARCHAR(255),
DROP COLUMN IF EXISTS phone;

COMMIT;
`),
		},
	}

	generator := NewMigrationGenerator()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generator.Generate(tt.diff)
			if strings.TrimSpace(got) != strings.TrimSpace(tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
