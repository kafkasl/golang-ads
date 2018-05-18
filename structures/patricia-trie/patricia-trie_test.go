package structures

import (
	"fmt"
	"testing"
)

func PrintFatal(t *testing.T, expected fmt.Stringer, got fmt.Stringer) {
	t.Fatalf("Expected: \n%v, but got: \n%v\n", expected, got)
}

func TestEmpty(t *testing.T) {
	pt := NewPatriciaTrie()
	var i uint64
	for i = 0; i < 10; i++ {
		if pt.Search(i) {
			t.Fatalf("Found %v and has not been inserted", i)
		}
	}

	if pt.header != nil {
		t.Fatalf("Header should be nil")
	}
}

func TestHeaderOnly(t *testing.T) {
	pt := NewPatriciaTrie()
	var i uint64 = 0
	pt.Insert(i)
	if !pt.Search(i) {
		t.Fatalf("Not found %v and has been inserted", i)
	}
	for i = 1; i < 10; i++ {
		if pt.Search(i) {
			t.Fatalf("Found %v and has not been inserted", i)
		}
	}

	if pt.header.key != 0 {
		t.Fatalf("Header key should be 0")
	}
	if pt.header.bit_index != 0 {
		t.Fatalf("Header bit index should be 0")
	}
	if pt.header.left != pt.header {
		t.Fatalf("Header left should point to itself but it is %v instead of %v", pt.header.left, pt.header)
	}
	if pt.header.right != nil {
		t.Fatalf("Header right should be nil but it is %v", pt.header.right)
	}
}

func TestRepeatedInsert(t *testing.T) {
	pt := NewPatriciaTrie()
	var i uint64
	for i = 1; i < 10; i++ {
		pt.Insert(0)
	}

	if !pt.Search(0) {
		t.Fatalf("Found %v and has not been inserted", i)
	}

	if pt.header.key != 0 {
		t.Fatalf("Header key should be 0")
	}
	if pt.header.bit_index != 0 {
		t.Fatalf("Header bit index should be 0")
	}
	if pt.header.left != pt.header {
		t.Fatalf("Header left should point to itself but it is %v instead of %v", pt.header.left, pt.header)
	}
	if pt.header.right != nil {
		t.Fatalf("Header right should be nil but it is %v", pt.header.right)
	}
}

func TestSearchHandbookStepByStep(t *testing.T) {
	inputs := []uint64{5, 0, 2, 8, 4, 10}

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
	pt.Print()
	pt.Insert(inputs[5])
	pt.Print()
	fmt.Printf("\n\n\n\n")

}

func TestDepth(t *testing.T) {
	inputs := []uint64{1, 3, 7, 15}
	pt := NewPatriciaTrie()
	if pt.Depth() != 0 {
		t.Fatalf("Depth is wrong, expected value: %v, found: %v", 0, pt.Depth())
	}

	for eDepth, i := range inputs {
		pt.Insert(i)
		pt.Print()
		d := pt.Depth()
		if d != eDepth+1 {
			t.Fatalf("Depth is wrong, expected value: %v, found: %v", eDepth+1, d)
		}
	}
}

func TestDepthReverse(t *testing.T) {
	inputs := []uint64{15, 7, 3, 1}
	pt := NewPatriciaTrie()
	if pt.Depth() != 0 {
		t.Fatalf("Depth is wrong, expected value: %v, found: %v", 0, pt.Depth())
	}

	for eDepth, i := range inputs {
		pt.Insert(i)
		pt.Print()
		d := pt.Depth()
		if d != eDepth+1 {
			t.Fatalf("Depth is wrong, expected value: %v, found: %v", eDepth+1, d)
		}
	}
}

func TestRightInsert(t *testing.T) {
	inputs := []uint64{1, 3, 7, 15}
	pt := NewPatriciaTrie()
	if pt.Depth() != 0 {
		t.Fatalf("Depth is wrong, expected value: %v, found: %v", 0, pt.Depth())
	}
	// pt.Insert(0)
	// t.Logf("Depth: %v", pt.Depth())

	for _, num := range inputs {
		pt.Insert(num)
		pt.Print()
		if !pt.Search(num) {
			t.Fatalf("Key %v should be present", num)
		}
	}
}

func TestSearchHandbook(t *testing.T) {
	inputs := []uint64{5, 0, 2, 8, 4, 10}
	no_inputs := []uint64{1, 3, 6, 7, 9}
	// searches := []string{"marx", "ordo", "mass", "hello", "malleus", "me"}
	// outputs := []bool{false, false, true, false, true, true}

	// var pt PatriciaTrie
	pt := NewPatriciaTrie()
	for _, num := range inputs {
		pt.Insert(num)
	}

	fmt.Printf("Patricia Trie:\n%v\n", pt)
	pt.Print()

	for _, num := range no_inputs {
		if pt.Search(num) {
			t.Fatalf("Num %v found and has not been inserted", num)
		}
	}
	fmt.Printf("\n\n\n\n")
}
