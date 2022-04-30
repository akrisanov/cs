package main

import (
	"reflect"
	"testing"
)

func TestMergeSortEmptySlice(t *testing.T) {
	got := MergeSort(slice{})
	want := slice{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestMergeSort(t *testing.T) {
	got := MergeSort(slice{1, 5, 7, 4, 2, 3, 6})
	want := slice{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestMergeSortSliceWithDuplicates(t *testing.T) {
	got := MergeSort(slice{1, 1, 1, 1, 1})
	want := slice{1, 1, 1, 1, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}

func TestMergeSortSliceWithNegativeNumber(t *testing.T) {
	got := MergeSort(slice{-1, -10, 0, 5, 2})
	want := slice{-10, -1, 0, 2, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %d, got %d\n", want, got)
	}
}
