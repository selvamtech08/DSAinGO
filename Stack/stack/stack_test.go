package stack

import (
	"testing"
)

func TestContains(t *testing.T) {
	colors := New[string]()
	colors.Push("green")
	colors.Push("red")
	colors.Push("orange")
	testcases := []struct {
		value    string
		expected bool
	}{
		{"green", true},
		{"red", true},
		{"yellow", false},
		{"pink", false},
		{"white", false},
	}
	for _, v := range testcases {
		t.Run("case: "+v.value, func(t *testing.T) {
			result := colors.Contains(v.value)
			if result != v.expected {
				t.Errorf("case: %s, expect: %v, got: %v", v.value, v.expected, result)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	colors := New[string]()
	result := colors.IsEmpty()
	if !result {
		t.Errorf("expected: %v, got: %v", true, result)
	}
	colors.Push("green")
	result = colors.IsEmpty()
	if result {
		t.Errorf("expected: %v, got: %v", true, result)
	}
}
