package sortimpl

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

// helper function to compare result slice with expected
func testCompare(a, b []int) error {
	// return error nil size not matched
	if len(a) != len(b) {
		return errors.New("array size not matched")
	}
	// iterate the array and return error if any unmatched found
	for i := range len(a) {
		if a[i] != b[i] {
			return fmt.Errorf("mismatched started from index '%d'", i)
		}
	}
	// if all the elemented matchec above iteration and return nill as success
	return nil
}

func testRandomArray(steps int) []int {
	var array = make([]int, steps)
	// add random value into array
	for i := range steps {
		array[i] = rand.Intn(500)
	}
	return array
}

var testArray []int = testRandomArray(200)

func TestSelection(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{3, 5, 9, 1, 4}, []int{1, 3, 4, 5, 9}},
		{[]int{11, 2, 7, 9, 1}, []int{1, 2, 7, 9, 11}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			result := Selection(tt.input)
			err := testCompare(result, tt.expect)
			if err != nil {
				t.Errorf("err: %s\n", err.Error())
				t.Errorf("sorted %v not matched with expected: %v\n", result, tt.expect)
			}
		})
	}
}

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Selection(slices.Clone(testArray))
	}
}

func TestBubble(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{3, 5, 9, 1, 4}, []int{1, 3, 4, 5, 9}},
		{[]int{11, 2, 7, 9, 1}, []int{1, 2, 7, 9, 11}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			result := Bubble(tt.input)
			err := testCompare(result, tt.expect)
			if err != nil {
				t.Errorf("err: %s\n", err.Error())
				t.Errorf("sorted %v not matched with expected: %v\n", result, tt.expect)
			}
		})
	}
}

func BenchmarkBubble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Bubble(slices.Clone(testArray))
	}
}

func TestInsertion(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{3, 5, 9, 1, 4}, []int{1, 3, 4, 5, 9}},
		{[]int{11, 2, 7, 9, 1}, []int{1, 2, 7, 9, 11}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			result := Insertion(tt.input)
			err := testCompare(result, tt.expect)
			if err != nil {
				t.Errorf("err: %s\n", err.Error())
				t.Errorf("sorted %v not matched with expected: %v\n", result, tt.expect)
			}
		})
	}
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Insertion(slices.Clone(testArray))
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{3, 5, 9, 1, 4}, []int{1, 3, 4, 5, 9}},
		{[]int{11, 2, 7, 9, 1}, []int{1, 2, 7, 9, 11}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			result := Merge(tt.input)
			err := testCompare(result, tt.expect)
			if err != nil {
				t.Errorf("err: %s\n", err.Error())
				t.Errorf("sorted %v not matched with expected: %v\n", result, tt.expect)
			}
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Merge(slices.Clone(testArray))
	}
}
