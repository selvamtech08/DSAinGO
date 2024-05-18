package queue

import "fmt"

// Queue data structure
type Queue[T comparable] struct {
	data []T
}

// init queue object
func New[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

// add new element
func (q *Queue[T]) Enqueue(element T) {
	q.data = append(q.data, element)
}

// remove the element at begin of the queue
func (q *Queue[T]) Dequeue() T {
	dequeue := q.data[0]
	q.data = q.data[1:]
	return dequeue
}

// check element is found or not
// return boolean result
func (q *Queue[T]) Contains(element T) bool {
	for _, v := range q.data {
		if v == element {
			return true
		}
	}
	return false
}

// check if queue is empty or not
// return boolean result
func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

// display all the items in the queue list
func (q *Queue[T]) Show() {
	for _, v := range q.data {
		fmt.Printf("%v,", v)
	}
	fmt.Println("\tlength:", len(q.data))
}
