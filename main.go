package main

import (
	lang "github.com/wasabi315/stg-go/language"
)

func main() {
	program := lang.Program{
		lang.Var("main"): &lang.LambdaForm{
			Body: &lang.Let{
				Binds: lang.Binds{
					lang.Var("add"): &lang.LambdaForm{
						Args: []lang.Var{
							lang.Var("x"),
							lang.Var("y"),
						},
						Body: &lang.PrimApp{
							Prim: lang.Prim("+#"),
							Args: []lang.Atom{
								lang.Var("x"),
								lang.Var("y"),
							},
						},
					},
					lang.Var("sub"): &lang.LambdaForm{
						Args: []lang.Var{
							lang.Var("x"),
							lang.Var("y"),
						},
						Body: &lang.PrimApp{
							Prim: lang.Prim("-#"),
							Args: []lang.Atom{
								lang.Var("x"),
								lang.Var("y"),
							},
						},
					},
				},
				Body: &lang.VarApp{
					Var: lang.Var("add"),
					Args: []lang.Atom{
						lang.Lit(10),
						lang.Lit(20),
					},
				},
			},
		},
		lang.Var("b"): &lang.LambdaForm{
			Upd: true,
			Body: &lang.Case{
				Expr: &lang.PrimApp{
					Prim: "+#",
					Args: []lang.Atom{
						lang.Lit(1),
						lang.Lit(2),
					},
				},
				Alts: []lang.Alt{
					&lang.PAlt{
						Lit: 0,
						Expr: &lang.PrimApp{
							Prim: "printInt#",
							Args: []lang.Atom{
								lang.Lit(0),
							},
						},
					},
					&lang.PAlt{
						Lit: 1,
						Expr: &lang.PrimApp{
							Prim: "printInt#",
							Args: []lang.Atom{
								lang.Lit(1),
							},
						},
					},
					&lang.DAlt{
						Expr: &lang.PrimApp{
							Prim: "printInt#",
							Args: []lang.Atom{
								lang.Lit(100),
							},
						},
					},
				},
			},
		},
	}

	lang.PrettyPrint(program)
}
