package main

import (
	"fmt"

	searchtree "github.com/selvamtech08/DSAinGO/BinaryTree/SearchTree"
)

func main() {

	tree := searchtree.New[int]()
	fmt.Println(tree)
	numbers := []int{15, 3, 7, 8, 12, 1, 25, 31, 4, 9}
	for _, v := range numbers {
		tree.Insert(v)
	}
	tree.PreOrder()
	tree.InOrder()
	tree.PostOrder()

	fmt.Println(tree.Search(3))
	fmt.Println(tree.Search(31))
	fmt.Println(tree.Search(15))
	fmt.Println(tree.Search(50))
	fmt.Println(tree.Search(0))

}
