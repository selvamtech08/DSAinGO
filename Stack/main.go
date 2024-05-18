package main

import (
	"fmt"

	"github.com/selvamtech08/DSAinGO/Stack/stack"
)

func main() {
	stack_1 := stack.New[int]()
	stack_1.Push(5)
	stack_1.Push(3)
	stack_1.Push(20)
	stack_1.Push(10)
	stack_1.Push(13)
	stack_1.ShowItems()

	colors := stack.New[string]()
	fmt.Println(colors.IsEmpty())
	colors.Push("green")
	fmt.Println(colors.IsEmpty())
	colors.Push("red")
	colors.Push("blue")
	colors.Push("pink")
	colors.ShowItems()

	fmt.Println(colors.Contains("yellow"))
	fmt.Println(colors.Contains("red"))

	points := stack.New[float32]()
	points.Push(25.30)
	points.Push(85.1)
	points.ShowItems()
	fmt.Println(points.Peek())
	fmt.Println(points.Pop())
	points.ShowItems()
	points.Pop()
	points.ShowItems()
}
