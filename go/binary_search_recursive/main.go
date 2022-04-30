package main

import (
	"fmt"
	"math"
	"sort"
)

// BinarySearch looks up for a key in an array and returns an index of the found key.
func BinarySearch(key int, s []int) int {
	if len(s) == 0 {
		return -1
	}

	middle := int(math.Floor(float64(len(s)) / 2))

	if key == s[middle] {
		return middle
	} else if key < s[middle] {
		return BinarySearch(key, s[:middle])
	} else {
		return BinarySearch(key, s[middle+1:])
	}
}

func main() {
	s := []int{10, 3, 45, 4, 0, -1, 2, 77, 2, 3}
	sort.Ints(s)

	key := 2
	result := BinarySearch(key, s)

	if result != -1 {
		fmt.Printf("%d has been found at the index %d in the array of %v\n", key, result, s)
	} else {
		fmt.Printf("%d is not found in the array of %v", key, s)
	}
}
