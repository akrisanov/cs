package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26 // lower case English letters

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie creates a new trie
func InitTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert takes in a word and adds it to the trie
func (t *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a' //  a is represented as 97
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search takes in a word and return true if that word is in a trie
func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}

func main() {
	testTrie := InitTrie()

	wordsToInsert := []string{
		"maria",
		"mark",
		"andrey",
	}

	for _, v := range wordsToInsert {
		testTrie.Insert(v)
	}

	fmt.Println(testTrie.Search("marc"))
}
