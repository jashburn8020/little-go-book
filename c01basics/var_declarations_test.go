package c01basics

import (
	"testing"
)

// --- Variables and declarations

// The first letter is capitalized, the identifier is exported from the package
func TestMostExplicitVariableDeclaration(t *testing.T) {
	// The most explicit way of variable declaration and assignment
	var power int
	power = 9000
	
	if power != 9000 {
		t.Error()
	}
}

func TestShorterVariableDeclaration(t *testing.T) {
	var power int = 9000

	if power != 9000 {
		t.Error()
	}
}

func TestDeclarationAndAssignmentTypeInferred(t *testing.T) {
	power := 9000

	if power != 9000 {
		t.Error()
	}
}

func TestAssignmentFromFunction(t *testing.T) {
	power := getPower()

	if power != 9000 {
		t.Error()
	}
}

// The first letter is not capitalized, the identifier is unexported (i.e., private)
func getPower() int {
	return 9000
}

func TestMultipleVariablesDeclarationAndAssignment(t *testing.T) {
	power := 8000

	// power is declared a second time here, but allowed because name is new
	name, power := "Goku", 9000
	if (name != "Goku") || (power != 9000) {
		t.Error()
	}
}
