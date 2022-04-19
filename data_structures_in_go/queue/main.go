package main

import "fmt"

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue adds a value at the end
func (s *Queue) Enqueue(i int) {
	s.items = append(s.items, i)
}

// Dequeue removes a value at the beginning and return the removed value
func (s *Queue) Dequeue() int {
	item := s.items[0]
	s.items = s.items[1:]
	return item
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)

	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)

	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
