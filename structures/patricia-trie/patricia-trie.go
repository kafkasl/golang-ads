package structures

import (
	"fmt"
	"math/bits"

	"github.com/disiqueira/gotree"
)

// type RuneSlice []rune
//
// func (p RuneSlice) Len() int           { return len(p) }
// func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
// func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type PatriciaTrieNode struct {
	key       uint64
	bit_index uint
	left      *PatriciaTrieNode
	right     *PatriciaTrieNode
}

// func (ptn PatriciaTrieNode) Search(key uint64) bool {
// 	if key&(1<<ptn.bit_index) == 0 {
// 	}
// 	return ptn.search(key, -1)
// }

func (ptn *PatriciaTrieNode) search(key uint64, prev_bi uint) bool {
	if prev_bi >= ptn.bit_index {
		return key == ptn.key
	}
	if key&(1<<ptn.bit_index) == 0 {
		return ptn.left.search(key, ptn.bit_index)
	} else {
		return ptn.right.search(key, ptn.bit_index)
	}
}

func (ptn *PatriciaTrieNode) find(key uint64, prev_bi int) (*PatriciaTrieNode, *PatriciaTrieNode) {
	fmt.Printf("prev_bi %v, bit_index %v [%v]\n", prev_bi, ptn.bit_index, uint(prev_bi))
	if prev_bi == -1 && ptn == ptn.left {
		return nil, ptn
	}
	if prev_bi > -1 && uint(prev_bi) >= ptn.bit_index {
		fmt.Printf("Returning node: %p\n", ptn)
		return ptn, nil
	}
	var parent, node *PatriciaTrieNode
	if key&(1<<ptn.bit_index) == 0 {
		node, parent = ptn.left.find(key, int(ptn.bit_index))
	} else {
		node, parent = ptn.right.find(key, int(ptn.bit_index))
	}
	if parent == nil {
		parent = ptn
	}
	fmt.Printf("Returning parent: %p\n", ptn)
	return node, parent

}

func (ptn *PatriciaTrieNode) depth(accDepth int32) int {
	var rd, ld int
	if ptn.right != nil {
		rd = ptn.right.depth(accDepth + 1)
	}
	if ptn.left != nil {
		ld = ptn.left.depth(accDepth + 1)
	}
	if rd > ld {
		return rd
	}
	return ld
}

// func (ptn PatriciaTrieNode) toString(acc, prefix string) string {
// 	if ptn.right == nil {
// 		acc += fmt.Sprintf("-%v\n", ptn.key)
// 	} else {
// 		acc += fmt.Sprintf("-%v", ptn.key)
// 		prefix += " |"
// 	}
// 	if ptn.left != nil {
// 		acc += fmt.Sprintf("%v |", prefix)
// 		ptn.left.toString(acc, prefix)
// 	}
// }

// func Print(ptn *PatriciaTrieNode, prefix string) {
// 	fmt.Printf("- %v", ptn.key)
// 	if ptn.right != nil {
// 		Print(ptn.right, fmt.Sprintf("%v |", prefix))
// 	} else {
// 		fmt.Println()
// 		fmt.Printf("%v |\n", prefix)
// 	}
// 	fmt.Printf("%v %v")
//
// }
type PatriciaTrie struct {
	header *PatriciaTrieNode
}

func NewPatriciaTrie() PatriciaTrie {
	return PatriciaTrie{nil}
}

func (pt *PatriciaTrie) Search(key uint64) bool {
	if pt.header == nil {
		return false
	}
	if pt.header == pt.header.left {
		return key == pt.header.key
	}
	return pt.header.left.search(key, 0)
}

func (pt *PatriciaTrie) Insert(key uint64) {
	if pt.header == nil {
		node := &PatriciaTrieNode{key, 0, nil, nil}
		node.left = node
		pt.header = node
		fmt.Printf("pt, %p\n", &pt)
	} else {

		if !pt.Search(key) {
			endNode, parentNode := pt.header.left.find(key, -1)
			fmt.Printf("Key to insert: %v, found %v\n", key, endNode.key)
			// fmt.Printf("Parent Node %p, endNode %p\n", parentNode, endNode)
			reachedKey := (*endNode).key
			// fmt.Printf("Reached key: %v\n", reachedKey)
			var rBitPos uint = 0
			for ; rBitPos < uint(bits.Len64(key)) && (key&(1<<rBitPos)) == (reachedKey&(1<<rBitPos)); rBitPos++ {

			}
			fmt.Printf("New rBitPos: %v\n", rBitPos)
			newNode := PatriciaTrieNode{key, rBitPos, nil, nil}
			if key&(1<<rBitPos) != 0 {
				newNode.right = &newNode
				newNode.left = endNode
				// Need to make p point to newNode and newNode to endNode
			} else {
				newNode.left = &newNode
				newNode.right = endNode
			}

			if parentNode.right == endNode {
				parentNode.right = &newNode
			} else {
				parentNode.left = &newNode
			}
		}
	}
	fmt.Printf("pt, %p\n", &pt)

}

func (pt *PatriciaTrie) Depth() int {
	if pt.header == nil {
		return 0
	}
	return pt.header.depth(0)
}

func (ptn *PatriciaTrieNode) populateTree(parent *gotree.Tree) {
	node := (*parent).Add(fmt.Sprintf("%v", ptn.key))
	if ptn.left != nil && ptn.bit_index < ptn.left.bit_index {
		ptn.left.populateTree(&node)
	}
	if ptn.right != nil && ptn.bit_index != ptn.right.bit_index {
		ptn.right.populateTree(&node)
	}
}

func print(ptn *PatriciaTrieNode) {
	if ptn != nil {
		if ptn.left != nil && ptn.bit_index < ptn.left.bit_index {
			print(ptn.left)
		}
		fmt.Printf("%v", ptn.key)
		if ptn.right != nil && ptn.bit_index < ptn.right.bit_index {
			print(ptn.right)
		}
	}
}

func (pt *PatriciaTrie) Print() {
	fmt.Printf("Header %v", pt.header)
	print(pt.header)
	header := gotree.New(fmt.Sprintf("%v", pt.header.key))

	pt.header.populateTree(&header)

	fmt.Println(header.Print())

}

// func (tn PatriciaTrieNode) String() string {
// 	text := ""
// 	for k, _ := range tn.children {
// 		text += string(k)
// 	}
// 	return fmt.Sprintf("Children: %v. EOW: %v", text, tn.endOfWord)
// }
//
// type Trie struct {
// 	root *PatriciaTrieNode
// }
//
// func NewTrie() Trie {
// 	tn := PatriciaTrieNode{make(map[rune]*PatriciaTrieNode), false}
// 	return Trie{&tn}
// }
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
// 			currentNode = val
// 		} else {
// 			nn := &PatriciaTrieNode{make(map[rune]*PatriciaTrieNode), false}
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
