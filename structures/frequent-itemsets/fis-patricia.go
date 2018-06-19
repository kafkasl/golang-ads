package frequent_itemsets

func (tn *TrieNode) Compress(myKey string, accKey string, accCount uint, ancestor **TrieNode) {

	if tn.count == accCount && len(tn.children) == 1 {

		for k, v := range tn.children {
			v.Compress(k, accKey+k, tn.count, ancestor)
		}
	} else if tn.count != accCount {
		delete((*ancestor).children, string(accKey[0]))

		children := make(map[string]*TrieNode)
		children[myKey] = tn
		nn := &TrieNode{children, accCount}
		(*ancestor).children[accKey[0:len(accKey)-1]] = nn

		for k, v := range nn.children {
			v.Compress(k, k, tn.count, &nn)
		}
	} else if len(tn.children) > 1 || len(tn.children) == 0 {
		// Have the same key as the last one
		delete((*ancestor).children, string(accKey[0]))

		nn := &TrieNode{tn.children, accCount}
		(*ancestor).children[accKey] = nn

		for k, v := range nn.children {
			v.Compress(k, k, v.count, &tn)
		}
	}

}

type FISPatriciaTrie struct {
	root *TrieNode
}

func NewFISPatriciaTrie(trie Trie) *FISPatriciaTrie {

	for k, v := range trie.root.children {
		v.Compress(k, k, v.count, &trie.root)
	}
	return &FISPatriciaTrie{trie.root}
}

func (t FISPatriciaTrie) String() string {
	return "Patricia trie:\n" + t.root.toString("")
}
