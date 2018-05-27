package frequent_itemsets

import (
	"bytes"
	"fmt"
)

// func (tn TrieNode) Words(currentWord string) []string {
// 	words := []string{}
// 	if tn.endOfWord {
// 		return append(words, currentWord)
// 	}
// 	if len(tn.children) > 0 {
// 		keys := make([]rune, 0)
// 		for k, _ := range tn.children {
// 			keys = append(keys, k)
// 		}
// 		sort.Sort(RuneSlice(keys))
//
// 		for _, key := range keys {
// 			words = append(words, tn.children[key].Words(currentWord+string(key))...)
// 		}
// 	}
// 	return words
// }
//
// func (tn TrieNode) toString(prefix string) string {
// 	var text string = ""
// 	var nprefix string
// 	keys := make([]rune, 0)
// 	for k, _ := range tn.children {
// 		keys = append(keys, k)
// 	}
// 	sort.Sort(RuneSlice(keys))
//
// 	idx := 0
// 	for _, key := range keys {
//
// 		if idx < len(tn.children)-1 {
// 			text += prefix + " ├─ " + string(key) + "[" + fmt.Sprintf("%v", tn.children[key].count) + "]" + "\n"
// 			nprefix = prefix + " │ "
// 		} else {
// 			text += prefix + " └─ " + string(key) + "[" + fmt.Sprintf("%v", tn.children[key].count) + "]" + "\n"
// 			nprefix = prefix + "   "
// 		}
// 		text += tn.children[key].toString(nprefix)
// 		idx++
// 	}
// 	return text
// }
// func (tn TrieNode) String() string {
// 	text := ""
// 	for k, _ := range tn.children {
// 		text += string(k)
// 	}
// 	return fmt.Sprintf("Children: %v. EOW: %v", text, tn.endOfWord)
// }
// func Walk(t *TrieNode, ch chan int) {
// 	if t.Left != nil {
// 		Walk(t.Left, ch)
// 	}
// 	ch <- t.Value
// 	if t.Right != nil {
// 		Walk(t.Right, ch)
// 	}
// }

func (tn *TrieNode) Compress(myKey rune, accKey []rune, accCount uint, ancestor **TrieNode) {
	fmt.Printf("Call: \n\tmyKey: %v[%v]\n\taccKey: %v \n\taccCount %v\n\tAncestor %v\n",
		string(myKey), tn.count, accKey, accCount, ancestor)
	var buffer bytes.Buffer

	// if tn.count == accCount && len(tn.children) == 1 {

	// 	for k, v := range nn.children {
	// 		// fmt.Printf("1. Compressing: k: %v v: %v\n", string(k), v)
	// 		v.Compress(k, append(accKey, myKey), tn.count, ancestor)
	// 	}
	// } else if tn.count != accCount {
	// 	if len(accKey) > 1 {
	// 		//compress

	// 	}
	// }
	if tn.count == accCount && len(tn.children) == 1 {

		for k, v := range tn.children {
			// fmt.Printf("1. Compressing: k: %v v: %v\n", string(k), v)
			v.Compress(k, append(accKey, myKey), tn.count, ancestor)
		}
	} else if tn.count != accCount {
		delete((*ancestor).children, accKey[0])
		for _, r := range accKey {
			buffer.WriteRune(r)

		}
		children := make(map[rune]*TrieNode)
		children[myKey] = tn
		nn := &TrieNode{children, accCount}
		newKey, _, _ := buffer.ReadRune()
		(*ancestor).children[newKey] = nn
		fmt.Printf("newKey %v %v", newKey, nn)

		for k, v := range nn.children {
			// fmt.Printf("1. Compressing: k: %v v: %v\n", string(k), v)
			v.Compress(k, []rune{k}, tn.count, &nn)
		}
	} else if len(tn.children) > 1 || tn.children == nil {
		// Have the same key as the last one
		delete((*ancestor).children, accKey[0])
		for _, r := range accKey {
			buffer.WriteRune(r)

		}
		buffer.WriteRune(myKey)
		nn := &TrieNode{tn.children, accCount}
		newKey, _, _ := buffer.ReadRune()
		(*ancestor).children[newKey] = nn

		fmt.Printf("newKey %v %v", newKey, nn)

		for k, v := range nn.children {
			fmt.Printf("2. Compressing: k: %v v: %v\n", string(k), v)
			v.Compress(k, []rune{k}, v.count, &tn)
		}
	}
	// else {

	// 	// Node has same key and a single child
	// 	for k, v := range tn.children {
	// 		// fmt.Printf("3. Compressing: k: %v v: %v\n", string(k), v)
	// 		v.Compress(k, append(accKey, k), v.count, ancestor)
	// 	}
	// }

}

type FISPatriciaTrie struct {
	root *TrieNode
}

// func (tn *TrieNode) dCompress(mykey, accKey rune, accCount uint) (key rune, count uint, nextNode *TrieNode) {
// 	if tn.count != accCount {
// 		return accKey, accCount, &tn
// 	}
// 	if len(tn.Children) > 1 {
// 		return accKey + myKey, accCount, tn
// 	}
// 	tn.children[0].Compress(accKey, accCount)
// }

func NewFISPatriciaTrie(trie Trie) Trie {

	for k, v := range trie.root.children {
		v.Compress(k, []rune{k}, v.count, &trie.root)
	}
	return trie
}

//
// func (t Trie) String() string {
// 	return "Trie:\n" + t.root.toString("")
// }
//
// func (t Trie) Words() []string {
// 	return t.root.Words("")
// }
//
// func (t Trie) Insert(key string) {
// 	currentNode := t.root
// 	for _, letter := range key {
// 		if val, ok := currentNode.children[letter]; ok {
// 			val.count++
// 			currentNode = val
// 		} else {
// 			nn := &TrieNode{make(map[rune]*TrieNode), 1, false}
// 			currentNode.children[letter] = nn
// 			currentNode = nn
// 		}
// 	}
// 	currentNode.endOfWord = true
// }
//
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
