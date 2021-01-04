package evaluator

import (
	"fmt"
	"strings"

	"github.com/wasabi315/stg-go/stg/ast"
)

var indentWidth = 4

func printIndent(n int) {
	fmt.Print(strings.Repeat(" ", indentWidth*n))
}

// PrintProgram unparse STG AST
func PrintProgram(bs []*ast.Bind) {
	printBinds(bs, 0)
}

func printBinds(bs []*ast.Bind, indent int) {
	for i, b := range bs {
		printBind(b, indent)
		fmt.Println()
		if i != len(bs)-1 {
			fmt.Println()
		}
	}
}

func printBind(b *ast.Bind, indent int) {
	printIndent(indent)
	printVar(b.Var)
	fmt.Print(" = ")
	printLF(b.LF, indent)
}

func printLF(lf *ast.LF, indent int) {
	printVars(lf.Free)
	if lf.Upd {
		fmt.Print(" \\u ")
	} else {
		fmt.Print(" \\n ")
	}
	printVars(lf.Args)
	fmt.Println(" ->")
	printExpr(lf.Body, indent+1)
}

func printExpr(e ast.Expr, indent int) {
	printIndent(indent)
	switch e := e.(type) {
	case *ast.Let:
		fmt.Println("let")
		printBinds(e.Binds, indent+1)
		printIndent(indent)
		fmt.Println("in")
		printExpr(e.Body, indent+1)

	case *ast.Case:
		fmt.Print("case")
		fmt.Print(" ")
		printExpr(e.Target, 0)
		fmt.Println(" of")
		printAlts(e.Alts, indent+1)

	case *ast.VarApp:
		printVar(e.Var)
		fmt.Print(" ")
		printAtoms(e.Atoms)

	case *ast.CtorApp:
		printCtor(e.Ctor)
		fmt.Print(" ")
		printAtoms(e.Atoms)

	case *ast.PrimApp:
		printPrim(e.Prim)
		fmt.Print(" ")
		printAtoms(e.Atoms)

	case ast.Lit:
		printLit(e)
	}
}

func printAlts(as []ast.Alt, indent int) {
	for i, a := range as {
		printAlt(a, indent)
		fmt.Println()
		if i != len(as)-1 {
			fmt.Println()
		}
	}
}

func printAlt(a ast.Alt, indent int) {
	printIndent(indent)
	switch a := a.(type) {
	case *ast.AAlt:
		printCtor(a.Ctor)
		fmt.Print(" ")
		printVars(a.Vars)
		fmt.Println(" ->")
		printExpr(a.Expr, indent+1)

	case *ast.PAlt:
		printLit(a.Lit)
		fmt.Println(" ->")
		printExpr(a.Expr, indent+1)

	case *ast.VAlt:
		printVar(a.Var)
		fmt.Println(" ->")
		printExpr(a.Expr, indent+1)

	case *ast.DAlt:
		fmt.Println("default ->")
		printExpr(a.Expr, indent+1)
	}
}

func printAtoms(as []ast.Atom) {
	fmt.Print("{")
	for i, a := range as {
		printAtom(a)
		if i != len(as)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}

func printVars(vs []ast.Var) {
	fmt.Print("{")
	for i, v := range vs {
		printVar(v)
		if i != len(vs)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}

func printVar(v ast.Var)   { fmt.Print(v) }
func printCtor(c ast.Ctor) { fmt.Print(c) }
func printPrim(p ast.Prim) { fmt.Print(p) }
func printLit(l ast.Lit)   { fmt.Print(l) }
func printAtom(a ast.Atom) {
	switch a := a.(type) {
	case ast.Var:
		printVar(a)

	case ast.Lit:
		printLit(a)
	}
}
