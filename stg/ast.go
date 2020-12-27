package stg

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
