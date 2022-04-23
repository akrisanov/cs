package main

import "fmt"

type slice []int

// bubbleSort sorts a slice in place by swapping elements.
// Largest element always bubbles to the top.
func (s *slice) bubbleSort() {
	swapped := true
	iterations := 1

	for swapped {
		swapped = false

		fmt.Printf("\nIteration #%d\n", iterations)

		for i := 0; i < len(*s)-1; i++ {
			fmt.Printf("%v -> ", s)
			if (*s)[i] > (*s)[i+1] {
				(*s)[i], (*s)[i+1] = (*s)[i+1], (*s)[i]
				swapped = true
			}
			fmt.Printf("%v\n", s)
		}

		iterations++
	}
}

func main() {
	s := &slice{5, 1, 4, 2, 8}
	fmt.Printf("Before sorting: %v\n", s)
	s.bubbleSort()
	fmt.Printf("\nAfter sorting:\t%v\n", s)
}
