package formula

type Formula interface {
	MkDNF() Formula
}

type OpWithChildren struct {
	Children []Formula
}

type Literal int32
type And OpWithChildren
type Or OpWithChildren

func MkOr(children ...Formula) (or Or) {
	or = Or{[]Formula{}}
	or.Flatten(children...)
	return
}

func MkAnd(children ...Formula) (and And) {
	and = And{[]Formula{}}
	and.Flatten(children...)
	return
}

func (lit Literal) MkDNF() Formula {
	return lit
}

func (or Or) MkDNF() Formula {
	for index, fm := range or.Children {
		or.Children[index] = fm.MkDNF()
	}
	or.FlattenChildren()
	return or
}

func (and And) MkDNF() Formula {
	for index, fm := range and.Children {
		and.Children[index] = fm.MkDNF()
	}
	and.FlattenChildren()
	return Distribute(and)
}

func Distribute(formula Formula) Formula {
	if lit, ok := formula.(Literal); ok {
		return lit
	}
	if and, ok := formula.(And); ok {
		return DistributeAnd(and)
	}
	return formula
}

func DistributeAnd(and And) Formula {
	if len(and.Children) == 1 {
		return and.Children[0]
	} else {
		first := and.Children[0]
		for _, next := range and.Children[1:] {
			first = DistributeTwo(first, next)
		}
		return first
	}
}

func DistributeTwo(left Formula, right Formula) Formula {
	if or1, ok := left.(Or); ok {
		var dnfs []Formula = []Formula{}
		for _, formula := range or1.Children {
			dnfs = append(dnfs, DistributeTwo(formula, right))
		}
		return MkOr(dnfs...)
	} else if _, ok := right.(Or); ok {
		return DistributeTwo(right, left)
	} else {
		return MkAnd(left, right)
	}
}

func (or *Or) Flatten(formulae ...Formula) {
	or.Children = append(or.Children, formulae...)
	or.FlattenChildren()
}

func (and *And) Flatten(formulae ...Formula) {
	and.Children = append(and.Children, formulae...)
	and.FlattenChildren()
}

func (or *Or) FlattenChildren() {
	var buffer []Formula = make([]Formula, 0, len(or.Children))
	for _, formula := range or.Children {
		t, ok := formula.(Or)
		if ok {
			buffer = append(buffer, t.Children...)
		} else {
			buffer = append(buffer, formula)
		}
	}
	or.Children = buffer
}

func (and *And) FlattenChildren() {
	var buffer []Formula = make([]Formula, 0, len(and.Children))
	for _, formula := range and.Children {
		t, ok := formula.(And)
		if ok {
			buffer = append(buffer, t.Children...)
		} else {
			buffer = append(buffer, formula)
		}
	}
	and.Children = buffer
}
