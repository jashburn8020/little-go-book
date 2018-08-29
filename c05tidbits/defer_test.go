package c05tidbits

import (
	"reflect"
	"testing"
)

// rune is an alias for int32, used, by convention, to distinguish character values from integer values.
var record = make([]rune, 0, 5)

func trace(ch rune) rune {
	record = append(record, ch)

	return ch + 'z' - 'a'
}

func un(ch rune) {
	record = append(record, ch)
}

func upper() {
	/* defer statement schedules a function call (the deferred function) to be run immediately before
	 * the function executing the defer returns, i.e., un() is run after lower().
	 * The arguments to the deferred function are evaluated when the defer executes, not when the
	 * call executes, i.e., trace('A') is run first, while un() is run last with the return value
	 * from trace('A') as its argument.
	 * If there are multiple defers in a function, they are executed in LIFO order.
	 */
	defer un(trace('A'))
	record = append(record, 'M')

	lower()
}

func lower() {
	defer un(trace('a'))
	record = append(record, 'm')
}

func TestDefer(t *testing.T) {
	upper()

	if !reflect.DeepEqual([]rune{'A', 'M', 'a', 'm', 'z', 'Z'}, record) {
		t.Error()
	}
}
