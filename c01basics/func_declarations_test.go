package c01basics

import (
	"fmt"
	"testing"
)

// --- Function declarations

func TestFunctionNoReturnValue(t *testing.T) {
	noReturnValue(t.Name())
}

func TestFunctionOneReturnValue(t *testing.T) {
	power := oneReturnValue()
	
	if power != 9000 {
		t.Error()
	}
}

func TestFunctionTwoReturnValues(t *testing.T) {
	power, exists := twoReturnValues()
	
	if (power != 9000) || (!exists) {
		t.Error()
	}
}

func TestBlankIdentifier(t *testing.T) {
	// _ is the blank identifier. The return value isn't actually assigned,
	// and you can use it over and over again regardless of the returned type.
	_, exists := twoReturnValues()
	
	if !exists {
		t.Error()
	}
}

func noReturnValue(message string) {
	fmt.Println("Output from", message)
}

func oneReturnValue() int {
	return 9000
}

func twoReturnValues() (int, bool) {
	return 9000, true
}
