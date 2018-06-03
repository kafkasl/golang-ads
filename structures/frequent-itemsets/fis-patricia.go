package frequent_itemsets

func (tn *TrieNode) Compress(myKey string, accKey string, accCount uint, ancestor **TrieNode) {
	// fmt.Printf("Call: \n\tmyKey: %v[%v]\n\taccKey: %v \n\taccCount %v\n\tAncestor %v\n",
	// string(myKey), tn.count, accKey, accCount, ancestor)

	// fmt.Printf("childs %v", tn.children)
	if tn.count == accCount && len(tn.children) == 1 {

		// ((tn.count != accCount || len(tn.children) > 1 || tn.children == nil) && *ancestor == tn) {

		for k, v := range tn.children {
			// fmt.Printf("1. Compressing: k: %v v: %v\n", string(k), v)
			v.Compress(k, accKey+k, tn.count, ancestor)
		}
	} else if tn.count != accCount {
		delete((*ancestor).children, string(accKey[0]))

		children := make(map[string]*TrieNode)
		children[myKey] = tn
		nn := &TrieNode{children, accCount}
		(*ancestor).children[accKey[0:len(accKey)-1]] = nn
		// fmt.Printf("newKey %v %v", newKey, nn)

		for k, v := range nn.children {
			// fmt.Printf("1. Compressing: k: %v v: %v\n", string(k), v)
			v.Compress(k, k, tn.count, &nn)
		}
	} else if len(tn.children) > 1 || len(tn.children) == 0 {
		// Have the same key as the last one
		delete((*ancestor).children, string(accKey[0]))

		nn := &TrieNode{tn.children, accCount}
		(*ancestor).children[accKey] = nn

		// fmt.Printf("newKey %v %v", newKey, nn)

		// fmt.Printf("2. Compressing: k: %v v: %v\n", string(accKey), nn)
		for k, v := range nn.children {
			v.Compress(k, k, v.count, &tn)
		}
	}

}

type FISPatriciaTrie struct {
	root *TrieNode
}

func NewFISPatriciaTrie(trie Trie) Trie {

	for k, v := range trie.root.children {
		v.Compress(k, k, v.count, &trie.root)
	}
	return trie
}
