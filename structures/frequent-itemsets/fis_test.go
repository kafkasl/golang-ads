package frequent_itemsets

import (
	"fmt"
	"testing"

	"github.com/OneOfOne/go-utils/memory"
)

func TestPatriciaBuilding(t *testing.T) {
	trie := NewTrie()
	output := "Trie:\n" +
		" └─ B[6]\n" +
		"    ├─ ADF[4]\n" +
		"    │  ├─ GH[1]\n" +
		"    │  └─ L[2]\n" +
		"    │     ├─ G[1]\n" +
		"    │     └─ H[1]\n" +
		"    └─ L[2]\n" +
		"       └─ GH[1]\n"

	var dataset *[]*[]uint = &[]*[]uint{
		&[]uint{1, 2, 4, 5, 6, 7, 8, 9},
		&[]uint{2, 3, 5, 12},
		&[]uint{1, 2, 4, 6, 8, 12},
		&[]uint{1, 2, 3, 4, 6, 7, 12},
		&[]uint{2, 7, 8, 12},
		&[]uint{1, 2, 4, 6, 9}}

	il := NewItemList(dataset, 3)
	fmt.Printf("Size of Itemlist: %v\n", memory.Sizeof(il))

	for _, tx := range il.txs {
		strTx := ""
		// orderedTx := orderTx(*tx, il)
		for _, e := range *tx {
			strTx += ToCharStr(e)
		}
		trie.Insert(strTx)
	}

	fmt.Printf("%v", trie)
	fmt.Printf("Size of Trie: %v\n", memory.Sizeof(trie))

	fisPTrie := NewFISPatriciaTrie(trie)

	str := fisPTrie.String()
	if str != output {
		t.Fatalf("Expected Patricia Trie: \n%s\nFound: \n%s\n", output, str)
	}

	fmt.Printf("IL :\n%v\nTrie: \n%v\nPatricia \n%v\n", il, trie, fisPTrie)
	fmt.Printf("Size of Frequent Itemset Patricia Trie: %v\n", memory.Sizeof(fisPTrie))

}
