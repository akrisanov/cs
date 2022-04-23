package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := &slice{5, 1, 4, 2, 8}
	sorted := &slice{1, 2, 4, 5, 8}
	list.bubbleSort()

	if !reflect.DeepEqual(list, sorted) {
		t.Errorf("got %v, wanted %v", list, sorted)
	}
}

func TestBubbleSortMore(t *testing.T) {
	list := &slice{10, 5, 3, 8, 2, 6, 4, 7, 9, 1}
	sorted := &slice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list.bubbleSort()

	if !reflect.DeepEqual(list, sorted) {
		t.Errorf("got %v, wanted %v", list, sorted)
	}
}
