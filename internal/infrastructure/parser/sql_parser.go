package parser

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/dawitel/schemadiff/internal/core/domain"
	"github.com/xwb1989/sqlparser"
)

type SQLParser struct{}

func NewSQLParser() *SQLParser {
	return &SQLParser{}
}

func (p *SQLParser) ParseDirectory(dir string) (*domain.Schema, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading directory: %w", err)
	}

	schema := &domain.Schema{
		Tables:      make(map[string]domain.Table),
		Sequences:   make(map[string]domain.Sequence),
		Functions:   make(map[string]domain.Function),
		Triggers:    make(map[string]domain.Trigger),
		Indexes:     make(map[string]domain.Index),
		Constraints: make(map[string]domain.Constraint),
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			content, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("reading file %s: %w", file.Name(), err)
			}

			if err := p.parseSQL(string(content), schema); err != nil {
				return nil, fmt.Errorf("parsing SQL from %s: %w", file.Name(), err)
			}
		}
	}

	return schema, nil
}

func (p *SQLParser) parseSQL(sql string, schema *domain.Schema) error {
	statements, err := sqlparser.SplitStatementToPieces(sql)
	if err != nil {
		return fmt.Errorf("splitting SQL statements: %w", err)
	}

	for _, stmt := range statements {
		parsed, err := sqlparser.Parse(stmt)
		if err != nil {
			return fmt.Errorf("parsing SQL statement: %w", err)
		}

		switch v := parsed.(type) {
		case *sqlparser.DDL:
			if err := p.handleDDL(v, schema); err != nil {
				return fmt.Errorf("handling DDL: %w", err)
			}
		}
	}

	return nil
}

func (p *SQLParser) handleDDL(ddl *sqlparser.DDL, schema *domain.Schema) error {
	switch ddl.Action {
	case "create":
		return p.handleCreate(ddl, schema)
	case "alter":
		return p.handleAlter(ddl, schema)
	}
	return nil
}
