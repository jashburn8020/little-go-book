package c05tidbits

import (
	"fmt"
	"testing"
)

type myStruct struct{}

/* All types/interfaces implicitly implement the empty interface (no methods), and so anything can
 * be passed in.
 */
func getType(i interface{}) string {
	/* Type switch can be used to discover the dynamic type of an interface variable. It uses the
	 * syntax of a type assertion with the keyword type inside the parentheses
	 */
	switch i.(type) {
	case int:
		return fmt.Sprintf("an %T", i)
	case string, myStruct:
		return fmt.Sprintf("a %T", i)
	default:
		return ""
	}
}

func TestGetTypeInt(t *testing.T) {
	if getType(1) != "an int" {
		t.Error()
	}
}

func TestGetTypeString(t *testing.T) {
	if getType("1") != "a string" {
		t.Error()
	}
}

func TestGetTypeCustom(t *testing.T) {
	if getType(myStruct{}) != "a c05tidbits.myStruct" {
		t.Error()
	}
}
