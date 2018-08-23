package c03datastructures

import (
	"sort"
	"testing"
)

func TestArray(t *testing.T) {
	// Array size is fixed
	// Initialise array of size 4
	scores := [4]int{10, 11, 12, 13}

	// Use len to get the length of the array
	if !(len(scores) == 4) {
		t.Error()
	}

	// Use range to iterate over various data structures
	for index, value := range scores {
		if value != index+10 {
			t.Error()
		}
	}
}

func TestSlice(t *testing.T) {
	// A slice is a lightweight structure that wraps and represents a portion of an array
	scores := []int{10, 20, 30, 40}

	if !(scores[3] == 40) {
		t.Error()
	}
}

func TestSliceLengthCapacity(t *testing.T) {
	// A slice is a descriptor of an array segment. It consists of a pointer to the array, the
	// length of the segment, and its capacity (the maximum length of the segment).
	// Length: number of elements referred to by the slice
	// Capacity: number of elements in the underlying array

	// Another way to create a slice, of length and capacity 10
	scores := make([]int, 10)
	scores[1] = 20

	if !(scores[0] == 0 && scores[1] == 20 && len(scores) == 10 && cap(scores) == 10) {
		t.Error()
	}
}

func TestSliceAppend(t *testing.T) {
	// Slice of length 0 and capacity 10
	scores := make([]int, 0, 10)

	if !(len(scores) == 0) {
		t.Error()
	}

	// Appending to a slice of length 0 sets the first element
	// You cannot set it directly, e.g., scores[0] = 10 because the length is 0
	scores = append(scores, 10)

	if !(scores[0] == 10 && len(scores) == 1) {
		t.Error()
	}
}

func TestReslice(t *testing.T) {
	scores := make([]int, 0, 10)

	// Re-slicing is done by specifying a half-open range with two indices separated by a colon
	// The start and end indices of a slice expression are optional; they default to zero and the
	// slice's length respectively

	// Create a slice including elements 0 through 7
	scores = scores[:8]
	scores[7] = 80

	// Slicing does not copy the slice's data. It creates a new slice value that points to the
	// original array.
	// Therefore cap(scores) is still 10, as it was originally created
	if !(len(scores) == 8 && cap(scores) == 10) {
		t.Error()
	}
}

func TestModifyReslice(t *testing.T) {
	scores := []int{10, 20, 30}

	slicedScores := scores[1:cap(scores)]

	if !(slicedScores[1] == 30) {
		t.Error()
	}

	// Modifying the elements (not the slice itself) of a re-slice modifies the elements of the
	// original slice
	slicedScores[1] = 300

	if !(scores[2] == 300) {
		t.Error()
	}
}

func TestGrowSlice(t *testing.T) {
	scores := make([]int, 0, 3)

	for i := 0; i < 10; i++ {
		scores = append(scores, i)
	}

	// Arrays are grown with a 2x algorithm: 3 to 6 to 12
	if !(cap(scores) == 12) {
		t.Error()
	}
}

func TestRemoveValueFromUnsortedSlice(t *testing.T) {
	scores := []int{10, 40, 30, 20, 50}
	scores = removeAtIndex(scores, 2)

	if !(scores[2] == 50 && len(scores) == 4 && cap(scores) == 5) {
		t.Error()
	}
}

func removeAtIndex(source []int, index int) []int {
	lastIndex := len(source) - 1

	// Swap value to be removed with the last value
	source[index], source[lastIndex] = source[lastIndex], source[index]
	// then return a slice up to and excluding the swapped value to be removed
	return source[:lastIndex]
}

func TestCopySlice(t *testing.T) {
	scores := []int{30, 50, 10, 20, 40}
	smallest := extractSmallest(scores, 3)

	if !(len(smallest) == 3 && cap(smallest) == 3) {
		t.Error()
	}
	if !(smallest[0] == 10 && smallest[1] == 20 && smallest[2] == 30) {
		t.Error()
	}
}

func extractSmallest(source []int, amount int) []int {
	sort.Ints(source)

	result := make([]int, amount)
	copy(result, source[:amount])

	return result
}
