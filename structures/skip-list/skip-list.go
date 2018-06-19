package skip_list

import (
	"fmt"
	"math"
	"math/rand"
)

type SkipListNode struct {
	key   int
	value string
	next  []*SkipListNode
}

func NewSkipListNode(k int, v string, h int) *SkipListNode {
	return &SkipListNode{k, v, make([]*SkipListNode, h, h*2)}
}

func (sln SkipListNode) String() (str string) {
	k := fmt.Sprintf("%v", sln.key)
	if sln.key == math.MaxInt64 {
		k = "+∞"
	} else if sln.key == math.MinInt64 {
		k = "-∞"
	}
	return fmt.Sprintf("[%v: %v (%v)]", k, sln.value, len(sln.next))
}

type Dictionary struct {
	header *SkipListNode
	height int
	p      float32
}

func NewDictionary(p float32) *Dictionary {

	start := NewSkipListNode(math.MinInt64, "start", 1)
	end := NewSkipListNode(math.MaxInt64, "end", 1)
	start.next[0] = end
	end.next[0] = nil
	return &Dictionary{start, 1, p}
}

func (d *Dictionary) Lookup(key int) *SkipListNode {
	p := d.header
	l := d.height - 1

	for l >= 0 {
		if p.next[l] == nil || key <= p.next[l].key {
			l--
		} else {
			p = p.next[l]
		}
	}
	if p.next[0] == nil || p.next[0].key != key {
		//key is not present
		return nil
	} else {
		return p.next[0]
	}
}

func (d *Dictionary) Delete(key int) bool {
	p := d.header
	l := d.height - 1
	pred := make([]*SkipListNode, d.height)
	pred[0] = d.header
	// find the key
	for l >= 0 {
		if p.next[l] == nil || key <= p.next[l].key {
			pred[l] = p
			l--
		} else {
			p = p.next[l]
		}
	}
	// key was present, return false
	if p.next[0] == nil || p.next[0].key != key {
		return false
		// key is present remove
	} else {

		// rewire predecessors
		for i := d.height - 1; i >= 0; i-- {
			if pred[i].next[i] == p.next[0] {
				pred[i].next[i] = p.next[0].next[i]
			}
		}
		//
		for i := d.height - 1; i > 0; i-- {
			if d.header.next[i] == nil || d.header.next[i].key == math.MaxInt64 {
				d.height--
			}
		}
		return true
	}
}

func (d *Dictionary) Insert(key int, value string) {
	p := d.header
	l := d.height - 1
	pred := make([]*SkipListNode, d.height, 2*d.height)
	pred[0] = d.header
	for l >= 0 {
		if p.next[l] == nil || key <= p.next[l].key {
			pred[l] = p
			l--
		} else {
			p = p.next[l]
		}
	}
	if p.next[0] == nil || p.next[0].key != key {
		// key not found, add new node
		h := 1
		for rand.Float32() > d.p && d.p > 0 {
			h++
		}
		nn := NewSkipListNode(key, value, h)
		if h > d.height {
			if h > cap(d.header.next) || h > cap(pred) {
				dhn := make([]*SkipListNode, h, 2*h)
				p := make([]*SkipListNode, h, 2*h)
				_ = copy(dhn, d.header.next)
				_ = copy(p, pred)
				d.header.next = dhn
				pred = p
			} else {
				d.header.next = d.header.next[:h]
				pred = pred[:h]
			}
			for i, _ := range pred[:] {
				if pred[i] == nil {
					pred[i] = d.header
				}
			}
		}
		for i := h - 1; i >= 0; i-- {
			nn.next[i] = pred[i].next[i]
			pred[i].next[i] = nn
		}
		if h > d.height {
			d.height = h
		}
	} else {
		p.next[0].value = value
	}
}

func (d *Dictionary) String() (str string) {
	str = fmt.Sprintf("Skip List (height: %v)\n", d.height)
	for p := d.header; p != nil; p = p.next[0] {
		str += "\n"

		for i := 0; i < d.height; i++ {
			if i < len(p.next) || p.key == math.MaxInt64 {
				str += "[ ] "
			} else {
				str += " |  "
			}
		}
		str += p.String() + "\n"

		if p.next[0] != nil {
			for i := 0; i < d.height; i++ {
				if i <= len(p.next) {
					str += " |  "
				} else {
					str += " |  "
				}

			}
			str += "\n"
			for i := 0; i < d.height; i++ {
				if i < len(p.next[0].next) || p.next[0].key == math.MaxInt64 {
					str += " V  "
				} else {
					str += " |  "
				}

			}
		}
	}

	return
}
