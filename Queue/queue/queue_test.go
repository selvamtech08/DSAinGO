package queue

import (
	"strconv"
	"testing"
)

func TestContains(t *testing.T) {
	que := New[int]()
	que.Enqueue(5)
	que.Enqueue(3)
	que.Enqueue(2)
	que.Enqueue(10)

	testcases := []struct {
		value  int
		expect bool
	}{
		{5, true},
		{8, false},
		{20, false},
		{10, true},
		{2, true},
		{16, false},
	}
	for _, v := range testcases {
		t.Run(strconv.Itoa(v.value), func(t *testing.T) {
			testresult := que.Contains(v.value)
			if testresult != v.expect {
				t.Errorf("value: %v, expected: %v, result: %v", v.value, v.expect, testresult)
			}
		})
	}
}

func TestEnqueue(t *testing.T) {
	colors := New[string]()
	my_colors := [...]string{"green", "red", "blue", "black"}
	for _, v := range my_colors {
		colors.Enqueue(v)
	}
	testcases := []struct {
		value    string
		expected string
	}{
		{"enqueue1", "green"},
		{"enqueue3", "blue"},
		{"enqueue2", "red"},
		{"enqueue4", "black"},
	}
	for _, v := range testcases {
		t.Run(v.value, func(t *testing.T) {
			result := colors.Dequeue()
			if result != v.expected {
				t.Errorf("expected: %v, got: %v", v.expected, result)
			}
		})
	}
}
