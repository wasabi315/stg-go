package stg

import (
	"fmt"
	"strings"
)

/*****************************************************************************/
/* AST                                                                       */
/*****************************************************************************/

type (
	Var  string
	Ctor string
	Prim string
)

type Lit int

type Atom interface {
	swchAtom(cases AtomCases)
}

type AtomCases struct {
	Var func(Var)
	Lit func(Lit)
}

func (v Var) swchAtom(cases AtomCases) { cases.Var(v) }
func (l Lit) swchAtom(cases AtomCases) { cases.Lit(l) }

type Let struct {
	Rec   bool
	Binds []*Bind
	Body  Expr
}

type Case struct {
	Target Expr
	Alts   []Alt
}

type Alt interface {
	swchAlt(cases AltCases)
}

type AltCases struct {
	AAlt func(*AAlt)
	PAlt func(*PAlt)
	VAlt func(*VAlt)
	DAlt func(*DAlt)
}

type AAlt struct {
	Ctor Ctor
	Vars []Var
	Expr Expr
}

type PAlt struct {
	Lit  Lit
	Expr Expr
}

type VAlt struct {
	Var  Var
	Expr Expr
}

type DAlt struct {
	Expr Expr
}

func (a *AAlt) swchAlt(cases AltCases) { cases.AAlt(a) }
func (p *PAlt) swchAlt(cases AltCases) { cases.PAlt(p) }
func (v *VAlt) swchAlt(cases AltCases) { cases.VAlt(v) }
func (d *DAlt) swchAlt(cases AltCases) { cases.DAlt(d) }

type VarApp struct {
	Var   Var
	Atoms []Atom
}

type CtorApp struct {
	Ctor  Ctor
	Atoms []Atom
}

type PrimApp struct {
	Prim  Prim
	Atoms []Atom
}

type Expr interface {
	swchExpr(cases ExprCases)
}

type ExprCases struct {
	Let     func(*Let)
	Case    func(*Case)
	VarApp  func(*VarApp)
	CtorApp func(*CtorApp)
	PrimApp func(*PrimApp)
	Lit     func(Lit)
}

func (l *Let) swchExpr(cases ExprCases)     { cases.Let(l) }
func (c *Case) swchExpr(cases ExprCases)    { cases.Case(c) }
func (v *VarApp) swchExpr(cases ExprCases)  { cases.VarApp(v) }
func (c *CtorApp) swchExpr(cases ExprCases) { cases.CtorApp(c) }
func (p *PrimApp) swchExpr(cases ExprCases) { cases.PrimApp(p) }
func (l Lit) swchExpr(cases ExprCases)      { cases.Lit(l) }

type LF struct {
	Free []Var
	Upd  bool
	Args []Var
	Body Expr
}

type Bind struct {
	Var Var
	LF  *LF
}

func SwchAtom(a Atom, cases AtomCases) { a.swchAtom(cases) }
func SwchAlt(a Alt, cases AltCases)    { a.swchAlt(cases) }
func SwchExpr(e Expr, cases ExprCases) { e.swchExpr(cases) }

/*****************************************************************************/
/* Print AST                                                                 */
/*****************************************************************************/

var indentWidth = 4

func printIndent(n int) {
	fmt.Print(strings.Repeat(" ", indentWidth*n))
}

func printVar(v Var)   { fmt.Print(v) }
func printCtor(c Ctor) { fmt.Print(c) }
func printPrim(p Prim) { fmt.Print(p) }
func printLit(l Lit)   { fmt.Print(l) }
func printAtom(a Atom) {
	SwchAtom(a, AtomCases{
		Var: printVar,
		Lit: printLit,
	})
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

func printBind(b *Bind, indent int) {
	printIndent(indent)
	printVar(b.Var)
	fmt.Print(" = ")
	printLF(b.LF, indent)
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
	SwchExpr(e, ExprCases{
		Let: func(l *Let) {
			fmt.Println("let")
			printBinds(l.Binds, indent+1)
			printIndent(indent)
			fmt.Println("in")
			printExpr(l.Body, indent+1)
		},
		Case: func(c *Case) {
			fmt.Print("case")
			fmt.Print(" ")
			printExpr(c.Target, 0)
			fmt.Println(" of")
			printAlts(c.Alts, indent+1)
		},
		VarApp: func(v *VarApp) {
			printVar(v.Var)
			fmt.Print(" ")
			printAtoms(v.Atoms)
		},
		CtorApp: func(c *CtorApp) {
			printCtor(c.Ctor)
			fmt.Print(" ")
			printAtoms(c.Atoms)
		},
		PrimApp: func(p *PrimApp) {
			printPrim(p.Prim)
			fmt.Print(" ")
			printAtoms(p.Atoms)
		},
		Lit: printLit,
	})
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
	SwchAlt(a, AltCases{
		AAlt: func(a *AAlt) {
			printCtor(a.Ctor)
			fmt.Print(" ")
			printVars(a.Vars)
			fmt.Println(" ->")
			printExpr(a.Expr, indent+1)
		},
		PAlt: func(p *PAlt) {
			printLit(p.Lit)
			fmt.Println(" ->")
			printExpr(p.Expr, indent+1)
		},
		VAlt: func(v *VAlt) {
			printVar(v.Var)
			fmt.Println(" ->")
			printExpr(v.Expr, indent+1)
		},
		DAlt: func(d *DAlt) {
			fmt.Println("default ->")
			printExpr(d.Expr, indent+1)
		},
	})
}

func PrintProgram(bs []*Bind) {
	printBinds(bs, 0)
}
