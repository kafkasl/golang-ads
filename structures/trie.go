package structures

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

func (tn TrieNode) printWords(prefix string) {
	if tn.endOfWord {
		fmt.Printf("Appending word: %v\n", prefix)
	}
	if len(tn.children) > 0 {
		// fmt.Printf("Children: %v", tn.children)
		for key, child := range tn.children {
			child.printWords(prefix + string(key))
		}
	}
}

func (tn TrieNode) recursiveToString(father TrieNode, prefix string, c chan string) {
	if tn.endOfWord {
		// *words = append(*words, prefix)
		fmt.Printf("Appending word: %v\n", prefix)
		c <- prefix
	}
	if len(tn.children) > 0 {
		for key, child := range tn.children {
			child.recursiveToString(father, prefix+string(key), c)
		}
	}
	// if &tn == &father {
	fmt.Println("closing channel")
	close(c)
	// }
}

func (tn TrieNode) String() string {
	// fmt.Printf("Printing trieNode%v\n", tn)
	words := []string{}
	c := make(chan string)
	tn.recursiveToString(tn, "", c)
	for word := range c {
		words = append(words, word)
		fmt.Printf("Received word: %v", word)
	}

	fmt.Printf("Printing words%v\n", words)
	return strings.Join(words, "\n")
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	tn := TrieNode{make(map[rune]*TrieNode), false}
	return Trie{&tn}
}

func (t Trie) String() string {
	return t.root.String()
}

func (t Trie) Insert(key string) {
	currentNode := t.root
	for _, letter := range key {
		if val, ok := currentNode.children[letter]; ok {
			currentNode = val
		} else {
			nn := &TrieNode{make(map[rune]*TrieNode), false}
			currentNode.children[letter] = nn
			currentNode = nn
		}
	}
	currentNode.endOfWord = true
}

func (t Trie) Search(key string) bool {
	currentNode := t.root
	for _, letter := range key {
		// fmt.Println("Proceeding to test key")
		// fmt.Printf("Node %v", currentNode)
		if val, ok := currentNode.children[letter]; ok {
			currentNode = val
		} else {
			return false
		}
	}
	return currentNode.endOfWord
}

// def _charToIndex(self,ch):
//
//     // private helper function
//     // Converts key current character into index
//     // use only 'a' through 'z' and lower case
//
//     return ord(ch)-ord('a')
//
//
// def insert(self,key):
//
//     // If not present, inserts key into trie
//     // If the key is prefix of trie node,
//     // just marks leaf node
//     pCrawl = self.root
//     length = len(key)
//     for level in range(length):
//         index = self._charToIndex(key[level])
//
//         // if current character is not present
//         if not pCrawl.children[index]:
//             pCrawl.children[index] = self.getNode()
//         pCrawl = pCrawl.children[index]
//
//     // mark last node as leaf
//     pCrawl.isEndOfWord = True
//
// def search(self, key):
//
//     // Search key in the trie
//     // Returns true if key presents
//     // in trie, else false
//     pCrawl = self.root
//     length = len(key)
//     for level in range(length):
//         index = self._charToIndex(key[level])
//         if not pCrawl.children[index]:
//             return False
//         pCrawl = pCrawl.children[index]
//
//     return pCrawl != None and pCrawl.isEndOfWord
