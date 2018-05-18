package structures

import (
	"fmt"

	"github.com/disiqueira/gotree"
)

func contains(a *PatriciaTrieNode, list []*PatriciaTrieNode) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

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

func (ptn *PatriciaTrieNode) search(key uint64, prev_bi uint) bool {
	// fmt.Printf("Search: \nCurrent node: %v\nprev_bi %v\nbit_index %v\n", ptn, prev_bi, ptn.bit_index)

	if prev_bi >= ptn.bit_index {
		return key == ptn.key
	}
	if key&((1<<63)>>ptn.bit_index) == 0 {
		return ptn.left.search(key, ptn.bit_index)
	} else {
		return ptn.right.search(key, ptn.bit_index)
	}
}

func (ptn *PatriciaTrieNode) find(key uint64, prev_bi uint) uint64 {
	// fmt.Printf("prev_bi %v, bit_index %v [%v]\n", prev_bi, ptn.bit_index, uint(prev_bi))
	if prev_bi >= ptn.bit_index {
		return ptn.key
	}
	if key&((1<<63)>>ptn.bit_index) == 0 {
		return ptn.left.find(key, ptn.bit_index)
	} else {
		return ptn.right.find(key, ptn.bit_index)
	}
}

func (ptn *PatriciaTrieNode) insertNode(node *PatriciaTrieNode, seenNodes []*PatriciaTrieNode) bool {

	var nextNode *PatriciaTrieNode
	if node.key&((1<<63)>>ptn.bit_index) == 0 {
		nextNode = ptn.left
	} else {
		nextNode = ptn.right
	}
	fmt.Printf("\n\n\nSTART\nCurrent Node: %v\nNode: %v\nInsert node: %v\n\n\n", ptn, nextNode, node)

	// if nextNode.bit_index < node.bit_index && ptn != endNode {
	// if nextNode.bit_index < node.bit_index && parentNode.bit_index < ptn.bit_index {
	if nextNode.bit_index < node.bit_index && !contains(nextNode, seenNodes) {

		return nextNode.insertNode(node, append(seenNodes, ptn))
	} else {
		// if node.bit_index > ptn.bit_index {
		// 	parentNode = ptn
		// 	ptn = nextNode
		// }
		// fmt.Printf("\n\n\nParent Node: %v\nNext Node: %v\nNode: %v\nCurrent node: %v\n\n\n", parentNode, nextNode, node, ptn)
		fmt.Printf("\n\n\nCurrent Node: %v\nNode: %v\nInsert node: %v\n\n\n", ptn, nextNode, node)

		if ptn.right == nextNode {
			ptn.right = node
		} else if ptn.left == nextNode {
			ptn.left = node
		} else {
			fmt.Println("ERROR: I am no son of my parent")
			return false
		}

		fmt.Printf("Comparison: %b & %b = %b (bit index: %v)\n", node.key, uint64((1<<63)>>node.bit_index), node.key&((1<<63)>>node.bit_index), node.bit_index)
		if node.key&((1<<63)>>node.bit_index) == 0 {
			node.left = node
			node.right = nextNode
		} else {
			node.right = node
			node.left = nextNode
		}
		// fmt.Printf("Parent Node: %v\nNew Node: %v\n", parentNode, node)
		return true
	}
}

