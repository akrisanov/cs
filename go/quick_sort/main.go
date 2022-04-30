package main

import "fmt"

type slice []int

func QuickSort(s slice) slice {
	if len(s) < 2 {
		return s
	}

	sorted := slice{}

	pivot := s[len(s)-1]
	left := slice{}
	right := slice{}

	for i := 0; i < len(s)-1; i++ {
		if s[i] <= pivot {
			left = append(left, s[i])
		} else {
			right = append(right, s[i])
		}
	}

	sortedLeft := QuickSort(left)
	sortedRight := QuickSort(right)

	sorted = append(sorted, sortedLeft...)
	sorted = append(sorted, pivot)
	sorted = append(sorted, sortedRight...)

	return sorted
}

func main() {
	s := slice{4, 9, 3, 5}
	fmt.Printf("Before sorting: %v\n", s)
	s = QuickSort(s)
	fmt.Printf("After sorting: %v\n", s)
}
