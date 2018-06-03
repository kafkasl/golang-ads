package frequent_itemsets

import (
	"fmt"
)

// type Pair struct {
// 	key   uint
// 	count uint
// }
//
// type PairSlice []Pair
//
// func (p PairSlice) Len() int { return len(p) }
// func (p PairSlice) Less(i, j int) bool {
// 	if p[i].count == p[j].count {
// 		return p[i].key > p[j].key
// 	}
// 	return p[i].count < p[j].count
// }
// func (p PairSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
//
// func orderTx(list []uint, itemlist *ItemList) (result []uint) {
// 	var ps PairSlice
// 	for _, key := range list {
// 		if item, ok := itemlist.v[key]; ok {
// 			ps = append(ps, Pair{key, item.count})
// 		}
// 	}
// 	sort.Sort(sort.Reverse(ps))
// 	for _, pair := range ps {
// 		result = append(result, pair.key)
// 	}
// 	return
// }

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
		// orderedTx := orderTx(*tx, il)
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
