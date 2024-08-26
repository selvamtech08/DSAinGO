package main

import (
	"fmt"
	"slices"

	"github.com/selvamtech08/sort/sortimpl"
)

func main() {

	numbers := []int{15, 3, 18, 19, 5, 6}

	fmt.Println("\nselection sort")
	result := sortimpl.Selection(slices.Clone(numbers))
	fmt.Println("orgrinal:", numbers)
	fmt.Println("sorted:", result)

	fmt.Println("\nbubble sort")
	result = sortimpl.Bubble(slices.Clone(numbers))
	fmt.Println("orgrinal:", numbers)
	fmt.Println("sorted:", result)

	fmt.Println("\ninsertion sort")
	result = sortimpl.Insertion(slices.Clone(numbers))
	fmt.Println("orgrinal:", numbers)
	fmt.Println("sorted:", result)

	fmt.Println("\nmerge sort")
	result = sortimpl.Merge(slices.Clone(numbers))
	fmt.Println("orgrinal:", numbers)
	fmt.Println("sorted:", result)

}
