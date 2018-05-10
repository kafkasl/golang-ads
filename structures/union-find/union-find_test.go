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

	for i := 0; i < 5; i++ {
		uf := NewUnionFindSet(i)

		// t.Logf("%T", uf1)
		// uf1p := *uf1
		// t.Logf("%T", uf1p)
		// p := uf1p.parent
		// t.Logf("%T", p)
		// t.Logf("%T", uf1.parent)
		// t.Logf("%v", p == uf1.parent)

		if uf.parent != uf {
			t.Fatalf("Parent of new UF Set is not itself\n")
		}
		if uf.rank != 0 {
			t.Fatalf("Rank of new UF Set is not 0, got, %v\n", uf.rank)
		}
		if uf.value != i {
			t.Fatalf("Value of new UF Set is not %v, got, %v\n", i, uf.value)
		}
	}

}

func TestFindAndUnion(t *testing.T) {
	// Basic Find
	value := 1
	uf1 := NewUnionFindSet(value)

	if uf1.Find() != uf1 {
		t.Fatalf("Node returned by Find is not himself.")
	}
	if uf1.Find() != uf1.parent {
		t.Fatalf("Node returned by Find is not his father.")
	}

	// Second find test, all elements in a group have the same parent (designed so the first element is going to be the overall parent)

	uf1 = NewUnionFindSet(1)
	uf2 := NewUnionFindSet(2)

	set1 := make([]*UnionFindSet, 2)
	for i := 0; i < len(set1); i++ {
		set1[i] = NewUnionFindSet(i + 10)
		uf1.Union(set1[i])
	}

	set2 := make([]*UnionFindSet, 2)
	for i := 0; i < len(set2); i++ {
		set2[i] = NewUnionFindSet(i + 20)
		uf2.Union(set2[i])
	}
	for _, node := range set1 {
		if node.Find() != uf1 {
			t.Fatalf("Parent of node %v should be %v but it is %v", node, uf1, node.Find())
		}

		if node.Find() == uf2 {
			t.Fatalf("Parent of node %v should not be %v but it is %v", node, uf2, node.Find())
		}
	}

	for _, node := range set2 {
		if node.Find() != uf2 {
			t.Fatalf("Parent of node %v should be %v but it is %v", node, uf2, node.Find())
		}
		if node.Find() == uf1 {
			t.Fatalf("Parent of node %v should not be %v but it is %v", node, uf1, node.Find())
		}
	}
	for _, node1 := range set1 {
		for _, node2 := range set2 {
			if node1.Find() == node2.Find() {
				t.Fatalf("Nodes %v and %v should not be in the same union-set (aka. they return the same Find())\n", node1, node2)

			}
		}
	}

	// Add a big UF to a recently created
	uf0 := NewUnionFindSet(666)
	uf0.Union(uf1)

	for _, node := range set1 {
		if node.Find() != uf0.Find() {
			t.Fatalf("Nodes %v and %v should have same ancestor. Expected %v but got %v", node, uf0, uf0.Find(), node.Find())
		}

		if node.Find() == uf2.Find() {
			t.Fatalf("Nodes %v and %v should not have same ancestor. Expected %v but got %v", node, uf0, uf0.Find(), node.Find())
		}
	}

	// Ensure that union with one self does not produce a different ancestor
	pAux := uf0.Find()
	uf0.Union(uf0)
	if pAux != uf0.Find() {
		t.Fatalf("Union of node with itself produces different ancestor. Expected %v got %v", pAux, uf0.Find())
	}

}

func TestToString(t *testing.T) {

	set := make([]*UnionFindSet, 2)
	set[0] = NewUnionFindSet(0)

	for i := 1; i < len(set); i++ {
		set[i] = NewUnionFindSet(i + 10)
		set[0].Union(set[i])
	}

	set2 := make([]*UnionFindSet, 2)
	set2[0] = NewUnionFindSet(1)
	for i := 1; i < len(set); i++ {
		set2[i] = NewUnionFindSet(i + 20)
		set2[0].Union(set2[i])
	}

	set[0].Union(set2[0])
	// for _, node := range set {
	// 	t.Logf("%v", node)
	// }
	// for _, node := range set2 {
	// 	t.Logf("%v", node)
	// }

	result := "[2]: 0"
	if set[0].String() != result {
		t.Fatalf("Node %v should be represented as %v", set[0], result)
	}
	result = "[0]: 11 -> [2]: 0"
	if set[1].String() != result {
		t.Fatalf("Node %v should be represented as %v", set[2], result)
	}
	result = "[1]: 1 -> [2]: 0"
	if set2[0].String() != result {
		t.Fatalf("Node %v should be represented as %v", set2[0], result)
	}
	result = "[0]: 21 -> [1]: 1 -> [2]: 0"
	if set2[1].String() != result {
		t.Fatalf("Node %v should be represented as %v", set2[1], result)
	}
}
