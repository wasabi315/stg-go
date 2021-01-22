package language

import (
	"fmt"
	"strings"
)

const indentWidth = 4

// PrettyPrint prints stg program
func PrettyPrint(program Program) {
	printBinds(program, 0)
}

func printIndent(n uint) {
	fmt.Print(strings.Repeat(" ", int(indentWidth*n)))
}

func printBinds(bs Binds, indent uint) {
	i := 0
	for v, lf := range bs {
		printIndent(indent)
		printVar(v)
		fmt.Print(" = ")
		printLambdaForm(lf, indent)
		fmt.Println()
		if i != len(bs)-1 {
			fmt.Println()
		}
		i++
	}
}

func printLambdaForm(lf *LambdaForm, indent uint) {
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

func printExpr(e Expr, indent uint) {
	printIndent(indent)
	switch e := e.(type) {
	case *Let:
		fmt.Println("let")
		printBinds(e.Binds, indent+1)
		printIndent(indent)
		fmt.Println("in")
		printExpr(e.Body, indent+1)

	case *Case:
		fmt.Print("case")
		fmt.Print(" ")
		printExpr(e.Expr, 0)
		fmt.Println(" of")
		printAlts(e.Alts, indent+1)

	case *VarApp:
		printVar(e.Var)
		fmt.Print(" ")
		printAtoms(e.Args)

	case *CtorApp:
		printCtor(e.Ctor)
		fmt.Print(" ")
		printAtoms(e.Args)

	case *PrimApp:
		printPrim(e.Prim)
		fmt.Print(" ")
		printAtoms(e.Args)

	case Lit:
		printLit(e)
	}
}

func printAlts(as []Alt, indent uint) {
	for i, a := range as {
		printAlt(a, indent)
		fmt.Println()
		if i != len(as)-1 {
			fmt.Println()
		}
	}
}

func printAlt(a Alt, indent uint) {
	printIndent(indent)
	switch a := a.(type) {
	case *AAlt:
		printCtor(a.Ctor)
		fmt.Print(" ")
		printVars(a.Vars)
		fmt.Println(" ->")
		printExpr(a.Expr, indent+1)

	case *PAlt:
		printLit(a.Lit)
		fmt.Println(" ->")
		printExpr(a.Expr, indent+1)

	case *VAlt:
		printVar(a.Var)
		fmt.Println(" ->")
		printExpr(a.Expr, indent+1)

	case *DAlt:
		fmt.Println("default ->")
		printExpr(a.Expr, indent+1)
	}
}

func printAtoms(as []Atom) {
	fmt.Print("{")
	for i, a := range as {
		printAtom(a)
		if i != len(as)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}

func printVars(vs []Var) {
	fmt.Print("{")
	for i, v := range vs {
		printVar(v)
		if i != len(vs)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}

func printVar(v Var)   { fmt.Print(v) }
func printCtor(c Ctor) { fmt.Print(c) }
func printPrim(p Prim) { fmt.Print(p) }
func printLit(l Lit)   { fmt.Print(l) }
func printAtom(a Atom) {
	switch a := a.(type) {
	case Var:
		printVar(a)

	case Lit:
		printLit(a)
	}
}
