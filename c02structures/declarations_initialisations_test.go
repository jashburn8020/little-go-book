package c02structures

import (
	"testing"
)

func TestDeclarationsAndInitialisations(t *testing.T) {

	goku := SaiyanBasic{
		Name:  "Goku",
		Power: 9000,
	}

	if (goku.Name != "Goku") || (goku.Power != 9000) {
		t.Error()
	}
}

func TestDeclarationsSetPartFields(t *testing.T) {
	// We don't have to set all or even any of the fields
	goku := SaiyanBasic{Name: "Goku"}

	if (goku.Name != "Goku") || (goku.Power != 0) {
		t.Error()
	}

	goku.Power = 9000

	if goku.Power != 9000 {
		t.Error()
	}
}

func TestDeclarationOrderOfFields(t *testing.T) {
	// You can rely on the order of the field declarations, but do this only for structures with few
	// fields
	goku := SaiyanBasic{"Goku", 9000}

	if (goku.Name != "Goku") || (goku.Power != 9000) {
		t.Error()
	}
}

func TestArgumentsToFunctionsAsCopies(t *testing.T) {
	goku := SaiyanBasic{"Goku", 9000}
	// Go passes arguments to a function as copies
	copyArg(t, goku)

	if (goku.Name != "Goku") || (goku.Power != 9000) {
		t.Error()
	}
}

func copyArg(t *testing.T, s SaiyanBasic) {
	s.Name = "Goku Copy"
	s.Power += 10000

	if (s.Name != "Goku Copy") || (s.Power != 19000) {
		t.Error()
	}
}

func TestArgumentsAddressAndPointers(t *testing.T) {
	// & is the address operator.
	goku := &SaiyanBasic{"Goku", 9000}
	pointerArg(goku)

	if (goku.Name != "Goku Pointer") || (goku.Power != 19000) {
		t.Error()
	}
}

// *SaiyanBasic is a pointer to value of type SaiyanBasic
func pointerArg(s *SaiyanBasic) {
	s.Name = "Goku Pointer"
	s.Power += 10000
}
