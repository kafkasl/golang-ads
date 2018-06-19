package patricia_trie

import (
	"fmt"
)

// Illustrative example described in Patricia Trie section of Handbook of
// Data Structures and Applications

func ExampleHandbook() {
	inputs := []uint64{5, 0, 2, 8, 4, 10}

	pt := NewPatriciaTrie()
	for _, i := range inputs[:] {
		pt.Insert(i)
	}

	fmt.Printf("%v", pt)

	// Output:Patricia Trie:
	// Header
	// └──  101[0] -> (1000, _)
	//     └── L 1000[61] -> (1010, _)
	//         ├── L 0[62] -> (100, _)
	//         │   ├── L 10[63] -> (10, _)
	//         │   └── R 100[64] -> (101, _)
	//         └── R 1010[63] -> (1010, _)

}
