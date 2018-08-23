package c03datastructures

import (
	"reflect"
	"testing"
)

func TestMapGet(t *testing.T) {
	// Create with the make function
	lookup := make(map[string]int)

	// Maps grow dynamically, but we can supply a second argument to make to set an initial size,
	// which can help with performance
	// lookup := make(map[string]int, 100)

	lookup["goku"] = 9000
	lookup["gohan"] = 1000

	power, exists := lookup["goku"]
	if !(power == 9000 && exists) {
		t.Error()
	}

	power, exists = lookup["vegeta"]
	if !(power == 0 && !exists) {
		t.Error()
	}
}

func TestMapLenDelete(t *testing.T) {
	// Create as a composite literal
	lookup := map[string]int{
		"goku":  9000,
		"gohan": 1000,
	}

	if !(len(lookup) == 2) {
		t.Error()
	}

	// No return value, can be called on a non-existing key
	delete(lookup, "gohan")
	if !(len(lookup) == 1) {
		t.Error()
	}
}

func TestMapIterate(t *testing.T) {
	lookup := map[string]int{
		"goku":  9000,
		"gohan": 1000,
	}

	clone := make(map[string]int, len(lookup))

	// Iteration over maps isn't ordered
	for key, value := range lookup {
		clone[key] = value
	}

	if !reflect.DeepEqual(lookup, clone) {
		t.Error()
	}
}

type SaiyanFriend struct {
	Name    string
	Friends map[string]*SaiyanFriend
}

func TestMapInStruct(t *testing.T) {
	vegeta := &SaiyanFriend{Name: "Vegeta"}

	bulma := &SaiyanFriend{Name: "Bulma"}

	krillin := &SaiyanFriend{
		Name:    "Krillin",
		Friends: map[string]*SaiyanFriend{"bulma": bulma},
	}

	goku := &SaiyanFriend{
		Name: "Goku",
		Friends: map[string]*SaiyanFriend{
			"krillin": krillin,
			"vegeta":  vegeta,
		},
	}

	if goku.Friends["krillin"].Name != "Krillin" {
		t.Error()
	}

	if goku.Friends["krillin"].Friends["bulma"] == nil {
		t.Error()
	}

	if goku.Friends["aka"] != nil {
		t.Error()
	}
}
