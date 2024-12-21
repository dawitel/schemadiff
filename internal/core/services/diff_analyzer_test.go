package services

import (
	"reflect"
	"testing"

	"github.com/dawitel/schemadiff/internal/core/domain"
)

func TestAnalyzeDiff(t *testing.T) {
	tests := []struct {
		name string
		dev  *domain.Schema
		prod *domain.Schema
		want *domain.SchemaDiff
	}{
		{
			name: "detect added table",
			dev: &domain.Schema{
				Tables: map[string]domain.Table{
					"users": {
						Name: "users",
						Columns: map[string]domain.Column{
							"id": {Name: "id", Type: "SERIAL"},
						},
					},
				},
			},
			prod: &domain.Schema{
				Tables: map[string]domain.Table{},
			},
			want: &domain.SchemaDiff{
				AddedTables: []domain.TableDiff{
					{
						Name: "users",
						AddedColumns: []domain.Column{
							{Name: "id", Type: "SERIAL"},
						},
					},
				},
			},
		},
		{
			name: "detect modified table",
			dev: &domain.Schema{
				Tables: map[string]domain.Table{
					"users": {
						Name: "users",
						Columns: map[string]domain.Column{
							"id":    {Name: "id", Type: "SERIAL"},
							"email": {Name: "email", Type: "VARCHAR(255)"},
						},
					},
				},
			},
			prod: &domain.Schema{
				Tables: map[string]domain.Table{
					"users": {
						Name: "users",
						Columns: map[string]domain.Column{
							"id": {Name: "id", Type: "SERIAL"},
						},
					},
				},
			},
			want: &domain.SchemaDiff{
				ModifiedTables: []domain.TableDiff{
					{
						Name: "users",
						AddedColumns: []domain.Column{
							{Name: "email", Type: "VARCHAR(255)"},
						},
					},
				},
			},
		},
	}

	analyzer := NewDiffAnalyzer()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := analyzer.AnalyzeDiff(tt.dev, tt.prod)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AnalyzeDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
