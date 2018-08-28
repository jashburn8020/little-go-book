package c04codeorginterfaces

import (
	"fmt"
	"testing"
)

// ----- Interface

type AddOn interface {
	Add(message string) string
}

// ----- PrefixAddOn

type PrefixAddOn struct{}

// A structure implementing an interface is implicit - as long as it fulfills the contract (has
// function(s) defined in the interface)
func (prefixAdder PrefixAddOn) Add(message string) string {
	return fmt.Sprintf("%s %s", "prefix", message)
}

// ----- PrefixAddOn

type SuffixAddOn struct{}

func (suffixAdder SuffixAddOn) Add(message string) string {
	return fmt.Sprintf("%s %s", message, "suffix")
}

// ----- Tester

type AddOnTester struct {
	// You can use interface in a structure's field
	MessageAdder AddOn
}

// or in a function parameter
func (tester AddOnTester) AddMessageWithInterface(adder AddOn, message string) string {
	return adder.Add(message)
}

func (tester AddOnTester) AddMessage(message string) string {
	return tester.MessageAdder.Add(message)
}

// Note we're using a pointer to the receiver as we're mutating the receiver
func (tester *AddOnTester) SwitchAddOn(adder AddOn) {
	tester.MessageAdder = adder
}

// ----- Tests

func TestPrefixSuffix(t *testing.T) {
	tester := &AddOnTester{
		MessageAdder: &PrefixAddOn{},
	}

	if tester.AddMessage("msg") != "prefix msg" {
		t.Error()
	}

	if tester.AddMessageWithInterface(&SuffixAddOn{}, "msg") != "msg suffix" {
		t.Error()
	}
}

func TestSuffixPrefix(t *testing.T) {
	tester := &AddOnTester{
		MessageAdder: &SuffixAddOn{},
	}

	if tester.AddMessage("msg") != "msg suffix" {
		t.Error()
	}

	if tester.AddMessageWithInterface(&PrefixAddOn{}, "msg") != "prefix msg" {
		t.Error()
	}
}

func TestChangeAddOn(t *testing.T) {
	tester := &AddOnTester{
		MessageAdder: &PrefixAddOn{},
	}

	if tester.AddMessage("msg") != "prefix msg" {
		t.Error()
	}

	tester.SwitchAddOn(&SuffixAddOn{})

	if tester.AddMessage("msg") != "msg suffix" {
		t.Error()
	}
}
