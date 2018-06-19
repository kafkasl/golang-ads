package skip_list

import "fmt"

func ExampleString() {
	d := NewDictionary(1.1) // with 1.1 height is never augmented

	var input = map[int]string{
		1: "aa",
		2: "ab",
		3: "ac",
		4: "ba",
		5: "bb",
		6: "bc",
		7: "ca",
		8: "cb",
		9: "cc",
	}

	for k, v := range input {
		d.Insert(k, v)
	}

	fmt.Printf("%v", d)
	// Output: Skip List (height: 1)
	// [ ] [-∞: start (1)]
	//  |
	//  V
	// [ ] [1: aa (1)]
	//  |
	//  V
	// [ ] [2: ab (1)]
	//  |
	//  V
	// [ ] [3: ac (1)]
	//  |
	//  V
	// [ ] [4: ba (1)]
	//  |
	//  V
	// [ ] [5: bb (1)]
	//  |
	//  V
	// [ ] [6: bc (1)]
	//  |
	//  V
	// [ ] [7: ca (1)]
	//  |
	//  V
	// [ ] [8: cb (1)]
	//  |
	//  V
	// [ ] [9: cc (1)]
	//  |
	//  V
	// [ ] [+∞: end (1)]

}
