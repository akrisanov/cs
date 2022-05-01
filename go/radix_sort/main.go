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
	sliceSize := len(s)

	sorted := make([]int, sliceSize)
	digits_count := make([]int, 10) // for digits 0, 1, 2, 3, 4, 5, 6, 7, 8, 9

	// store count of digit occurrences
	for i := 0; i < sliceSize; i++ {
		digit := (s[i] / exp) % 10
		digits_count[digit]++
	}

	// change digits_count[i] so that digits_count[i] now contains actual position of this digit in sorted[]
	for i := 1; i < 10; i++ {
		digits_count[i] += digits_count[i-1]
	}

	// build the output array
	for i := sliceSize - 1; i >= 0; i-- {
		digit := (s[i] / exp) % 10
		sorted[digits_count[digit]-1] = s[i]
		digits_count[digit]--
	}

	return sorted
}

// RadixSort is a non comparison sort.
// It also can be used for sorting strings (say 26 buckets for letterns), binary and hex numbers.
// NOTE: this implementation doesn't support negative numbers.
func RadixSort(s []int) []int {
	if len(s) == 0 {
		return s
	}

	// find the largest element in the slice to know max number of digits
	max := findMax(s)

	// do counting sort for every digit: instead of passing digit number, exp is passed.
	// exp is 10^exp, where i is current digit number
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
