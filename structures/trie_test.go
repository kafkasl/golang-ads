package structures

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	inputs := []string{"mars", "malleus", "mass"}
	outputs := []string{"marx", "ordo", "mass"}

	trie := NewTrie()
	for _, word := range inputs {
		trie.Insert(word)
	}

	//fmt.Printf("Printing trie%v\n", trie)
	trie.root.printWords("")
	// trie.Search(outputs[1])
	// trie.Search(outputs[1])

	var found bool
	for _, word := range outputs {
		found = trie.Search(word)
		fmt.Printf("Word %v found? %v\n", word, found)
	}

}
