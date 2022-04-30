package main

import "fmt"

// Stack represents a stack that holds a slice
type Stack struct {
	items []int
}

// Push adds a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop removes a value at the end and return the removed value
func (s *Stack) Pop() int {
	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)

	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)
}
