package main

import (
	"sort"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	s := []int{10, 3, 45, 4, 0, -1, 2, 77, 2, 3}
	sort.Ints(s)

	key := 2
	got := BinarySearch(key, s)
	want := 2

	if got != want {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestBinarySearchEmptySlice(t *testing.T) {
	s := []int{}

	key := 2
	got := BinarySearch(key, s)
	want := -1

	if got != want {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestBinarySearchOneItemSlice(t *testing.T) {
	s := []int{1}

	key := 1
	got := BinarySearch(key, s)
	want := 0

	if got != want {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	s := []int{10, 3, 45, 4, 0, -1, 2, 77, 2, 3}
	sort.Ints(s)

	key := 999
	got := BinarySearch(key, s)
	want := -1

	if got != want {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}
