package ast

import (
	"fmt"
	"strings"
)

/*****************************************************************************/
/* AST Nodes                                                                 */
/*****************************************************************************/

// Bind bind
type Bind struct {
	Var Var
	LF  *LF
}

// LF lambda form
type LF struct {
	Free []Var
	Upd  bool
	Args []Var
	Body Expr
}

// Expr expression
type Expr interface {
	isExpr()
}

// Alt alternative in case expression
type Alt interface {
	isAlt()
}

// Atom atom
type Atom interface {
	isAtom()
}

type (
	// Let let expression
	Let struct {
		Rec   bool
		Binds []*Bind
		Body  Expr
	}

	// Case case expression
	Case struct {
		Target Expr
		Alts   []Alt
	}

	// VarApp variable application expression
	VarApp struct {
		Var   Var
		Atoms []Atom
	}

	// CtorApp constructor application expression
	CtorApp struct {
		Ctor  Ctor
		Atoms []Atom
	}

	// PrimApp primitive application expression
	PrimApp struct {
		Prim  Prim
		Atoms []Atom
	}
)

func (*Let) isExpr()     {}
func (*Case) isExpr()    {}
func (*VarApp) isExpr()  {}
func (*CtorApp) isExpr() {}
func (*PrimApp) isExpr() {}
func (Lit) isExpr()      {}

type (
	// AAlt algebraic alternative
	AAlt struct {
		Ctor Ctor
		Vars []Var
		Expr Expr
	}

	// PAlt primitive alternative
	PAlt struct {
		Lit  Lit
		Expr Expr
	}

	// VAlt default alternative with binding
	VAlt struct {
		Var  Var
		Expr Expr
	}

	// DAlt default alternative
	DAlt struct {
		Expr Expr
	}
)

func (*AAlt) isAlt() {}
func (*PAlt) isAlt() {}
func (*VAlt) isAlt() {}
func (*DAlt) isAlt() {}

type (
	// Var variable name
	Var string

	// Ctor constructor name
	Ctor string

	// Prim primitive operation name
	Prim string

	// Lit literal
	Lit int
)

func (Var) isAtom() {}
func (Lit) isAtom() {}

/*****************************************************************************/
/* Print AST                                                                 */
/*****************************************************************************/

var indentWidth = 4

func printIndent(n int) {
	fmt.Print(strings.Repeat(" ", indentWidth*n))
}

// PrintProgram unparse STG AST
func PrintProgram(bs []*Bind) {
	printBinds(bs, 0)
}

func printBinds(bs []*Bind, indent int) {
	for i, b := range bs {
		printBind(b, indent)
		fmt.Println()
		if i != len(bs)-1 {
			fmt.Println()
		}
	}
}

func printBind(b *Bind, indent int) {
	printIndent(indent)
	printVar(b.Var)
	fmt.Print(" = ")
	printLF(b.LF, indent)
}

func printLF(lf *LF, indent int) {
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

func printExpr(e Expr, indent int) {
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
		printExpr(e.Target, 0)
		fmt.Println(" of")
		printAlts(e.Alts, indent+1)

	case *VarApp:
		printVar(e.Var)
		fmt.Print(" ")
		printAtoms(e.Atoms)

	case *CtorApp:
		printCtor(e.Ctor)
		fmt.Print(" ")
		printAtoms(e.Atoms)

	case *PrimApp:
		printPrim(e.Prim)
		fmt.Print(" ")
		printAtoms(e.Atoms)

	case Lit:
		printLit(e)
	}
}

func printAlts(as []Alt, indent int) {
	for i, a := range as {
		printAlt(a, indent)
		fmt.Println()
		if i != len(as)-1 {
			fmt.Println()
		}
	}
}

func printAlt(a Alt, indent int) {
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
