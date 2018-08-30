package c05tidbits

import (
	"errors"
	"fmt"
	"testing"
)

/* Common pattern in Go's standard library of using error variables as a
- package variable: defined outside of a function
- publicly accessible: upper-case first letter
*/
var LBE = errors.New("Lower bound error")
var UBE = errors.New("Upper bound error")

func process(lower int, upper int) (int, error) {
	if lower < 1 {
		return 0, LBE
	}

	if upper > 9 {
		return 0, UBE
	}

	processed := lower + upper

	if processed == 13 {
		//  Create our own errors by importing the errors package and using it in the New function
		return 0, errors.New("Unlucky")
	}

	return processed, nil
}

func TestLowerError(t *testing.T) {
	// Initialised if: err is initialised prior to the condition being evaluated
	if _, err := process(0, 5); !(err != nil && err == LBE) {
		t.Error()
	}
}

func TestUpperError(t *testing.T) {
	_, err := process(5, 10)

	if !(err != nil && err == UBE) {
		t.Error()
	}
}

func TestUnlucky(t *testing.T) {
	_, err := process(6, 7)

	// The error built-in interface type has only 1 method
	if !(err != nil && err.Error() == "Unlucky") {
		t.Error()
	}
}

func TestOK(t *testing.T) {
	processed, err := process(2, 3)

	if !(err == nil && processed == 5) {
		t.Error()
	}
}

// ----- Custom error

type CustomError struct {
	Lower int
	Upper int
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Lower is %d, upper is %d", e.Lower, e.Upper)
}

func customProcess(lower int, upper int) (int, error) {
	if lower < 1 || upper > 9 {
		return 0, &CustomError{lower, upper}
	}

	return lower + upper, nil
}

func TestCustomError(t *testing.T) {
	_, err := customProcess(0, 3)

	if !(err != nil && err.Error() == "Lower is 0, upper is 3") {
		t.Error()
	}

	// Type assertion - returns the underlying value and whether the assertion succeeded
	customErr, ok := err.(*CustomError)
	if !(ok && customErr.Lower == 0) {
		t.Error()
	}
}
