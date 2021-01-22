package interp

import (
	"fmt"

	"github.com/wasabi315/stg-go/lang"
)

func Run(program lang.Program) {
	NewInterpreter().Run(program)
}

type Interpreter struct {
	state *state
}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Run(program lang.Program) {
	i.state = newState(program)
	fmt.Println("Running ...")
}

type state struct {
	code        code
	argStack    []value
	returnStack []returnStackFrame
	updateStack []updateStackFrame
	globalEnv   env
	fresh       int
}

func newState(program lang.Program) *state {
	globals := env{}

	for v, lf := range program {
		globals[v] = &closure{lf, []value{}}
	}

	return &state{
		code: eval(&lang.VarApp{
			Var:  lang.Var("main"),
			Args: []lang.Atom{},
		}, env{}),
		argStack:    []value{},
		returnStack: []returnStackFrame{},
		updateStack: []updateStackFrame{},
		globalEnv:   globals,
		fresh:       0,
	}
}

type code func(*state)

type returnStackFrame struct {
	alts []lang.Alt
	env  env
}

type updateStackFrame struct {
	argStack    []value
	returnStack []returnStackFrame
	addr        addr
}

type env map[lang.Var]value

type closure struct {
	lambdaFrom *lang.LambdaForm
	freeValues []value
}

type (
	value interface {
		isValue()
	}

	imm  int
	addr = *closure // nil for blackhole
)

func (imm) isValue()  {}
func (addr) isValue() {}
