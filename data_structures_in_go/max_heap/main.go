package main

import "fmt"

// MaxHeap struct has a slice that holds an array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("Can't extract from an empty heap")
		return -1
	}
	extracted := h.array[0]
	// move the last item to the root of the heap
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]
	// normalize the heap again
	h.maxHeapifyDown(0)
	// finally, return the extracted element
	return extracted
}

// maxHeapifyUp normalizes the heap from the bottom to the top
func (h *MaxHeap) maxHeapifyUp(index int) {
	// stops when both index and parent(index) are zeroes
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown normalizes the heap from the top to the bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)

	childToCompare := 0
	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else { // if the value of the root element is larger than value of the larger child, we're done
			return
		}
	}
}

// get an index of the parent node
func parent(childIndex int) int {
	return (childIndex - 1) / 2
}

// get an index of the left node
func left(parentIndex int) int {
	return parentIndex*2 + 1
}

// get an index of the right node
func right(parentIndex int) int {
	return parentIndex*2 + 2
}

// swap two array elements by their indexes
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
