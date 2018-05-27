package frequent_itemsets

import (
	"fmt"
	"testing"
)

func TestPatriciaBuilding(t *testing.T) {
	trie := NewTrie()

	var dataset *[]*[]uint = &[]*[]uint{
		&[]uint{1, 2, 4, 5, 6, 7, 8, 9},
		&[]uint{2, 3, 5, 12},
		&[]uint{1, 2, 4, 6, 8, 12},
		&[]uint{1, 2, 3, 4, 6, 7, 12},
		&[]uint{2, 7, 8, 12},
		&[]uint{1, 2, 4, 6, 9}}

	il := NewItemList(dataset, 3)

	for _, tx := range *dataset {
		strTx := ""
		orderedTx := orderTx(*tx, il)
		for _, e := range orderedTx {
			strTx += ToCharStr(e)
		}
		trie.Insert(strTx)
	}

	fmt.Printf("%v", trie)

	fisPTrie := NewFISPatriciaTrie(trie)

	fmt.Printf("Patricia version\n%s\n", fisPTrie)
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
