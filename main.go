package main

import (
	"github.com/wasabi315/stg-go/stg"
)

func main() {
	ast := []*stg.Bind{
		&stg.Bind{
			Var: stg.Var("a"),
			LF: &stg.LF{
				Body: &stg.LetExpr{
					Binds: []*stg.Bind{
						&stg.Bind{
							Var: stg.Var("add"),
							LF: &stg.LF{
								Args: []stg.Var{
									stg.Var("x"),
									stg.Var("y"),
								},
								Body: &stg.PrimAppExpr{
									Prim: stg.Prim("+#"),
									Atoms: []stg.Atom{
										stg.Var("x"),
										stg.Var("y"),
									},
								},
							},
						},
						&stg.Bind{
							Var: stg.Var("sub"),
							LF: &stg.LF{
								Args: []stg.Var{
									stg.Var("x"),
									stg.Var("y"),
								},
								Body: &stg.PrimAppExpr{
									Prim: stg.Prim("-#"),
									Atoms: []stg.Atom{
										stg.Var("x"),
										stg.Var("y"),
									},
								},
							},
						},
					},
					Body: &stg.VarAppExpr{
						Var: stg.Var("add"),
						Atoms: []stg.Atom{
							stg.Lit(10),
							stg.Lit(20),
						},
					},
				},
			},
		},
		&stg.Bind{
			Var: stg.Var("b"),
			LF: &stg.LF{
				Upd: true,
				Body: &stg.CaseExpr{
					Target: &stg.PrimAppExpr{
						Prim: "+#",
						Atoms: []stg.Atom{
							stg.Lit(1),
							stg.Lit(2),
						},
					},
					Alts: []stg.Alt{
						&stg.PAlt{
							Lit: 0,
							Expr: &stg.PrimAppExpr{
								Prim: "printInt#",
								Atoms: []stg.Atom{
									stg.Lit(0),
								},
							},
						},
						&stg.PAlt{
							Lit: 1,
							Expr: &stg.PrimAppExpr{
								Prim: "printInt#",
								Atoms: []stg.Atom{
									stg.Lit(1),
								},
							},
						},
						&stg.DAlt{
							Expr: &stg.PrimAppExpr{
								Prim: "printInt#",
								Atoms: []stg.Atom{
									stg.Lit(100),
								},
							},
						},
					},
				},
			},
		},
	}

	stg.PrintProgram(ast)
}
