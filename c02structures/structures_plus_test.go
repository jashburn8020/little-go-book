package c02structures

import (
	"testing"
)

func TestFunctionsOnStructures(t *testing.T) {
	goku := &SaiyanBasic{"Goku", 9000}
	goku.power()

	if goku.Power != 19000 {
		t.Error()
	}
}

// Type *SaiyanBasic is the receiver of the power method - adds the method to the structure
func (s *SaiyanBasic) power() {
	s.Power += 10000
}

func TestConstructor(t *testing.T) {
	// Structures don't have constructors. Instead, you can create a function like a factory
	goku := newSaiyan("Goku", 9000)

	if (goku.Name != "Goku") || (goku.Power != 9000) {
		t.Error()
	}
}

func newSaiyan(name string, power int) *SaiyanBasic {
	saiyanAddr := &SaiyanBasic{
		Name:  name,
		Power: power,
	}
	return saiyanAddr
}

func TestNew(t *testing.T) {
	// The new built-in function allocates memory required by a type
	gokuNew := new(SaiyanBasic)
	gokuNew.Name = "Goku"
	gokuNew.Power = 9000

	// is the same as
	goku := &SaiyanBasic{Name: "Goku", Power: 9000}

	if (gokuNew.Name != goku.Name) || (gokuNew.Power != goku.Power) {
		t.Error()
	}
}

func TestFieldsOfAStructure(t *testing.T) {
	// Fields can be of any type, including other structures, arrays, maps, intefaces and functions
	gohan := &SaiyanFather{
		Name:  "Gohan",
		Power: 1000,
		Father: &SaiyanFather{
			Name:   "Goku",
			Power:  9000,
			Father: nil,
		},
	}

	if (gohan.Father.Name != "Goku") || (gohan.Father.Power != 9000) {
		t.Error()
	}
}

// Function added to a composed type
func (p *Person) GetName() string {
	return p.Name
}

func TestComposition(t *testing.T) {
	goku := &SaiyanPerson{
		Person: &Person{"Goku"},
		Power:  9000,
	}
	// With composition there's no need to duplicate methods from composed types
	if goku.GetName() != "Goku" {
		t.Error()
	}

	// 2 ways to refer to a field from a composed type
	if goku.Name != goku.Person.Name {
		t.Error()
	}
}

// Go doesn't support overloading, but you can override the functions of a composed type.
// This overrides GetName() in Person (above)
func (s *SaiyanOverride) GetName() string {
	return s.Name + " overridden"
}

func TestOverride(t *testing.T) {
	goku := &SaiyanOverride{
		Person: &Person{"Goku"},
		Power:  9000,
	}

	if goku.GetName() == goku.Person.GetName() {
		t.Error()
	}
}
