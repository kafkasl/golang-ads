package frequent_itemsets

import (
	"fmt"
	"sort"
)

type RuneSlice []string

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type TrieNode struct {
	children map[string]*TrieNode
	count    uint
}

func (tn TrieNode) toString(prefix string) string {
	var text string = ""
	var nprefix string
	keys := make([]string, 0)
	for k, _ := range tn.children {
		keys = append(keys, k)
	}
	sort.Sort(RuneSlice(keys))

	idx := 0
	for _, key := range keys[:] {

		if idx < len(tn.children)-1 {
			text += prefix + " ├─ " + string(key) + "[" + fmt.Sprintf("%v", tn.children[key].count) + "]" + "\n"
			nprefix = prefix + " │ "
		} else {
			text += prefix + " └─ " + string(key) + "[" + fmt.Sprintf("%v", tn.children[key].count) + "]" + "\n"
			nprefix = prefix + "   "
		}
		text += tn.children[key].toString(nprefix)
		idx++
	}
	return text
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	tn := TrieNode{make(map[string]*TrieNode), 0}
	return Trie{&tn}
}

func (t Trie) String() string {
	return "Trie:\n" + t.root.toString("")
}

func (t Trie) Insert(key string) {
	currentNode := t.root
	for _, letter := range key[:] {
		if val, ok := currentNode.children[string(letter)]; ok {
			val.count++
			currentNode = val
		} else {
			nn := &TrieNode{make(map[string]*TrieNode), 1}
			currentNode.children[string(letter)] = nn
			currentNode = nn
		}
	}
}
