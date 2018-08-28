package c04codeorginterfaces

import (
	"testing"
	"fmt"
)

// ----- Interface

type AddOn interface {
	Add(message string) string
}

// ----- PrefixAddOn

type PrefixAddOn struct {}

func (prefixAdder PrefixAddOn) Add(message string) string {
	return fmt.Sprintf("%s %s", "prefix", message)
}

// ----- PrefixAddOn

type SuffixAddOn struct {}

func (suffixAdder SuffixAddOn) Add(message string) string {
	return fmt.Sprintf("%s %s", message, "suffix")
}

// ----- Tester

type AddOnTester struct {
	MessageAdder AddOn
}

func (tester AddOnTester) AddMessageWithInterface(adder AddOn, message string) string {
	return adder.Add(message)
}

func (tester AddOnTester) AddMessage(message string) string {
	return tester.MessageAdder.Add(message)
}

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

func TestSwitchAddOn(t *testing.T) {
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
