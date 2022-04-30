package main

import (
	"reflect"
	"testing"
)

func TestRadixSortEmptySlice(t *testing.T) {
	got := RadixSort([]int{})
	want := []int{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestRadixSort(t *testing.T) {
	got := RadixSort([]int{170, 45, 75, 90, 802, 24, 2, 66})
	want := []int{2, 24, 45, 66, 75, 90, 170, 802}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestRadixSortSliceWithDuplicates(t *testing.T) {
	got := RadixSort([]int{10, 100, 1, 1, 10})
	want := []int{1, 1, 10, 10, 100}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}
