package structures

import (
	"fmt"
	"testing"
)

func PrintFatal(t *testing.T, expected fmt.Stringer, got fmt.Stringer) {
	t.Fatalf("Expected: \n%v, but got: \n%v\n", expected, got)
}

//
// func contains(list []string, target string) bool {
// 	for _, element := range list {
// 		if target == element {
// 			return true
// 		}
// 	}
// 	return false
// }
//
func TestInitializer(t *testing.T) {

	value := 1
	uf1 := NewUnionFindSet(value)

	// t.Logf("%T", uf1)
	// uf1p := *uf1
	// t.Logf("%T", uf1p)
	// p := uf1p.parent
	// t.Logf("%T", p)
	// t.Logf("%T", uf1.parent)
	// t.Logf("%v", p == uf1.parent)

	if uf1.parent != uf1 {
		t.Fatalf("Parent of new UF Set is not itself\n")
	}
	if uf1.rank != 0 {
		t.Fatalf("Rank of new UF Set is not 0, got, %v\n", uf1.rank)
	}
	if uf1.value != value {
		t.Fatalf("Value of new UF Set is not %v, got, %v\n", value, uf1.value)
	}
}

func TestFind(t *testing.T) {
	// Basic Find
	value := 1
	uf1 := NewUnionFindSet(value)

	if uf1.Find() != uf1 {
		t.Fatalf("Node returned by Find is not himself.")
	}
	if uf1.Find() != uf1.parent {
		t.Fatalf("Node returned by Find is not his father.")
	}
}

func TestToString(t *testing.T) {
	uf0 := NewUnionFindSet(0)
	uf1 := NewUnionFindSet(1)
	uf2 := NewUnionFindSet(2)
	uf3 := NewUnionFindSet(3)

	uf0.Union(uf1)
	uf0.Union(uf2)
	uf0.Union(uf3)
	t.Logf("%v", uf0)
	t.Logf("%v", uf1)
	t.Logf("%v", uf2)
	t.Logf("%v", uf3)

}

//
// 	resultWords := trie.Words()
//
// 	if len(resultWords) != len(words) {
// 		t.Fatalf("Expected: \n%v, but got: \n%v\n", words, resultWords)
// 	}
// 	for i := 0; i < len(words); i++ {
// 		if !contains(words, resultWords[i]) {
// 			t.Fatalf("Word: %v not reported to be in %v\n", resultWords[i], words)
// 		}
// 	}
//
// }
//
// func TestSearch(t *testing.T) {
// 	inputs := []string{"mars", "malleus", "me", "mass"}
// 	searches := []string{"marx", "ordo", "mass", "hello", "malleus", "me"}
// 	outputs := []bool{false, false, true, false, true, true}
//
// 	trie := NewTrie()
// 	for _, word := range inputs {
// 		trie.Insert(word)
// 	}
//
// 	var found bool
// 	for i, word := range searches {
// 		found = trie.Search(word)
// 		if found != outputs[i] {
// 			t.Fatalf("Expected: %v, but got: %v, for word %v\n", found, outputs[i], word)
// 		}
// 	}
// }
//
// func TestPrint(t *testing.T) {
// 	inputs := []string{"mars", "malleus", "me"}
// 	output := `Trie:
//  └─ m
//     ├─ a
//     │  ├─ l
//     │  │  └─ l
//     │  │     └─ e
//     │  │        └─ u
//     │  │           └─ s
//     │  └─ r
//     │     └─ s
//     └─ e
// `
// 	trie := NewTrie()
// 	for _, word := range inputs {
// 		trie.Insert(word)
// 	}
//
// 	resultOutput := trie.String()
//
// 	if resultOutput != output {
// 		t.Fatalf("Expected output: \n%v\n but got: \n%v\n", output, resultOutput)
// 	}
// }
