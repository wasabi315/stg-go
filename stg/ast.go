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
	VisitAtom(vis AtomVistor)
}

type AtomVistor struct {
	Var func(Var)
	Lit func(Lit)
}

func (v Var) VisitAtom(vis AtomVistor) {
	vis.Var(v)
}

func (l Lit) VisitAtom(vis AtomVistor) {
	vis.Lit(l)
}

type LetExpr struct {
	Rec   bool
	Binds []*Bind
	Body  Expr
}

type CaseExpr struct {
	Target Expr
	Alts   []Alt
}

type Alt interface {
	VisitAlt(vis AltVisitor)
}

type AltVisitor struct {
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

func (a *AAlt) VisitAlt(vis AltVisitor) {
	vis.AAlt(a)
}

func (p *PAlt) VisitAlt(vis AltVisitor) {
	vis.PAlt(p)
}

func (v *VAlt) VisitAlt(vis AltVisitor) {
	vis.VAlt(v)
}

func (d *DAlt) VisitAlt(vis AltVisitor) {
	vis.DAlt(d)
}

type VarAppExpr struct {
	Var   Var
	Atoms []Atom
}

type CtorAppExpr struct {
	Ctor  Ctor
	Atoms []Atom
}

type PrimAppExpr struct {
	Prim  Prim
	Atoms []Atom
}

type Expr interface {
	VisitExpr(vis ExprVisitor)
}

type ExprVisitor struct {
	Let     func(*LetExpr)
	Case    func(*CaseExpr)
	VarApp  func(*VarAppExpr)
	CtorApp func(*CtorAppExpr)
	PrimApp func(*PrimAppExpr)
	Lit     func(Lit)
}

func (l *LetExpr) VisitExpr(vis ExprVisitor) {
	vis.Let(l)
}

func (c *CaseExpr) VisitExpr(vis ExprVisitor) {
	vis.Case(c)
}

func (v *VarAppExpr) VisitExpr(vis ExprVisitor) {
	vis.VarApp(v)
}

func (c *CtorAppExpr) VisitExpr(vis ExprVisitor) {
	vis.CtorApp(c)
}

func (p *PrimAppExpr) VisitExpr(vis ExprVisitor) {
	vis.PrimApp(p)
}

func (l Lit) VisitExpr(vis ExprVisitor) {
	vis.Lit(l)
}

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

/*****************************************************************************/
/* Print AST                                                                 */
/*****************************************************************************/

const indentWidth = 4

func PrintProgram(bs []*Bind) {
	PrintBinds(bs, 0)
}

func PrintBinds(bs []*Bind, indent int) {
	for i, b := range bs {
		PrintBind(b, indent)
		fmt.Println()
		if i != len(bs)-1 {
			fmt.Println()
		}
	}
}

func PrintBind(b *Bind, indent int) {
	PrintIndent(indent)
	PrintVar(b.Var)
	fmt.Print(" = ")
	PrintLF(b.LF, indent)
}

func PrintVar(v Var) {
	fmt.Print(v)
}

func PrintVars(vs []Var) {
	fmt.Print("{")
	for i, v := range vs {
		PrintVar(v)
		if i != len(vs)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}

func PrintLF(lf *LF, indent int) {
	PrintVars(lf.Free)

	if lf.Upd {
		fmt.Print(" \\u ")
	} else {
		fmt.Print(" \\n ")
	}

	PrintVars(lf.Args)

	fmt.Println(" ->")

	PrintExpr(lf.Body, indent+1)
}

func PrintExpr(e Expr, indent int) {
	PrintIndent(indent)
	e.VisitExpr(ExprVisitor{
		Let: func(l *LetExpr) {
			fmt.Println("let")
			PrintBinds(l.Binds, indent+1)
			PrintIndent(indent)
			fmt.Println("in")
			PrintExpr(l.Body, indent+1)
		},
		Case: func(c *CaseExpr) {
			fmt.Print("case")
			fmt.Print(" ")
			PrintExpr(c.Target, 0)
			fmt.Println(" of")
			PrintAlts(c.Alts, indent+1)
		},
		VarApp: func(v *VarAppExpr) {
			PrintVar(v.Var)
			fmt.Print(" ")
			PrintAtoms(v.Atoms)
		},
		CtorApp: func(c *CtorAppExpr) {
			PrintCtor(c.Ctor)
			fmt.Print(" ")
			PrintAtoms(c.Atoms)
		},
		PrimApp: func(p *PrimAppExpr) {
			PrintPrim(p.Prim)
			fmt.Print(" ")
			PrintAtoms(p.Atoms)
		},
		Lit: PrintLit,
	})
}

func PrintCtor(c Ctor) {
	fmt.Print(c)
}

func PrintPrim(p Prim) {
	fmt.Print(p)
}

func PrintLit(l Lit) {
	fmt.Print(l)
}

func PrintAtom(a Atom) {
	a.VisitAtom(AtomVistor{
		Var: PrintVar,
		Lit: PrintLit,
	})
}

func PrintAtoms(as []Atom) {
	fmt.Print("{")
	for i, a := range as {
		PrintAtom(a)
		if i != len(as)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}

func PrintAlts(as []Alt, indent int) {
	for i, a := range as {
		PrintAlt(a, indent)
		fmt.Println()
		if i != len(as)-1 {
			fmt.Println()
		}
	}
}

func PrintAlt(a Alt, indent int) {
	PrintIndent(indent)
	a.VisitAlt(AltVisitor{
		AAlt: func(a *AAlt) {
			PrintCtor(a.Ctor)
			fmt.Print(" ")
			PrintVars(a.Vars)
			fmt.Println(" ->")
			PrintExpr(a.Expr, indent+1)
		},
		PAlt: func(p *PAlt) {
			PrintLit(p.Lit)
			fmt.Println(" ->")
			PrintExpr(p.Expr, indent+1)
		},
		VAlt: func(v *VAlt) {
			PrintVar(v.Var)
			fmt.Println(" ->")
			PrintExpr(v.Expr, indent+1)
		},
		DAlt: func(d *DAlt) {
			fmt.Println("default ->")
			PrintExpr(d.Expr, indent+1)
		},
	})
}

func PrintIndent(n int) {
	fmt.Print(strings.Repeat(" ", n*indentWidth))
}
