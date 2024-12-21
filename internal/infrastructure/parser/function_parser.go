package parser

import (
	"github.com/dawitel/schemadiff/internal/core/domain"
	"github.com/xwb1989/sqlparser"
)

type FunctionParser struct{}

func NewFunctionParser() *FunctionParser {
	return &FunctionParser{}
}

func (p *FunctionParser) ParseCreateFunction(ddl *sqlparser.DDL) (*domain.Function, error) {
	function := &domain.Function{
		Name:       ddl.Function.Name.String(),
		ReturnType: ddl.Function.ReturnType.String(),
		Body:       ddl.Function.Body.String(),
	}

	for _, arg := range ddl.Function.Args {
		function.Arguments = append(function.Arguments, domain.Argument{
			Name: arg.Name.String(),
			Type: arg.Type.String(),
		})
	}

	return function, nil
}
