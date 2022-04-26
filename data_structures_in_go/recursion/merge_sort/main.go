package main

import "fmt"

type slice []int

func MergeSort(s slice) slice {
	if len(s) < 2 {
		return s
	}
	m := len(s)/2 + len(s)%2 // math.Floor()

	left := MergeSort(s[:m])
	right := MergeSort(s[m:])
	sorted := Merge(left, right)

	return sorted
}

// Merge creates a new sorted slice out of two sorted slices
func Merge(left, right slice) slice {
	sorted := slice{}

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}

	sorted = append(sorted, left...)
	sorted = append(sorted, right...)

	return sorted
}

func main() {
	s := slice{1, 5, 7, 4, 2, 3, 6}
	fmt.Printf("Before sorting: %v\n", s)
	s = MergeSort(s)
	fmt.Printf("After sorting: %v\n", s)
}
