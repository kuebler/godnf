package formula

import (
	"reflect"
	"testing"
)

func TestFlattenOr(t *testing.T) {
	or1 := MkOr(Literal(1), Literal(2))
	or2 := MkOr(Literal(3), Literal(4))

	or1.Flatten(or2.Children...)

	if len(or1.Children) != 4 {
		t.Errorf("Flattening failed: %d, expected 4", len(or1.Children))
	}
}

func TestFlattenAnd(t *testing.T) {
	and1 := MkAnd(Literal(1), Literal(2))
	and2 := MkAnd(Literal(3), Literal(4))

	and1.Flatten(and2.Children...)

	if len(and1.Children) != 4 {
		t.Errorf("Flattening failed: %d, expected 4", len(and1.Children))
	}
}

func TestLiteralDnf(t *testing.T) {
	lit := Literal(1)

	if lit.MkDNF() != lit {
		t.Errorf("DNF representation supposed to be %s, got %s", lit, lit)
	}
}

func TestAllLiteralsOrDnf(t *testing.T) {
	or1 := MkOr(Literal(1), Literal(2), Literal(3))
	dnf := MkOr(Literal(1), Literal(2), Literal(3))

	if !reflect.DeepEqual(or1.MkDNF(), dnf) {
		t.Errorf("DNF representation supposed to be %s, got %s", dnf, or1)
	}
}

func TestAllLiteralsAndDnf(t *testing.T) {
	and := MkAnd(Literal(1), Literal(2), Literal(3))
	dnf := MkAnd(Literal(1), Literal(2), Literal(3))

	if !reflect.DeepEqual(and.MkDNF(), dnf) {
		t.Errorf("DNF representation supposed to be %s, got %s", dnf, and)
	}
}

func TestOneTermAndLiteralsDnf(t *testing.T) {
	or1 := MkOr(Literal(1), MkAnd(Literal(1), Literal(2)), Literal(3))
	dnf := MkOr(Literal(1), MkAnd(Literal(1), Literal(2)), Literal(3))

	if !reflect.DeepEqual(or1.MkDNF(), dnf) {
		t.Errorf("DNF representation supposed to be %s, got %s", dnf, or1)
	}
}

func TestTwoClausesDnf(t *testing.T) {
	fm := MkAnd(
		MkOr(Literal(1), Literal(2)),
		MkOr(Literal(3), Literal(4)))
	expected := MkOr(MkAnd(Literal(3), Literal(1)),
		MkAnd(Literal(4), Literal(1)),
		MkAnd(Literal(3), Literal(2)),
		MkAnd(Literal(4), Literal(2)))
	dnf := fm.MkDNF()

	if !reflect.DeepEqual(dnf, expected) {
		t.Errorf("DNF representation supposed to be %s, got %s", expected, dnf)
	}
}

func TestThreeClausesDnf(t *testing.T) {
	fm := MkAnd(
		MkOr(Literal(1), Literal(2)),
		MkOr(Literal(3), Literal(4)),
		MkOr(Literal(5), Literal(6)))
	expected := MkOr(MkAnd(Literal(5), Literal(3), Literal(1)),
		MkAnd(Literal(6), Literal(3), Literal(1)),
		MkAnd(Literal(5), Literal(4), Literal(1)),
		MkAnd(Literal(6), Literal(4), Literal(1)),
		MkAnd(Literal(5), Literal(3), Literal(2)),
		MkAnd(Literal(6), Literal(3), Literal(2)),
		MkAnd(Literal(5), Literal(4), Literal(2)),
		MkAnd(Literal(6), Literal(4), Literal(2)))
	dnf := fm.MkDNF()

	if !reflect.DeepEqual(dnf, expected) {
		t.Errorf("DNF representation supposed to be %s, got %s", expected, dnf)
	}
}
