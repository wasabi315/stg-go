package evaluator

import (
	"github.com/wasabi315/stg-go/ast"
)

// Evaluator STG program evaluator
type Evaluator interface {
	Eval(program ast.Program)
}
