package main

import "fmt"

func main() {

	// test
	stack := NewStack()

	stack.push(1)
	stack.push(2)
	stack.push(3)

	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
