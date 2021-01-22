package language

/*****************************************************************************/
/* AST Nodes                                                                 */
/*****************************************************************************/

type Program = Binds

type Binds map[Var]*LambdaForm

// LambdaForm lambda form
type LambdaForm struct {
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
		Binds Binds
		Body  Expr
	}

	// Case case expression
	Case struct {
		Expr Expr
		Alts []Alt
	}

	// VarApp variable application expression
	VarApp struct {
		Var  Var
		Args []Atom
	}

	// CtorApp constructor application expression
	CtorApp struct {
		Ctor Ctor
		Args []Atom
	}

	// PrimApp primitive application expression
	PrimApp struct {
		Prim Prim
		Args []Atom
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
