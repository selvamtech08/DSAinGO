package stack

import "fmt"

// implement stack data type
type Stack[T comparable] struct {
	data []T // items stored in slice
}

// init stack object
func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

// add element at end of stack list
func (s *Stack[T]) Push(element T) {
	s.data = append(s.data, element)
}

// return the last element
func (s *Stack[T]) Peek() T {
	return s.data[len(s.data)-1]
}

// remove last element and return
func (s *Stack[T]) Pop() T {
	popValue := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return popValue
}

// check given value present in the stack
// return true if value found else return false
func (s *Stack[T]) Contains(element T) bool {
	for _, v := range s.data {
		if element == v {
			return true
		}
	}
	return false
}

// show all the elements in the stack list
func (s *Stack[T]) ShowItems() {
	if len(s.data) == 0 {
		fmt.Println("stack is empty!")
		return
	}
	for _, v := range s.data {
		fmt.Printf("%v, ", v)
	}
	fmt.Println("\tlength:", len(s.data))
}

// check if stack is empty or not
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
