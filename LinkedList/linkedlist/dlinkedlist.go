package linkedlist

import "fmt"

// object for int node
type DNode[T comparable] struct {
	data T         // variable hold node value
	next *DNode[T] // next node reference
	prev *DNode[T] // previous node reference
}

// Doubly Linked List implementation
type DoublyLinkedList[T comparable] struct {
	head   *DNode[T] // start node
	tail   *DNode[T] // end node
	length int
}

// constructor for instance
func NewDLL[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// insert new node at begining
func (d *DoublyLinkedList[T]) Prepend(element T) {
	newNode := DNode[T]{data: element}
	newNode.next = d.head
	if d.length == 0 {
		d.head = &newNode
		d.tail = d.head
	} else {
		d.head.prev = &newNode
		d.head = &newNode
	}
	d.length++
}

// insert new node at end
func (d *DoublyLinkedList[T]) Append(element T) {
	newNode := DNode[T]{data: element}
	newNode.prev = d.tail
	d.tail.next = &newNode
	d.tail = &newNode
	d.length++
}

// insert node at given index position
func (d *DoublyLinkedList[T]) InsertPos(element T, pos int) {
	if pos <= 0 {
		d.Prepend(element)
		return
	}
	currentNode := d.head.next
	previousNode := d.head
	currentPos := 1
	for currentNode != nil {
		if currentPos == pos {
			break
		}
		previousNode = currentNode
		currentNode = currentNode.next
		currentPos++
	}
	newNode := DNode[T]{data: element}
	newNode.next = currentNode
	newNode.prev = previousNode
	previousNode.next = &newNode
	if currentNode != nil {
		currentNode.prev = &newNode
	} else {
		d.tail = &newNode
	}
	d.length++
}

// search given value in in list and return boolean value
func (d *DoublyLinkedList[T]) Constains(element T) bool {
	if d.head.data == element || d.tail.data == element {
		return true
	}
	currentNode := d.head
	for currentNode != nil {
		if currentNode.data == element {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// remove the node element from the list
func (d *DoublyLinkedList[T]) Remove(element T) bool {
	if d.length == 0 {
		return false
	}
	if d.head.data == element {
		if d.head.next != nil {
			d.head = d.head.next
		}
		d.head.prev = nil
		d.length--
		return true
	}
	if d.tail.data == element {
		d.tail = d.tail.prev
		d.tail.next = nil
		d.length--
		return true
	}
	currentNode := d.head
	for currentNode != nil {
		if currentNode.data == element {
			break
		}
		currentNode = currentNode.next
	}
	currentNode.next.prev = currentNode.prev
	currentNode.prev.next = currentNode.next
	d.length--
	return false
}

// return the node and it's index position by search the value
// if index out of range will return nil
func (d DoublyLinkedList[T]) Search(element T) (*DNode[T], int) {
	currentNode := d.head
	var currentPos int
	for currentNode != nil {
		if currentNode.data == element {
			return currentNode, currentPos
		}
		currentNode = currentNode.next
		currentPos++
	}
	return nil, 0
}

// show head and tail in console
func (d *DoublyLinkedList[T]) ShowHeadTail() {
	fmt.Println(d.head, d.tail)
}

// display the node values in console, iterate from head to tail
func (d *DoublyLinkedList[T]) IterFromHead() {
	if d.length == 0 {
		fmt.Println("empty list!")
		return
	}
	currentNode := d.head
	for currentNode != nil {
		fmt.Printf("%v, ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("\tlength:", d.length)
}

// display the node values in console, iterate from tail to head
func (d *DoublyLinkedList[T]) IterFromTail() {
	if d.length == 0 {
		fmt.Println("list is empty!")
		return
	}
	currentNode := d.tail
	for currentNode != nil {
		fmt.Printf("%v, ", currentNode.data)
		currentNode = currentNode.prev
	}
	fmt.Println("\tlength:", d.length)
}