func (ptn *PatriciaTrieNode) depth(accDepth int) int {
	fmt.Printf("right %v\nleft %v\nptn: %v\n", ptn.right, ptn.left, ptn)
	if !((ptn.right != nil && ptn.right.bit_index > ptn.bit_index) ||
		(ptn.left != nil && ptn.left.bit_index > ptn.bit_index)) {
		return accDepth
	}
	rd, ld := 0, 0
	if ptn.right != nil && ptn.right.bit_index > ptn.bit_index {
		rd = ptn.right.depth(accDepth + 1)
	}
	if ptn.left != nil && ptn.left.bit_index > ptn.bit_index {
		ld = ptn.left.depth(accDepth + 1)
	}
	if rd > ld {
		return rd
	}
	return ld
}

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
		if node.key&((1<<63)>>node.bit_index) == 0 {
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
		reachedKey := pt.header.left.find(key, 0)
		lBitPos := firstDiffBit(key, reachedKey)
		node := &PatriciaTrieNode{key, lBitPos, nil, nil}

		// if node needs to be inserted after header do it manually
		if node.bit_index < pt.header.left.bit_index {
			fmt.Println("Insta inserting")
			fmt.Printf("Comparison: %b & %b = %b (bit index: %v)\n", node.key, uint64((1<<63)>>node.bit_index), node.key&((1<<63)>>node.bit_index), node.bit_index)

			if node.key&((1<<63)>>node.bit_index) == 0 {
				node.left = node
				node.right = pt.header.left
			} else {
				node.right = node
				node.left = pt.header.left
			}
			fmt.Printf("Final header: %v -> %v, _\nNew node: %v -> %v, %v\n%v\n",
				pt.header.key, pt.header.left.key, node.key, node.left.key, node.right.key, pt.toString(false))
			// fmt.Printf("Final header: %v, %v\n",
			// 	pt.header.key, pt.header)
			// fmt.Printf("New node: %v -> %v, %v\n%v\n", pt.header.key, pt.header.left.key, pt.header.right.key, node.key, node.left.key, node.right.key, pt.toString(false))

			fmt.Printf("%v\n", pt.toString(false))

			pt.header.left = node
			fmt.Printf("Final header: %v -> %v, _\nNew node: %v -> %v, %v\n%v\n",
				pt.header.key, pt.header.left.key, node.key, node.left.key, node.right.key, pt.toString(false))

			fmt.Printf("%v\n", pt.toString(false))

		} else {
			ok := pt.header.left.insertNode(node, []*PatriciaTrieNode{pt.header})
			if !ok {
				fmt.Printf("ERROR: insertion of key %v was unsuccesful\n", key)
			}
		}

	}

	fmt.Printf("pt, %p\n", &pt)
	fmt.Printf("Key %v inserted\n", key)
}

func (pt *PatriciaTrie) Depth() int {
	if pt.header == nil {
		return 0
	}
	return pt.header.depth(1)
}

func (ptn *PatriciaTrieNode) populateTree(parent *gotree.Tree, dir string, binary bool) {
	var txt string
	if binary {
		txt = fmt.Sprintf("%v %b[%v]", dir, ptn.key, ptn.bit_index)
	} else {
		txt = fmt.Sprintf("%v %v[%v]", dir, ptn.key, ptn.bit_index)
	}
	if ptn.left != nil || ptn.right != nil {
		l := "_"
		r := "_"
		if ptn.left != nil {
			if binary {
				l = fmt.Sprintf("%b", ptn.left.key)
			} else {
				l = fmt.Sprintf("%v", ptn.left.key)
			}
		}
		if ptn.right != nil {
			if binary {
				l = fmt.Sprintf("%b", ptn.right.key)
			} else {
				l = fmt.Sprintf("%v", ptn.right.key)
			}
		}
		txt += fmt.Sprintf(" -> (%v, %v)", l, r)
	}
	node := (*parent).Add(txt)
	if ptn.left != nil && ptn.bit_index < ptn.left.bit_index {
		ptn.left.populateTree(&node, "L", binary)
	}
	if ptn.right != nil && ptn.bit_index < ptn.right.bit_index {
		ptn.right.populateTree(&node, "R", binary)
	}
}

func (pt *PatriciaTrie) String() string {

	header := gotree.New("Header")

	pt.header.populateTree(&header, "", true)

	return header.Print()

}

func (pt *PatriciaTrie) toString(binary bool) string {

	header := gotree.New("Header")

	pt.header.populateTree(&header, "", binary)

	return header.Print()

}
