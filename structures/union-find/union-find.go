package union_find

import (
	"bytes"
	"fmt"
)

type UnionFindSet struct {
	g []int
}

func NewUnionFindSet(n int) *UnionFindSet {
	g := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = i
	}
	uf := UnionFindSet{g}

	return &uf
}

func (uf *UnionFindSet) Find(x int) int {
	if uf.g[x] == x {
		return x
	}
	uf.g[x] = uf.Find(uf.g[x])
	return uf.g[x]

}

func (uf *UnionFindSet) Union(x, y int) int {
	ufParent := uf.Find(x)
	otherParent := uf.Find(y)
	uf.g[ufParent] = otherParent
	return ufParent
}

func (uf *UnionFindSet) String() string {
	var buffer bytes.Buffer
	for i := 0; i < len(uf.g); i++ {
		buffer.WriteString(fmt.Sprintf("%v ", uf.g[i]))
	}

	return buffer.String()
}

//
// func (tn UnionFindSetNode) Words(currentWord string) []string {
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
// func (tn UnionFindSetNode) toString(prefix string) string {
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
// 			text += prefix + " ├─ " + string(key) + "\n"
// 			nprefix = prefix + " │ "
// 		} else {
// 			text += prefix + " └─ " + string(key) + "\n"
// 			nprefix = prefix + "   "
// 		}
// 		text += tn.children[key].toString(nprefix)
// 		idx++
// 	}
// 	return text
// }
// func (tn UnionFindSetNode) String() string {
// 	text := ""
// 	for k, _ := range tn.children {
// 		text += string(k)
// 	}
// 	return fmt.Sprintf("Children: %v. EOW: %v", text, tn.endOfWord)
// }
//
// type UnionFindSet struct {
// 	root *UnionFindSetNode
// }
//
//
//
// func (t UnionFindSet) String() string {
// 	return "UnionFindSet:\n" + t.root.toString("")
// }
//
// func (t UnionFindSet) Words() []string {
// 	return t.root.Words("")
// }
//
// func (t UnionFindSet) Insert(key string) {
// 	currentNode := t.root
// 	for _, letter := range key {
// 		if val, ok := currentNode.children[letter]; ok {
// 			currentNode = val
// 		} else {
// 			nn := &UnionFindSetNode{make(map[rune]*UnionFindSetNode), false}
// 			currentNode.children[letter] = nn
// 			currentNode = nn
// 		}
// 	}
// 	currentNode.endOfWord = true
// }
//
// func (t UnionFindSet) Search(key string) bool {
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
