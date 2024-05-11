package linkedlist

import (
	"errors"
	"fmt"
)

// object for int node
type Node struct {
	data int   // variable hold node value
	next *Node // next node reference
}

// Single Linked list implemention
type SingleLinkedList struct {
	head   *Node // pointing to root(head) node for SLL object
	length int   // track total nodes in list
}

// insert given node at beginning and it will be new head
func (s *SingleLinkedList) Prepand(element int) {
	newNode := Node{data: element}
	newNode.next = s.head
	s.head = &newNode
	s.length++
}

// return head node
// if list is empty return dummy node with error
func (s *SingleLinkedList) Head() (Node, error) {
	if s.head == nil { // nil assignment will panic, so added error object as return
		return Node{}, errors.New("empty list")
	}
	return *s.head, nil
}

// append node at end of list
func (s *SingleLinkedList) Append(element int) {
	if s.head == nil { // check if list is empty, if Yes call prepand
		s.Prepand(element)
		return
	}
	// traverse list still end and insert new element
	currentNode := s.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	newNode := Node{data: element}
	currentNode.next = &newNode
	s.length++
}

// show all nodes in the list
func (s *SingleLinkedList) Show() {
	// Note: here we use list point, careful while doing operations.
	// Any assignment to s object, will affect original list
	currentNode := s.head
	// traverse the nodes and show the it's data
	for currentNode != nil {
		fmt.Printf("%d, ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("\tLength:", s.length)
}

// search node in the list and return boolean value
func (s *SingleLinkedList) Search(element int) bool {
	if s.head == nil { // empty list
		return false
	}
	currentNode := s.head
	// traverse list and find the node
	for currentNode.next != nil {
		if currentNode.data == element {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// check if list is empty and return true or false
func (s *SingleLinkedList) IsEmpty() bool {
	return s.head == nil
}

// remove given node from list, return true if removed or false where node not found
func (s *SingleLinkedList) Remove(element int) bool {
	if s.head == nil { // empty list
		return false
	}
	if s.head.data == element { // remove node if it is head
		s.head = s.head.next
		s.length--
		return true
	}
	currentNode := s.head
	previousNode := currentNode
	for currentNode != nil { // traverse until match found
		if currentNode.data == element { // compare current node data with given value
			previousNode.next = currentNode.next // ressign current previos's next to current's next
			s.length--                           // above step technically unplugin the current(matched) node from the list
			return true
		}
		previousNode = currentNode
		currentNode = currentNode.next
	}
	return false // return false if given value not found in the list
}

// inserted new node in given position,
// return true for successfully update otherwise return false
func (s *SingleLinkedList) InsertPos(element, pos int) {
	// two possible cases here will match
	// 1. empty list, 2. position at 0
	if s.head == nil || pos <= 0 {
		s.Prepand(element)
		return
	}
	// inserting after the head
	currentNode := s.head.next
	previousNode := s.head
	currentPos := 1
	for currentNode != nil {
		if currentPos == pos {
			break // break the foor loop if index match found
		}
		previousNode = currentNode
		currentNode = currentNode.next
		currentPos++
	}
	// Note: if index position not found then insert at end
	newNode := &Node{data: element}
	previousNode.next = newNode
	// if position is not at end
	if currentNode != nil {
		newNode.next = currentNode.next
	}
	s.length++
}

// return number of nodes contains same value and return 0 if value not found
func (s *SingleLinkedList) Count(value int) int {
	currentNode := s.head
	count := 0
	for currentNode != nil {
		if currentNode.data == value {
			count++
		}
		currentNode = currentNode.next
	}
	return count
}

// return mininum value in the list
// return false if list is empty
func (s *SingleLinkedList) Min() (int, bool) {
	// list will accept signed int, so can't return -1 if list is empty
	if s.head == nil {
		return -1, false
	}
	minValue := s.head.data
	currentNode := s.head.next
	for currentNode != nil {
		if currentNode.data < minValue {
			minValue = currentNode.data
		}
		currentNode = currentNode.next
	}
	return minValue, true
}

// return maxinum value in the list, return
// return false if list is empty
func (s *SingleLinkedList) Max() (int, bool) {
	if s.head == nil {
		return -1, false
	}
	maxValue := s.head.data
	currentNode := s.head.next
	for currentNode != nil {
		if currentNode.data > maxValue {
			maxValue = currentNode.data
		}
		currentNode = currentNode.next
	}
	return maxValue, true
}

// TODO
// sort the list and return new list
func (s *SingleLinkedList) Sort() *SingleLinkedList {
	newList := SingleLinkedList{}
	return &newList
}
