package main

import (
	"fmt"
	"math"
)

// findMax returns the largest element in the slice
func findMax(s []int) int {
	max := s[0]
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

// countSort do counting sort according to the digit represented by exp
func countSort(s []int, exp int) []int {
	sorted := make([]int, len(s))
	count := make([]int, 10) // for digits 0, 1, 2, 3, 4, 5, 6, 7, 8, 9

	// store count of occurrences
	for i := 0; i < len(s); i++ {
		count[(s[i]/exp)%10]++
	}

	// change count[i] so that count[i] now contains actual position of this digit in sorted[]
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// build the output array
	for i := len(s) - 1; i >= 0; i-- {
		sorted[count[(s[i]/exp)%10]-1] = s[i]
		count[(s[i]/exp)%10]--
	}

	return sorted
}

// RadixSort is non comparison sort.
// NOTE: this implementation doesn't support negative numbers.
func RadixSort(s []int) []int {
	if len(s) == 0 {
		return s
	}

	// find the largest element in the slice to know max number of digits
	max := findMax(s)

	// do counting sort for every digit: instead of passing digit number, exp is passed.
	// exp is 10^i, where i is current digit number
	for exp := 1; math.Floor(float64(max)/float64(exp)) > 0; exp *= 10 {
		s = countSort(s, exp)
	}

	return s
}

func main() {
	s := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Printf("Before sorting: %v\n", s)
	s = RadixSort(s)
	fmt.Printf("After sorting: %v\n", s)
}
