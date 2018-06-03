package frequent_itemsets

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

	// Output:B [6]: [B A D F G H] + [B L] + [B A D F L H] + [B A D F L G] + [B L G H] + [B A D F]
	//A [4]: [B A D F G H] + [B A D F L H] + [B A D F L G] + [B A D F]
	//D [4]: [B A D F G H] + [B A D F L H] + [B A D F L G] + [B A D F]
	//F [4]: [B A D F G H] + [B A D F L H] + [B A D F L G] + [B A D F]
	//L [4]: [B L] + [B A D F L H] + [B A D F L G] + [B L G H]
	//G [3]: [B A D F G H] + [B A D F L G] + [B L G H]
	//H [3]: [B A D F G H] + [B A D F L H] + [B L G H]

}

func ExampleBuildTrie() {
	t := NewTrie()

	var dataset *[]*[]uint = &[]*[]uint{
		&[]uint{1, 2, 4, 5, 6, 7, 8, 9},
		&[]uint{2, 3, 5, 12},
		&[]uint{1, 2, 4, 6, 8, 12},
		&[]uint{1, 2, 3, 4, 6, 7, 12},
		&[]uint{2, 7, 8, 12},
		&[]uint{1, 2, 4, 6, 9}}

	il := NewItemList(dataset, 3)

	for _, tx := range il.txs {
		strTx := ""
		for _, e := range *tx {
			strTx += ToCharStr(e)
		}
		t.Insert(strTx)
	}

	fmt.Printf("%v", t)

	// Output: Trie:
	//  └─ B[6]
	//     ├─ A[4]
	//     │  └─ D[4]
	//     │     └─ F[4]
	//     │        ├─ G[1]
	//     │        │  └─ H[1]
	//     │        └─ L[2]
	//     │           ├─ G[1]
	//     │           └─ H[1]
	//     └─ L[2]
	//        └─ G[1]
	//           └─ H[1]

}
