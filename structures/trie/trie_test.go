package trie

import (
	"fmt"
	"testing"
)

func PrintFatal(t *testing.T, expected fmt.Stringer, got fmt.Stringer) {
	t.Fatalf("Expected: \n%v, but got: \n%v\n", expected, got)
}

func contains(list []string, target string) bool {
	for _, element := range list {
		if target == element {
			return true
		}
	}
	return false
}

func TestWords(t *testing.T) {
	words := []string{"mars", "malleus", "me", "stavro", "dent", "wiggin"}

	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}

	resultWords := trie.Words()

	if len(resultWords) != len(words) {
		t.Fatalf("Expected: \n%v, but got: \n%v\n", words, resultWords)
	}
	for i := 0; i < len(words); i++ {
		if !contains(words, resultWords[i]) {
			t.Fatalf("Word: %v not reported to be in %v\n", resultWords[i], words)
		}
	}

}

func TestSearch(t *testing.T) {
	inputs := []string{"mars", "malleus", "me", "mass"}
	searches := []string{"marx", "ordo", "mass", "hello", "malleus", "me"}
	outputs := []bool{false, false, true, false, true, true}

	trie := NewTrie()
	for _, word := range inputs {
		trie.Insert(word)
	}

	var found bool
	for i, word := range searches {
		found = trie.Search(word)
		if found != outputs[i] {
			t.Fatalf("Expected: %v, but got: %v, for word %v\n", found, outputs[i], word)
		}
	}
}

func TestPrint(t *testing.T) {
	inputs := []string{"mars", "malleus", "me"}
	output := `Trie:
 └─ m
    ├─ a
    │  ├─ l
    │  │  └─ l
    │  │     └─ e
    │  │        └─ u
    │  │           └─ s
    │  └─ r
    │     └─ s
    └─ e
`
	trie := NewTrie()
	for _, word := range inputs {
		trie.Insert(word)
	}

	resultOutput := trie.String()

	if resultOutput != output {
		t.Fatalf("Expected output: \n%v\n but got: \n%v\n", output, resultOutput)
	}
}
