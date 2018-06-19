package union_find

import (
	"fmt"
	"testing"
)

func TestInitializerwRP(t *testing.T) {

	uf := make([]*UnionFindSetwRP, 5)
	for i := 0; i < 5; i++ {
		uf[i] = NewUnionFindSetwRP(i)
		if uf[i].Find() != uf[i] {
			t.Fatalf("Element not pointing to itself\n")
		}
	}

}

func TestFindAndUnionwRP(t *testing.T) {
	// Basic Find
	value := 1
	uf1 := NewUnionFindSetwRP(value)

	if uf1.Find().value != value {
		t.Fatalf("Node returned by Find is not himself.")
	}

	// Second find test, all elements in a group have the same parent (designed so the first element is going to be the overall parent)

	set := make([]*UnionFindSetwRP, 10)
	for i := 0; i < 10; i++ {
		set[i] = NewUnionFindSetwRP(i)
	}

	for i := 0; i < 4; i++ {
		set[i].Union(set[i+1])
	}
	for i := 8; i >= 5; i-- {
		set[i].Union(set[i+1])
	}

	for i := 0; i < 5; i++ {
		for j := 5; j < 10; j++ {
			if set[i].Find() == set[j].Find() {
				t.Fatalf("Nodes %v and %v should not be in the same union-set (aka. they return the same Find())\n", i, j)

			}
		}
	}

	// Ensure that union with one self does not produce a different ancestor
	pAux := set[0].Find()
	set[0].Union(set[0])
	if pAux != set[0].Find() {
		t.Fatalf("Union of node with itself produces different ancestor. Expected %v got %v", pAux, set[0].Find())
	}

}

func TestToStringSingleElement(t *testing.T) {

	set := NewUnionFindSetwRP(10)

	output := set.String()
	correct := "[0]: 10"

	if output != correct {
		t.Fatalf("Output of String() is not correct. Expected %v, Got: %v", correct, output)
	}
}

func TestToStringwRP(t *testing.T) {

	set := make([]*UnionFindSetwRP, 4)
	for i := 0; i < 4; i++ {
		set[i] = NewUnionFindSetwRP(i)
	}
	set[1].Union(set[0])
	set[3].Union(set[2])
	// set[4].Union(set[2])

	set[2].Union(set[1])
	output := ""
	for i := 0; i < 4; i++ {
		output += fmt.Sprintf("\n%v", set[i])
	}
	correct := `
[0]: 0 -> [1]: 1 -> [2]: 3
[1]: 1 -> [2]: 3
[0]: 2 -> [2]: 3
[2]: 3`

	if output != correct {
		t.Fatalf("Output of String() is not correct. Expected %v, Got: %v", correct, output)
	}
}
