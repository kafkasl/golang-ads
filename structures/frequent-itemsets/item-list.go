package structures

import (
	"fmt"
	"sort"
)

// Auxiliar Functions
func ToCharStr(i uint) string {
	return string('A' - 1 + i)
}

type ItemSlice []Item

func (p ItemSlice) Len() int { return len(p) }
func (p ItemSlice) Less(i, j int) bool {
	if p[i].count == p[j].count {
		return p[i].element > p[j].element
	}
	return p[i].count < p[j].count
}
func (p ItemSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Data structures

type ThreadedList struct {
	list *[]uint
	ptr  *ThreadedList
}

func ListToString(list []uint) string {
	str := "["
	for i, e := range list {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprintf("%v", ToCharStr(e))
	}
	str += "]"
	return str
}

func NewThreadedList(list *[]uint) *ThreadedList {
	return &ThreadedList{list, nil}
}

func (tl *ThreadedList) String() string {
	if tl.ptr != nil {
		return fmt.Sprintf("%v + %v", ListToString(*tl.list), tl.ptr)
	} else {
		return fmt.Sprintf("%v", ListToString(*tl.list))
	}
}

type Item struct {
	element uint
	count   uint
	ptr     *ThreadedList
}

func NewItem(elem uint, list *[]uint) *Item {
	i := &Item{elem, 1, nil}
	i.Insert(list)
	return i
}

func (item *Item) Insert(list *[]uint) {
	ntl := NewThreadedList(list)
	if item.ptr == nil {
		item.ptr = ntl
	} else {
		cu := item.ptr
		for ; cu.ptr != nil; cu = cu.ptr {
		}
		cu.ptr = ntl
	}
}

type ItemList struct {
	v map[uint]*Item
}

// Data structures methods
func NewItemList(dataset *[]*[]uint, minSupport uint) *ItemList {
	var v map[uint]*Item = make(map[uint]*Item)

	for _, tx := range *dataset {
		for _, l := range *tx {
			if item, ok := v[l]; ok {
				item.count++
				item.Insert(tx)
			} else {
				v[l] = NewItem(l, tx)
			}

		}
	}
	for _, l := range v {
		if l.count < minSupport {
			delete(v, l.element)
		}
	}
	return &ItemList{v}
}

func (il ItemList) String() string {
	var is ItemSlice
	for _, e := range il.v {
		is = append(is, *e)
	}
	sort.Sort(sort.Reverse(is))
	str := ""
	for _, item := range is {
		str += fmt.Sprintf("%v [%v]: %v\n", ToCharStr(item.element), item.count, item.ptr)
	}
	return str
}

// type ItemList struct {
//
// }
//
// func (ptn *PatriciaTrieNode) String() string {
// 	l, r := "_", "_"
// 	if ptn.left != nil {
// 		l = fmt.Sprintf("%b", ptn.left.key)
// 	}
// 	if ptn.right != nil {
// 		r = fmt.Sprintf("%b", ptn.right.key)
// 	}
// 	return fmt.Sprintf("[%b, %v] -> (%v, %v)", ptn.key, ptn.bit_index, l, r)
// }
//
// func (ptn *PatriciaTrieNode) search(key uint64, prev_bi uint) bool {
// 	// fmt.Printf("Search: \nCurrent node: %v\nprev_bi %v\nbit_index %v\n", ptn, prev_bi, ptn.bit_index)
//
// 	if prev_bi >= ptn.bit_index {
// 		return key == ptn.key
// 	}
// 	if key&((1<<63)>>(ptn.bit_index-1)) == 0 {
// 		return ptn.left.search(key, ptn.bit_index)
// 	} else {
// 		return ptn.right.search(key, ptn.bit_index)
// 	}
// }
//
// func (ptn *PatriciaTrieNode) find(key uint64, prev_bi uint) uint64 {
// 	// fmt.Printf("prev_bi %v, bit_index %v [%v]\n", prev_bi, ptn.bit_index, uint(prev_bi))
// 	if prev_bi >= ptn.bit_index {
// 		return ptn.key
// 	}
// 	if key&((1<<63)>>(ptn.bit_index-1)) == 0 {
// 		return ptn.left.find(key, ptn.bit_index)
// 	} else {
// 		return ptn.right.find(key, ptn.bit_index)
// 	}
// }
//
// func (ptn *PatriciaTrieNode) insertNode(node *PatriciaTrieNode, seenNodes []*PatriciaTrieNode) bool {
//
// 	var nextNode *PatriciaTrieNode
// 	if node.key&((1<<63)>>(ptn.bit_index-1)) == 0 {
// 		nextNode = ptn.left
// 	} else {
// 		nextNode = ptn.right
// 	}
//
// 	if nextNode.bit_index < node.bit_index && !contains(nextNode, seenNodes) {
//
// 		return nextNode.insertNode(node, append(seenNodes, ptn))
// 	} else {
//
// 		if ptn.right == nextNode {
// 			ptn.right = node
// 		} else if ptn.left == nextNode {
// 			ptn.left = node
// 		}
// 		node.reassignNode(node, nextNode)
//
// 		return true
// 	}
// }
//
// func (ptn *PatriciaTrieNode) depth(accDepth int) int {
// 	if !((ptn.right != nil && ptn.right.bit_index > ptn.bit_index) ||
// 		(ptn.left != nil && ptn.left.bit_index > ptn.bit_index)) {
// 		return accDepth
// 	}
// 	rd, ld := 0, 0
// 	if ptn.right != nil && ptn.right.bit_index > ptn.bit_index {
// 		rd = ptn.right.depth(accDepth + 1)
// 	}
// 	if ptn.left != nil && ptn.left.bit_index > ptn.bit_index {
// 		ld = ptn.left.depth(accDepth + 1)
// 	}
// 	if rd > ld {
// 		return rd
// 	}
// 	return ld
// }
//
// func (ptn *PatriciaTrieNode) reassignNode(node1, node2 *PatriciaTrieNode) {
// 	if ptn.key&((1<<63)>>(ptn.bit_index-1)) == 0 {
// 		ptn.left = node1
// 		ptn.right = node2
// 	} else {
// 		ptn.right = node1
// 		ptn.left = node2
// 	}
// }
//
// func firstDiffBit(key1, key2 uint64) uint {
// 	var lBitPos uint = 63
// 	for ; lBitPos > 0 && (key1&(1<<lBitPos)) == (key2&(1<<lBitPos)); lBitPos-- {
//
// 	}
// 	return 64 - lBitPos
// }
//
// type PatriciaTrie struct {
// 	header *PatriciaTrieNode
// }
//
// func NewPatriciaTrie() *PatriciaTrie {
// 	return &PatriciaTrie{nil}
// }
//
// func (pt *PatriciaTrie) Search(key uint64) bool {
// 	if pt.header == nil {
// 		return false
// 	}
// 	if pt.header == pt.header.left {
// 		return key == pt.header.key
// 	}
// 	return pt.header.left.search(key, 0)
// }
//
// func (pt *PatriciaTrie) Insert(key uint64) {
// 	if pt.header == nil {
// 		// HEADER IS EMPTY
//
// 		node := &PatriciaTrieNode{key, 0, nil, nil}
// 		node.left = node
// 		pt.header = node
// 		fmt.Printf("pt, %p\n", &pt)
//
// 	} else if pt.header == pt.header.left && key != pt.header.key {
// 		// ONLY HEADER EXISTS
//
// 		lBitPos := firstDiffBit(key, pt.header.key)
// 		node := &PatriciaTrieNode{key, lBitPos, nil, nil}
//
// 		node.reassignNode(node, pt.header)
//
// 		pt.header.left = node
//
// 	} else if !pt.Search(key) {
// 		// HEADER AND +1 NODES
//
// 		reachedKey := pt.header.left.find(key, 0)
// 		lBitPos := firstDiffBit(key, reachedKey)
// 		node := &PatriciaTrieNode{key, lBitPos, nil, nil}
//
// 		// if node needs to be inserted after header do it manually
//
// 		if node.bit_index < pt.header.left.bit_index {
// 			node.reassignNode(node, pt.header.left)
// 			pt.header.left = node
//
// 		} else {
// 			ok := pt.header.left.insertNode(node, []*PatriciaTrieNode{pt.header})
// 			if !ok {
// 				// This statement should not be reachable
// 				fmt.Printf("ERROR: insertion of key %v was unsuccesful\n", key)
// 			}
// 		}
//
// 	}
//
// }
//
// func (pt *PatriciaTrie) Depth() int {
// 	if pt.header == nil {
// 		return 0
// 	}
// 	return pt.header.depth(1)
// }
//
// func (ptn *PatriciaTrieNode) populateTree(parent *gotree.Tree, dir string, binary bool) {
// 	var txt string
// 	if binary {
// 		txt = fmt.Sprintf("%v %b[%v]", dir, ptn.key, ptn.bit_index)
// 	} else {
// 		txt = fmt.Sprintf("%v %v[%v]", dir, ptn.key, ptn.bit_index)
// 	}
// 	if ptn.left != nil || ptn.right != nil {
// 		l := "_"
// 		r := "_"
// 		if ptn.left != nil {
// 			if binary {
// 				l = fmt.Sprintf("%b", ptn.left.key)
// 			} else {
// 				l = fmt.Sprintf("%v", ptn.left.key)
// 			}
// 		}
// 		if ptn.right != nil {
// 			if binary {
// 				l = fmt.Sprintf("%b", ptn.right.key)
// 			} else {
// 				l = fmt.Sprintf("%v", ptn.right.key)
// 			}
// 		}
// 		txt += fmt.Sprintf(" -> (%v, %v)", l, r)
// 	}
// 	node := (*parent).Add(txt)
// 	if ptn.left != nil && ptn.bit_index < ptn.left.bit_index {
// 		ptn.left.populateTree(&node, "L", binary)
// 	}
// 	if ptn.right != nil && ptn.bit_index < ptn.right.bit_index {
// 		ptn.right.populateTree(&node, "R", binary)
// 	}
// }
//
// func (pt *PatriciaTrie) String() string {
//
// 	header := gotree.New("Header")
//
// 	pt.header.populateTree(&header, "", true)
//
// 	return header.Print()
//
// }
//
// func (pt *PatriciaTrie) toString(binary bool) string {
//
// 	header := gotree.New("Header")
//
// 	pt.header.populateTree(&header, "", binary)
//
// 	return header.Print()
//
// }