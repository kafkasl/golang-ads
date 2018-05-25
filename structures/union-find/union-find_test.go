package union_find

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

	uf := NewUnionFindSet(5)
	for i := 0; i < 5; i++ {

		if uf.g[i] != i {
			t.Fatalf("Element not pointing to itself\n")
		}
	}

}

func TestFindAndUnion(t *testing.T) {
	// Basic Find
	value := 1
	uf1 := NewUnionFindSet(5)

	if uf1.Find(value) != value {
		t.Fatalf("Node returned by Find is not himself.")
	}

	// Second find test, all elements in a group have the same parent (designed so the first element is going to be the overall parent)

	set := NewUnionFindSet(10)

	for i := 0; i < 4; i++ {
		set.Union(set.g[i], set.g[i+1])
	}
	for i := 5; i < 9; i++ {
		set.Union(set.g[i], set.g[i+1])
	}

	for i := 0; i < 5; i++ {
		for j := 5; j < 10; j++ {
			if set.Find(i) == set.Find(j) {
				t.Fatalf("Nodes %v and %v should not be in the same union-set (aka. they return the same Find())\n", i, j)

			}
		}
	}

	// Ensure that union with one self does not produce a different ancestor
	pAux := set.Find(0)
	set.Union(set.g[0], set.g[0])
	if pAux != set.Find(set.g[0]) {
		t.Fatalf("Union of node with itself produces different ancestor. Expected %v got %v", pAux, set.Find(set.g[0]))
	}

}

func TestToString(t *testing.T) {

	set := NewUnionFindSet(10)

	output := set.String()
	correct := "0 1 2 3 4 5 6 7 8 9 "

	if output != correct {
		t.Fatalf("Output of String() is not correct. Expected %v, Got: %v", correct, output)
	}
}
