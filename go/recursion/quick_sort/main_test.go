package main

import (
	"reflect"
	"testing"
)

func TestQuickSortEmptySlice(t *testing.T) {
	got := QuickSort(slice{})
	want := slice{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestQuickSortOneElementSlice(t *testing.T) {
	got := QuickSort(slice{1})
	want := slice{1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestQuickSort(t *testing.T) {
	got := QuickSort(slice{4, 9, 3, 5})
	want := slice{3, 4, 5, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestQuickSortAlreadySorted(t *testing.T) {
	got := QuickSort(slice{1, 2, 3, 4, 5, 6, 7})
	want := slice{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}
