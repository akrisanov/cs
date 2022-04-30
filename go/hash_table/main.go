package main

import "fmt"

const ArraySize = 7

// HashTable with a separate chaining
type HashTable struct {
	array [ArraySize]*bucket
}

// Insert take in a key and adds it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search takes in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete takes in a key and deletes it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// bucket is represented as a linked list
type bucket struct {
	head *bucketNode
}

// bucketNode
type bucketNode struct {
	key  string
	next *bucketNode
}

// insert takes in a key, creates a node with the key and insert the node in the bucket
func (b *bucket) insert(key string) {
	if b.search(key) {
		fmt.Printf("%s already exists\n", key)
	}
	newNode := &bucketNode{key: key}
	newNode.next = b.head
	b.head = newNode
}

// search takes in a key and returns true if the bucket has a key
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take in a key and delete the node from the bucket
func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	currentNode := previousNode.next

	for currentNode != nil {
		if currentNode.key == key {
			currentNode = currentNode.next
		}
		previousNode = currentNode
	}
}

func hash(key string) int {
	sum := 0
	for _, c := range key {
		sum += int(c)
	}
	return sum % ArraySize
}

// InitHashTable creates a new hash table
func InitHashTable() *HashTable {
	hashTable := &HashTable{}
	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}
	return hashTable
}

func main() {
	hashTable := InitHashTable()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _, name := range list {
		hashTable.Insert(name)
	}
	hashTable.Delete("STAN")
	hashTable.Search("TOKEN")
	fmt.Println(hashTable.Search("TOKEN"))
}
