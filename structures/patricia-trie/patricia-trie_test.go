package structures

import (
	"fmt"
	"testing"
)

func PrintFatal(t *testing.T, expected fmt.Stringer, got fmt.Stringer) {
	t.Fatalf("Expected: \n%v, but got: \n%v\n", expected, got)
}

// func TestSearch(t *testing.T) {
// 	inputs := []uint64{0, 1, 2, 3, 4}
// 	// searches := []string{"marx", "ordo", "mass", "hello", "malleus", "me"}
// 	// outputs := []bool{false, false, true, false, true, true}
//
// 	// var pt PatriciaTrie
// 	pt := NewPatriciaTrie()
//
// 	pt.Insert(inputs[0])
// 	if !pt.Search(inputs[0]) {
// 		t.Fatalf("Not found element after 1 insertion")
// 	}
// 	pt.Insert(inputs[1])
// 	if !pt.Search(inputs[1]) {
// 		t.Fatalf("Not found element after 1 insertion")
// 	}
//
// 	if pt.Search(inputs[2]) {
// 		t.Fatalf("Found element after 1 insertion")
// 	}
//
// 	pt.Insert(inputs[2])
// 	fmt.Printf("Header %v\nFirst child %v\nLeft %v\nRight %v\n", pt.header, pt.header.left, pt.header.left.left, pt.header.left.right)
// 	if pt.header.left.key != 2 {
// 		t.Fatalf("Not found element 2 where expected, found %v", pt.header.left.key)
// 	}
// 	if pt.header.left.right.key != 1 {
// 		t.Fatalf("Not found element 1 where expected, found %v", pt.header.left.right.key)
// 	}
// 	fmt.Printf("\n\n\n\n")
// }

