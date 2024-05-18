package main

import (
	"fmt"

	"github.com/selvamtech08/DSAinGO/linkedlist"
)

func main() {
	// sll_1 := linkedlist.SingleLinkedList{}
	// fmt.Println(sll_1)
	// sll_1.Prepand(8)
	// sll_1.Prepand(5)
	// sll_1.Prepand(15)
	// sll_1.Prepand(2)
	// fmt.Println(sll_1)
	// fmt.Println(sll_1.Head())
	// sll_1.Append(23)
	// sll_1.Show()
	// sll_1.Append(35)
	// sll_1.Prepand(4)
	// sll_1.Show()
	// fmt.Println(sll_1.Search(5))
	// fmt.Println(sll_1.Search(100))
	// fmt.Println(sll_1.IsEmpty())

	// sll_2 := linkedlist.SingleLinkedList{}
	// head, err := sll_2.Head()
	// if err == nil {
	// 	fmt.Println(head)
	// }
	// fmt.Println(sll_2.IsEmpty())
	// sll_2.Append(1)
	// sll_2.Append(5)
	// sll_2.Append(3)
	// sll_2.Append(8)
	// fmt.Println(sll_2.Head())
	// sll_2.Show()
	// fmt.Println(sll_2.Remove(8))
	// fmt.Println(sll_2.Remove(80))
	// fmt.Println(sll_2.Remove(1))
	// fmt.Println(sll_2.Remove(3))
	// fmt.Println(sll_2.Remove(5))
	// fmt.Println(sll_2.Remove(5))
	// sll_2.Show()
	// sll_2.InsertPos(8, 3)
	// sll_2.Show()
	// sll_2.InsertPos(5, -3)
	// sll_2.InsertPos(12, 100)
	// sll_2.InsertPos(3, 100)
	// sll_2.InsertPos(45, 100)
	// sll_2.InsertPos(45, 100)
	// sll_2.Show()
	// fmt.Println(sll_2.Count(45))
	// fmt.Println(sll_2.Count(12))
	// fmt.Println(sll_2.Count(2))

	// myMin, ok := sll_2.Max()
	// if ok {
	// 	fmt.Println("Min value is:", myMin)
	// }
	// fmt.Println(sll_2.Max())
	// fmt.Println(sll_2.Min())

	// sll_3 := linkedlist.SingleLinkedList{}
	// fmt.Println(sll_3.Max())

	// doubly linked list

	dll_1 := linkedlist.NewDLL[int]()
	dll_1.IterFromHead()
	dll_1.Prepend(5)
	dll_1.Prepend(1)
	dll_1.Append(8)
	dll_1.Append(13)
	dll_1.IterFromHead()
	dll_1.InsertPos(17, 2)
	dll_1.InsertPos(3, -2)
	dll_1.InsertPos(3, 20)
	dll_1.IterFromTail()
	dll_1.IterFromHead()

	fmt.Println(dll_1.Search(5))
	fmt.Println(dll_1.Search(15))
	dll_1.Remove(17)
	dll_1.IterFromHead()
	dll_1.Remove(3)
	dll_1.IterFromHead()
	//dll_1.Remove(3)
	//dll_1.Remove(13)
	//dll_1.Remove(8)
	//dll_1.Remove(5)
	//dll_1.Remove(1)
	dll_1.ShowHeadTail()
	dll_1.Remove(1)
	dll_1.ShowHeadTail()
	dll_1.Remove(5)
	dll_1.ShowHeadTail()
	dll_1.Remove(8)
	dll_1.ShowHeadTail()
	dll_1.Remove(13)
	dll_1.ShowHeadTail()
	dll_1.IterFromHead()
}
