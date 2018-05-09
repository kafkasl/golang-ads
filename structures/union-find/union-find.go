package structures

import "fmt"

type UnionFind interface {
	Find() UnionFind
	Union(other UnionFind) UnionFind
}

type UnionFindSet struct {
	parent *UnionFindSet
	rank   int
	value  int
}

func NewUnionFindSet(value int) *UnionFindSet {
	uf := UnionFindSet{nil, 0, value}
	uf.parent = &uf
	return &uf
}

func (uf *UnionFindSet) Find() *UnionFindSet {
	if uf.parent == uf {
		return uf
	}
	return uf.parent.Find()

}

func (uf *UnionFindSet) Union(other *UnionFindSet) *UnionFindSet {
	ufParent := uf.Find()
	otherParent := other.Find()
	if ufParent == otherParent {
		return ufParent
	}
	if ufParent.rank > otherParent.rank {
		otherParent.parent = ufParent
		return ufParent
	} else {
		ufParent.parent = otherParent
		if ufParent.rank == otherParent.rank {
			otherParent.rank++
		}
		return otherParent
	}

}

func (uf *UnionFindSet) String() string {
	if uf.parent == uf {
		return fmt.Sprintf("[%v]: %v", uf.rank, uf.value)
	}
	str := fmt.Sprintf("[%v]: %v -> %v", uf.rank, uf.value, uf.parent)
	return str
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
