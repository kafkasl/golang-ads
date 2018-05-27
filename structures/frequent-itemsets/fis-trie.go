package frequent_itemsets

import (
	"fmt"
	"sort"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type TrieNode struct {
	children map[rune]*TrieNode
	count    uint
}

func (tn TrieNode) toString(prefix string) string {
	var text string = ""
	var nprefix string
	keys := make([]rune, 0)
	for k, _ := range tn.children {
		keys = append(keys, k)
	}
	sort.Sort(RuneSlice(keys))

	idx := 0
	for _, key := range keys {

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
func (tn TrieNode) String() string {
	text := ""
	for k, _ := range tn.children {
		text += string(k) + " "
	}
	return fmt.Sprintf("Children: %v.", text)
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	tn := TrieNode{make(map[rune]*TrieNode), 0}
	return Trie{&tn}
}

func (t Trie) String() string {
	return "Trie:\n" + t.root.toString("")
}

func (t Trie) Insert(key string) {
	currentNode := t.root
	for _, letter := range key {
		if val, ok := currentNode.children[letter]; ok {
			val.count++
			currentNode = val
		} else {
			nn := &TrieNode{make(map[rune]*TrieNode), 1}
			currentNode.children[letter] = nn
			currentNode = nn
		}
	}
}

// func (t Trie) Search(key string) bool {
// 	currentNode := t.root
// 	for _, letter := range key {
// 		if val, ok := currentNode.children[letter]; ok {
// 			currentNode = val
// 		} else {
// 			return false
// 		}
// 	}
// 	return currentNode.endOfWord
// }
