package main

import "fmt"

type slice []int

func MergeSort(s slice) slice {
	return slice{}
}

// merge creates a new sorted slice out of two sorted slices
func merge(left, right slice) slice {
	if len(left) < len(right) {
		err := fmt.Errorf("Left slice is empty when the right slice is not")
		fmt.Println(err.Error())
	}

	sorted := slice{}
	for i := 0; i < len(left); i++ {
		if i > len(right)-1 {
			sorted = append(sorted, left[i])
			break
		}

		if left[i] < right[i] {
			sorted = append(sorted, left[i], right[i])
		} else {
			sorted = append(sorted, right[i], left[i])
		}
	}
	return sorted
}

func main() {
	s := slice{1, 5, 7, 4, 2, 3, 6}
	fmt.Printf("Before sorting: %v\n", s)
	s = MergeSort(s)
	fmt.Printf("After sorting: %v\n", s)
}
