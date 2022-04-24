package main

import "fmt"

type slice []int

func (s *slice) insertionSort(logSwaps bool) {
	for i := 1; i < len(*s); i++ {
		cursor := i

		for j := i - 1; j >= 0; j-- {
			if (*s)[cursor] < (*s)[j] {
				(*s)[cursor], (*s)[j] = (*s)[j], (*s)[cursor]
				cursor = j
				if logSwaps {
					fmt.Printf("Iteration #%d - %v\n", i, s)
				}
			}
		}
	}
}

func main() {
	s := &slice{5, 2, 4, 3, 1}
	fmt.Printf("Before sorting: %v\n", s)
	s.insertionSort(false)
	fmt.Printf("After sorting:\t%v\n", s)
}
