package structures

import (
	"fmt"

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

func (ptn *PatriciaTrieNode) String() string {
	l, r := "_", "_"
	if ptn.left != nil {
		l = fmt.Sprintf("%b", ptn.left.key)
	}
	if ptn.right != nil {
		r = fmt.Sprintf("%b", ptn.right.key)
	}
	return fmt.Sprintf("[%b, %v] -> (%v, %v)", ptn.key, ptn.bit_index, l, r)
}

// func (ptn PatriciaTrieNode) Search(key uint64) bool {
// 	if key&(1<<ptn.bit_index) == 0 {
// 	}
// 	return ptn.search(key, -1)
// }

func (ptn *PatriciaTrieNode) search(key uint64, prev_bi uint) bool {
	// fmt.Printf("Search: \nCurrent node: %v\nprev_bi %v\nbit_index %v [%v]\n", ptn, prev_bi, ptn.bit_index, uint(prev_bi))

	if prev_bi >= ptn.bit_index {
		return key == ptn.key
	}
	if key&((1<<63)>>ptn.bit_index) == 0 {
		return ptn.left.search(key, ptn.bit_index)
	} else {
		return ptn.right.search(key, ptn.bit_index)
	}
}

func (ptn *PatriciaTrieNode) find(key uint64, prev_bi uint) *PatriciaTrieNode {
	// fmt.Printf("prev_bi %v, bit_index %v [%v]\n", prev_bi, ptn.bit_index, uint(prev_bi))
	if prev_bi >= ptn.bit_index {
		return ptn
	}
	if key&((1<<63)>>ptn.bit_index) == 0 {
		return ptn.left.find(key, ptn.bit_index)
	} else {
		return ptn.right.find(key, ptn.bit_index)
	}
}

func (ptn *PatriciaTrieNode) insertNode(node, endNode, parentNode *PatriciaTrieNode) bool {
	fmt.Printf("\n\n\nSTART\nParent Node: %v\nEnd Node: %v\nNode: %v\nCurrent Node: %v", parentNode, endNode, node, ptn)
	// if ptn.bit_index <= parentNode.bit_index {
	// 	return false
	// }
	var nextNode *PatriciaTrieNode
	if node.key&((1<<63)>>ptn.bit_index) == 0 {
		nextNode = ptn.left
	} else {
		nextNode = ptn.right
	}
	// if nextNode.bit_index < node.bit_index && ptn != endNode {
	if nextNode.bit_index < node.bit_index && parentNode.bit_index < ptn.bit_index {
		return nextNode.insertNode(node, endNode, ptn)
	} else {
		// if node.bit_index > ptn.bit_index {
		// 	parentNode = ptn
		// 	ptn = nextNode
		// }
		fmt.Printf("\n\n\nParent Node: %v\nNext Node: %v\nNode: %v\nCurrent node: %v\n\n\n", parentNode, nextNode, node, ptn)
		if parentNode.right == ptn {
			parentNode.right = node
		} else if parentNode.left == ptn {
			parentNode.left = node
		} else {
			fmt.Println("ERROR: I am no son of my parent")
			return false
		}
		fmt.Printf("COMPARISON: %08b, %v\n", node.key, node.key&((1<<63)>>node.bit_index))

		if node.key&((1<<63)>>node.bit_index) == 0 {
			node.left = node
			node.right = ptn
		} else {
			node.right = node
			node.left = ptn
		}
		// fmt.Printf("Parent Node: %v\nNew Node: %v\n", parentNode, node)
		return true
	}
}

// 	if prev_bi < ptn.bit_index {
// 		return
// 	}
// 	if prev_bi > -1 && uint(prev_bi) >= ptn.bit_index {
// 		fmt.Printf("Returning node: %p\n", ptn)
// 		return ptn, nil
// 	}
// 	var parent, node *PatriciaTrieNode
// 	if key&(1<<ptn.bit_index) == 0 {
// 		node, parent = ptn.left.find(key, int(ptn.bit_index))
// 	} else {
// 		node, parent = ptn.right.find(key, int(ptn.bit_index))
// 	}
// 	if parent == nil {
// 		parent = ptn
// 	}
// 	fmt.Printf("Returning parent: %p\n", ptn)
// 	return node, parent
//
// }

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

func firstDiffBit(key1, key2 uint64) uint {
	var lBitPos uint = 63
	for ; lBitPos > 0 && (key1&(1<<lBitPos)) == (key2&(1<<lBitPos)); lBitPos-- {

	}
	return 63 - lBitPos
}

type PatriciaTrie struct {
	header *PatriciaTrieNode
}

func NewPatriciaTrie() *PatriciaTrie {
	return &PatriciaTrie{nil}
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
	fmt.Printf("Inserting key %v\n", key)
	if pt.header == nil {
		// HEADER IS EMPTY

		fmt.Println("Header nil so inserting manually")
		node := &PatriciaTrieNode{key, 0, nil, nil}
		node.left = node
		pt.header = node
		fmt.Printf("pt, %p\n", &pt)

	} else if pt.header == pt.header.left && key != pt.header.key {
		// ONLY HEADER EXISTS

		fmt.Println("Only header present inserting left child")
		lBitPos := firstDiffBit(key, pt.header.key)
		node := &PatriciaTrieNode{key, lBitPos, nil, nil}
		// fmt.Printf("lbit %v, ")
		if key&((1<<63)>>lBitPos) == 0 {
			node.left = node
			node.right = pt.header
		} else {
			node.right = node
			node.left = pt.header
		}
		pt.header.left = node
	} else if !pt.Search(key) {
		// HEADER AND +1 NODES

		fmt.Println("More than 2 nodes and key not present in trie")
		endNode := pt.header.left.find(key, 0)
		reachedKey := endNode.key
		lBitPos := firstDiffBit(key, reachedKey)
		node := &PatriciaTrieNode{key, lBitPos, nil, nil}

		ok := pt.header.left.insertNode(node, endNode, pt.header)
		if !ok {
			fmt.Printf("ERROR: insertion of key %v was unsuccesful\n", key)
		}
	}

	fmt.Printf("pt, %p\n", &pt)
	fmt.Printf("Key %v inserted\n", key)
}

func (pt *PatriciaTrie) Depth() int {
	if pt.header == nil {
		return 0
	}
	return pt.header.depth(0)
}

func (ptn *PatriciaTrieNode) populateTree(parent *gotree.Tree, dir string) {
	txt := fmt.Sprintf("%v %b[%v,%v]", dir, ptn.key, 63-ptn.bit_index, ptn.bit_index)
	if ptn.left != nil || ptn.right != nil {
		l := "_"
		r := "_"
		if ptn.left != nil {
			l = fmt.Sprintf("%b", ptn.left.key)
		}
		if ptn.right != nil {
			r = fmt.Sprintf("%b", ptn.right.key)
		}
		txt += fmt.Sprintf(" -> (%v, %v)", l, r)
	}
	node := (*parent).Add(txt)
	if ptn.left != nil && ptn.bit_index < ptn.left.bit_index {
		ptn.left.populateTree(&node, "L")
	}
	if ptn.right != nil && ptn.bit_index < ptn.right.bit_index {
		ptn.right.populateTree(&node, "R")
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
	// fmt.Printf("Header %v", pt.header)
	// print(pt.header)
	header := gotree.New("Header")

	pt.header.populateTree(&header, "")

	fmt.Println(header.Print())

}
