package main

import "fmt"

type slice []int

// bubbleSort sorts a slice in place by swapping elements.
// Largest element always bubbles to the top.
func (s *slice) bubbleSort() {
	for i := 0; i < len(*s)-1; i++ {
		fmt.Printf("\nIteration #%d\n", i+1)

		for j := 0; j < len(*s)-1; j++ {
			fmt.Printf("%v -> ", s)
			if (*s)[j] > (*s)[j+1] {
				(*s)[j], (*s)[j+1] = (*s)[j+1], (*s)[j]
			}
			fmt.Printf("%v\n", s)
		}
	}
}

func main() {
	s := &slice{5, 1, 4, 2, 8}
	fmt.Printf("Before sorting: %v\n", s)
	s.bubbleSort()
	fmt.Printf("\nAfter sorting:\t%v\n", s)
}
