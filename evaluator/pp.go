package evaluator

import (
	"fmt"
	"strings"

	"github.com/wasabi315/stg-go/ast"
)

// PP (pretty printer) implements Evaluator
type PP struct {
	indentWidth uint
}

// NewPP create PP instance
func NewPP(indentWidth uint) *PP {
	return &PP{indentWidth}
}

// Eval pretty print program
func (pp *PP) Eval(program []*ast.Bind) {
	pp.printBinds(program, 0)
}

func (pp *PP) printIndent(n uint) {
	fmt.Print(strings.Repeat(" ", int(pp.indentWidth*n)))
}

func (pp *PP) printBinds(bs []*ast.Bind, indent uint) {
	for i, b := range bs {
		pp.printBind(b, indent)
		fmt.Println()
		if i != len(bs)-1 {
			fmt.Println()
		}
	}
}

func (pp *PP) printBind(b *ast.Bind, indent uint) {
	pp.printIndent(indent)
	pp.printVar(b.Var)
	fmt.Print(" = ")
	pp.printLF(b.LF, indent)
}

func (pp *PP) printLF(lf *ast.LF, indent uint) {
	pp.printVars(lf.Free)
	if lf.Upd {
		fmt.Print(" \\u ")
	} else {
		fmt.Print(" \\n ")
	}
	pp.printVars(lf.Args)
	fmt.Println(" ->")
	pp.printExpr(lf.Body, indent+1)
}

func (pp *PP) printExpr(e ast.Expr, indent uint) {
	pp.printIndent(indent)
	switch e := e.(type) {
	case *ast.Let:
		fmt.Println("let")
		pp.printBinds(e.Binds, indent+1)
		pp.printIndent(indent)
		fmt.Println("in")
		pp.printExpr(e.Body, indent+1)

	case *ast.Case:
		fmt.Print("case")
		fmt.Print(" ")
		pp.printExpr(e.Target, 0)
		fmt.Println(" of")
		pp.printAlts(e.Alts, indent+1)

	case *ast.VarApp:
		pp.printVar(e.Var)
		fmt.Print(" ")
		pp.printAtoms(e.Atoms)

	case *ast.CtorApp:
		pp.printCtor(e.Ctor)
		fmt.Print(" ")
		pp.printAtoms(e.Atoms)

	case *ast.PrimApp:
		pp.printPrim(e.Prim)
		fmt.Print(" ")
		pp.printAtoms(e.Atoms)

	case ast.Lit:
		pp.printLit(e)
	}
}

func (pp *PP) printAlts(as []ast.Alt, indent uint) {
	for i, a := range as {
		pp.printAlt(a, indent)
		fmt.Println()
		if i != len(as)-1 {
			fmt.Println()
		}
	}
}

func (pp *PP) printAlt(a ast.Alt, indent uint) {
	pp.printIndent(indent)
	switch a := a.(type) {
	case *ast.AAlt:
		pp.printCtor(a.Ctor)
		fmt.Print(" ")
		pp.printVars(a.Vars)
		fmt.Println(" ->")
		pp.printExpr(a.Expr, indent+1)

	case *ast.PAlt:
		pp.printLit(a.Lit)
		fmt.Println(" ->")
		pp.printExpr(a.Expr, indent+1)

	case *ast.VAlt:
		pp.printVar(a.Var)
		fmt.Println(" ->")
		pp.printExpr(a.Expr, indent+1)

	case *ast.DAlt:
		fmt.Println("default ->")
		pp.printExpr(a.Expr, indent+1)
	}
}

func (pp *PP) printAtoms(as []ast.Atom) {
	fmt.Print("{")
	for i, a := range as {
		pp.printAtom(a)
		if i != len(as)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}

func (pp *PP) printVars(vs []ast.Var) {
	fmt.Print("{")
	for i, v := range vs {
		pp.printVar(v)
		if i != len(vs)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}

func (*PP) printVar(v ast.Var)   { fmt.Print(v) }
func (*PP) printCtor(c ast.Ctor) { fmt.Print(c) }
func (*PP) printPrim(p ast.Prim) { fmt.Print(p) }
func (*PP) printLit(l ast.Lit)   { fmt.Print(l) }
func (pp *PP) printAtom(a ast.Atom) {
	switch a := a.(type) {
	case ast.Var:
		pp.printVar(a)

	case ast.Lit:
		pp.printLit(a)
	}
}
