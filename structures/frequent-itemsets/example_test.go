package structures

import (
	"fmt"
)

func ExamplePaper() {

	var dataset *[]*[]uint = &[]*[]uint{
		&[]uint{1, 2, 4, 5, 6, 7, 8, 9},
		&[]uint{2, 3, 5, 12},
		&[]uint{1, 2, 4, 6, 8, 12},
		&[]uint{1, 2, 3, 4, 6, 7, 12},
		&[]uint{2, 7, 8, 12},
		&[]uint{1, 2, 4, 6, 9}}

	il := NewItemList(dataset, 3)

	fmt.Printf("%v", il)

	// Output:B [6]: [A B D E F G H I] + [B C E L] + [A B D F H L] + [A B C D F G L] + [B G H L] + [A B D F I]
	// A [4]: [A B D E F G H I] + [A B D F H L] + [A B C D F G L] + [A B D F I]
	// D [4]: [A B D E F G H I] + [A B D F H L] + [A B C D F G L] + [A B D F I]
	// F [4]: [A B D E F G H I] + [A B D F H L] + [A B C D F G L] + [A B D F I]
	// L [4]: [B C E L] + [A B D F H L] + [A B C D F G L] + [B G H L]
	// G [3]: [A B D E F G H I] + [A B C D F G L] + [B G H L]
	// H [3]: [A B D E F G H I] + [A B D F H L] + [B G H L]
}

func ExampleBuildTrie() {
	t := NewTrie()
}
