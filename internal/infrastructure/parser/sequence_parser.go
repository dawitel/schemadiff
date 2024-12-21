package parser

import (
	"github.com/xwb1989/sqlparser"
	"github.com/dawitel/schemadiff/internal/core/domain"
)

type SequenceParser struct{}

func NewSequenceParser() *SequenceParser {
	return &SequenceParser{}
}

func (p *SequenceParser) ParseCreateSequence(ddl *sqlparser.DDL) (*domain.Sequence, error) {
	sequence := &domain.Sequence{
		Name:      ddl.Sequence.Name.String(),
		StartWith: ddl.Sequence.StartWith,
		Increment: ddl.Sequence.Increment,
	}

	return sequence, nil
}