func TestSearchHandbookStepByStep(t *testing.T) {
	inputs := []uint64{5, 0, 2, 8, 4, 10}
	// searches := []string{"marx", "ordo", "mass", "hello", "malleus", "me"}
	// outputs := []bool{false, false, true, false, true, true}

	// var pt PatriciaTrie
	pt := NewPatriciaTrie()

	// INSERT 5
	pt.Insert(inputs[0])
	if pt.header.key != inputs[0] {
		t.Fatalf("Expected %v in header, found %v", inputs[0], pt.header.key)
	}

	// INSERT 0
	pt.Insert(inputs[1])
	if pt.header.key != inputs[0] {
		t.Fatalf("Expected %v in header, found %v", inputs[0], pt.header.key)
	}
	if pt.header.left.key != inputs[1] {
		t.Fatalf("Expected %v in header.left, found %v", inputs[1], pt.header.left.key)
	}
	if pt.header.left.left.key != inputs[1] {
		t.Fatalf("Expected %v in header.left.left, found %v", inputs[1], pt.header.left.left.key)
	}
	if pt.header.left.right.key != inputs[0] {
		t.Fatalf("Expected %v in header.left.right.key, found %v", inputs[0], pt.header.left.right.key)
	}
	if pt.header.bit_index != 0 {
		t.Fatalf("Expected %v bit_index in header, found %v", 0, pt.header.bit_index)
	}
	if pt.header.left.bit_index != 61 {
		t.Fatalf("Expected %v bit_index in header.left, found %v", 61, pt.header.left.bit_index)
	}

	// INSERT 2
	pt.Insert(inputs[2])

	pt.Print()

	if pt.header.key != inputs[0] {
		t.Fatalf("Expected %v in header, found %v", inputs[0], pt.header.key)
	}
	if pt.header.left.key != inputs[1] {
		t.Fatalf("Expected %v in header.left, found %v", inputs[1], pt.header.left.key)
	}
	if pt.header.left.left.key != inputs[2] {
		t.Fatalf("Expected %v in header.left.left, found %v", inputs[2], pt.header.left.left.key)
	}
	if pt.header.left.right.key != inputs[0] {
		t.Fatalf("Expected %v in header.left.right.key, found %v", inputs[0], pt.header.left.right.key)
	}
	if pt.header.bit_index != 0 {
		t.Fatalf("Expected %v bit_index in header, found %v", 0, pt.header.bit_index)
	}
	if pt.header.left.bit_index != 61 {
		t.Fatalf("Expected %v bit_index in header.left, found %v", 61, pt.header.left.bit_index)
	}
	if pt.header.left.left.bit_index != 62 {
		t.Fatalf("Expected %v bit_index in header.left.left, found %v", 62, pt.header.left.left.bit_index)
	}
	if pt.header.left.left.right.key != inputs[2] {
		t.Fatalf("Expected %v in header.left.left.right, found %v", inputs[2], pt.header.left.left.right.key)
	}
	if pt.header.left.left.left.key != inputs[1] {
		t.Fatalf("Expected %v in header.left.left.left, found %v", inputs[1], pt.header.left.left.left.key)
	}

	// INSERT 8
	pt.Insert(inputs[3])
	if pt.header.key != 5 {
		t.Fatalf("Expected %v in header, found %v", 5, pt.header.key)
	}
	if pt.header.left.key != 8 {
		t.Fatalf("Expected %v in header.left, found %v", 8, pt.header.left.key)
	}
	if pt.header.left.left.key != 0 {
		t.Fatalf("Expected %v in header.left.left, found %v", 0, pt.header.left.left.key)
	}
	if pt.header.left.right.key != 8 {
		t.Fatalf("Expected %v in header.left.right.key, found %v", 8, pt.header.left.right.key)
	}
	if pt.header.left.left.right.key != 5 {
		t.Fatalf("Expected %v in header.left.left.right, found %v", 5, pt.header.left.left.right.key)
	}
	if pt.header.left.left.left.key != 2 {
		t.Fatalf("Expected %v in header.left.left.left, found %v", 2, pt.header.left.left.left.key)
	}
	if pt.header.left.left.left.left.key != 0 {
		t.Fatalf("Expected %v in header.left.left.left.left, found %v", 0, pt.header.left.left.left.left.key)
	}
	if pt.header.left.left.left.right.key != 2 {
		t.Fatalf("Expected %v in header.left.left.left.right, found %v", 2, pt.header.left.left.left.right.key)
	}

	if pt.header.bit_index != 0 {
		t.Fatalf("Expected %v bit_index in header, found %v", 0, pt.header.bit_index)
	}
	if pt.header.left.bit_index != 60 {
		t.Fatalf("Expected %v bit_index in header.left, found %v", 61, pt.header.left.bit_index)
	}
	if pt.header.left.left.bit_index != 61 {
		t.Fatalf("Expected %v bit_index in header.left, found %v", 61, pt.header.left.left.bit_index)
	}
	if pt.header.left.left.left.bit_index != 62 {
		t.Fatalf("Expected %v bit_index in header.left, found %v", 61, pt.header.left.left.left.bit_index)
	}

	fmt.Printf("Patricia Trie:\n%v\n", pt)
	pt.Print()
	pt.Insert(inputs[4])

	fmt.Printf("\n\n\n\n")

}

// func TestSearchHandbook(t *testing.T) {
// 	inputs := []uint64{5, 0, 2, 8, 4, 10}
// 	// searches := []string{"marx", "ordo", "mass", "hello", "malleus", "me"}
// 	// outputs := []bool{false, false, true, false, true, true}
//
// 	// var pt PatriciaTrie
// 	pt := NewPatriciaTrie()
// 	for _, num := range inputs {
// 		pt.Insert(num)
// 	}
//
// 	fmt.Printf("Patricia Trie:\n%v\n", pt)
// 	pt.Print()
//
// 	for _, num := range inputs {
// 		if !pt.Search(num) {
// 			t.Fatalf("Num not found %v and has been inserted", num)
// 		}
// 	}
// 	fmt.Printf("\n\n\n\n")
// }

//
// func TestPrint(t *testing.T) {
// 	inputs := []uint64{0, 1, 2, 3, 4}
//
// 	pt := NewPatriciaTrie()
//
// 	pt.Insert(inputs[0])
// 	pt.Insert(inputs[1])
// 	pt.Insert(inputs[2])
// 	pt.Insert(inputs[3])
// 	// pt.Insert(inputs[4])
//
// 	fmt.Printf("Patricia Trie:\n%v\n", pt)
// 	pt.Print()
//
// 	fmt.Printf("Header %v\n", pt.header)
// 	fmt.Printf("Node: %v\n", pt.header.left)
// }
