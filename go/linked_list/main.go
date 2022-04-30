package main

import "fmt"

// node is a linked list building block
type Node struct {
	data int
	next *Node
}

// LinkedList is represented by a heading node and the length
type LinkedList struct {
	head   *Node
	length int
}

// printData prints out data the linked list stores
func (l LinkedList) printData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

// prepend adds a node at the beginning of the linked list
func (l *LinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

// deleteByValue deletes a node by value from the linked list
func (l *LinkedList) deleteByValue(value int) {
	// edge case: immediately return from an empty linked list
	if l.length == 0 {
		return
	}

	// edge case: removing the head
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	// removing a node by value
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	linkedList := LinkedList{}

	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}
	node4 := &Node{data: 11}
	node5 := &Node{data: 7}
	node6 := &Node{data: 2}

	linkedList.prepend(node1)
	linkedList.prepend(node2)
	linkedList.prepend(node3)
	linkedList.prepend(node4)
	linkedList.prepend(node5)
	linkedList.prepend(node6)

	linkedList.printData()

	linkedList.deleteByValue(100)
	linkedList.deleteByValue(2)
	linkedList.deleteByValue(16)
	linkedList.printData()

	emptyLinkedList := LinkedList{}
	emptyLinkedList.deleteByValue(10)
	emptyLinkedList.printData()
}
