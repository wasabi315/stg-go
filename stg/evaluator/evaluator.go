package evaluator

import (
	"github.com/wasabi315/stg-go/stg/ast"
)

// Evaluator STG program evaluator
type Evaluator interface {
	Eval(program []*ast.Bind)
}
