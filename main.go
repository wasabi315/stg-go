package main

import (
	"github.com/wasabi315/stg-go/stg/ast"
)

func main() {
	nodes := []*ast.Bind{
		&ast.Bind{
			Var: ast.Var("a"),
			LF: &ast.LF{
				Body: &ast.Let{
					Binds: []*ast.Bind{
						&ast.Bind{
							Var: ast.Var("add"),
							LF: &ast.LF{
								Args: []ast.Var{
									ast.Var("x"),
									ast.Var("y"),
								},
								Body: &ast.PrimApp{
									Prim: ast.Prim("+#"),
									Atoms: []ast.Atom{
										ast.Var("x"),
										ast.Var("y"),
									},
								},
							},
						},
						&ast.Bind{
							Var: ast.Var("sub"),
							LF: &ast.LF{
								Args: []ast.Var{
									ast.Var("x"),
									ast.Var("y"),
								},
								Body: &ast.PrimApp{
									Prim: ast.Prim("-#"),
									Atoms: []ast.Atom{
										ast.Var("x"),
										ast.Var("y"),
									},
								},
							},
						},
					},
					Body: &ast.VarApp{
						Var: ast.Var("add"),
						Atoms: []ast.Atom{
							ast.Lit(10),
							ast.Lit(20),
						},
					},
				},
			},
		},
		&ast.Bind{
			Var: ast.Var("b"),
			LF: &ast.LF{
				Upd: true,
				Body: &ast.Case{
					Target: &ast.PrimApp{
						Prim: "+#",
						Atoms: []ast.Atom{
							ast.Lit(1),
							ast.Lit(2),
						},
					},
					Alts: []ast.Alt{
						&ast.PAlt{
							Lit: 0,
							Expr: &ast.PrimApp{
								Prim: "printInt#",
								Atoms: []ast.Atom{
									ast.Lit(0),
								},
							},
						},
						&ast.PAlt{
							Lit: 1,
							Expr: &ast.PrimApp{
								Prim: "printInt#",
								Atoms: []ast.Atom{
									ast.Lit(1),
								},
							},
						},
						&ast.DAlt{
							Expr: &ast.PrimApp{
								Prim: "printInt#",
								Atoms: []ast.Atom{
									ast.Lit(100),
								},
							},
						},
					},
				},
			},
		},
	}

	ast.PrintProgram(nodes)
}
