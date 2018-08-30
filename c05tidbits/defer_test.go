package c05tidbits

import (
	"reflect"
	"testing"
)

type trace struct {
	// rune is an alias for int32, used, by convention, to distinguish character values from integer
	// values.
	record []rune
}

func (tr *trace) trace(ch rune) rune {
	tr.record = append(tr.record, ch)

	return ch + 'z' - 'a'
}

func (tr *trace) un(ch rune) {
	tr.record = append(tr.record, ch)
}

func (tr *trace) upper() {
	/* defer statement schedules a function call (the deferred function) to be run immediately before
	 * the function executing the defer returns, i.e., un() is run after lower().
	 * The arguments to the deferred function are evaluated when the defer executes, not when the
	 * call executes, i.e., trace('A') is run first, while un() is run last with the return value
	 * from trace('A') as its argument.
	 * If there are multiple defers in a function, they are executed in LIFO order.
	 */
	defer tr.un(tr.trace('A'))
	tr.record = append(tr.record, 'M')

	tr.lower()
}

func (tr *trace) lower() {
	defer tr.un(tr.trace('a'))
	tr.record = append(tr.record, 'm')
}

func TestDefer(t *testing.T) {
	tr := &trace{make([]rune, 0, 5)}
	tr.upper()

	if !reflect.DeepEqual([]rune{'A', 'M', 'a', 'm', 'z', 'Z'}, tr.record) {
		t.Error()
	}
}
