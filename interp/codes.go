package interp

import (
	"github.com/wasabi315/stg-go/lang"
)

func eval(expr lang.Expr, locals env) code {
	return func(s *state) {}
}

func enter(addr addr) code {
	return func(s *state) {}
}

func returnCon(ctor lang.Ctor, args []value) code {
	return func(s *state) {}
}

func returnInt(n imm) code {
	return func(s *state) {}
}
