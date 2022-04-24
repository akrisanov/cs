package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	list := &slice{5, 2, 4, 3, 1}
	sorted := &slice{1, 2, 3, 4, 5}
	list.insertionSort(false)

	if !reflect.DeepEqual(list, sorted) {
		t.Errorf("got %v, wanted %v", list, sorted)
	}
}

func TestInsertionSortMore(t *testing.T) {
	list := &slice{10, 5, 3, 8, 2, 6, 4, 7, 9, 1}
	sorted := &slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list.insertionSort(false)

	if !reflect.DeepEqual(list, sorted) {
		t.Errorf("got %v, wanted %v", list, sorted)
	}
}
func TestBubbleSortSliceWithOneElement(t *testing.T) {
	list := &slice{2}
	sorted := &slice{2}
	list.insertionSort(false)

	if !reflect.DeepEqual(list, sorted) {
		t.Errorf("got %v, wanted %v", list, sorted)
	}
}

func TestBubbleSortEmptySlice(t *testing.T) {
	list := &slice{}
	sorted := &slice{}
	list.insertionSort(false)

	if !reflect.DeepEqual(list, sorted) {
		t.Errorf("got %v, wanted %v", list, sorted)
	}
}
