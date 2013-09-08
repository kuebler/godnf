package main

import (
	"fmt"
	"formula"
)

func main() {
	cnf := formula.MkAnd(
		formula.MkOr(formula.Literal(0), formula.Literal(1)),
		formula.MkOr(formula.Literal(2), formula.Literal(3)),
		formula.MkOr(formula.Literal(4), formula.Literal(5)),
		formula.MkOr(formula.Literal(6), formula.Literal(7)),
		formula.MkOr(formula.Literal(8), formula.Literal(9)),
		formula.MkOr(formula.Literal(10), formula.Literal(11)),
		formula.MkOr(formula.Literal(12), formula.Literal(13)),
		formula.MkOr(formula.Literal(14), formula.Literal(15)),
		formula.MkOr(formula.Literal(16), formula.Literal(17)),
		formula.MkOr(formula.Literal(18), formula.Literal(19)),
		formula.MkOr(formula.Literal(20), formula.Literal(21)),
		formula.MkOr(formula.Literal(22), formula.Literal(23)),
		formula.MkOr(formula.Literal(24), formula.Literal(25)),
		formula.MkOr(formula.Literal(26), formula.Literal(27)),
		formula.MkOr(formula.Literal(28), formula.Literal(29)),
		formula.MkOr(formula.Literal(30), formula.Literal(31)),
		formula.MkOr(formula.Literal(32), formula.Literal(33)),
		formula.MkOr(formula.Literal(34), formula.Literal(35)))
	dnf := cnf.MkDNF()
	fmt.Println(dnf)
	or, ok := dnf.(formula.Or)
	if ok {
		fmt.Println("Num of terms: ", len(or.Children))
	}
}
