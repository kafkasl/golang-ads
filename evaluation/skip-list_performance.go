package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"testing"

	"github.com/OneOfOne/go-utils/memory"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func PrintFatal(t *testing.T, expected fmt.Stringer, got fmt.Stringer) {
	t.Fatalf("Expected: \n%v, but got: \n%v\n", expected, got)
}

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
	length int
}

func NewDictionary(p float32) *Dictionary {
	start := NewSkipListNode(math.MinInt64, "start", 1)
	end := NewSkipListNode(math.MaxInt64, "end", 1)
	start.next[0] = end
	end.next[0] = nil
	return &Dictionary{start, 1, p, 0}
}

func (d *Dictionary) Lookup(key int) (*SkipListNode, int) {
	p := d.header
	l := d.height - 1
	searchLength := 0

	for l >= 0 {
		if p.next[l] == nil || key <= p.next[l].key {
			l--
			searchLength++
		} else {
			p = p.next[l]
			searchLength++
		}
	}
	if p.next[0] == nil || p.next[0].key != key {
		//key is not present
		return nil, searchLength
	} else {
		return p.next[0], searchLength
	}
}

func (d *Dictionary) Delete(key int) bool {
	p := d.header
	l := d.height - 1
	pred := make([]*SkipListNode, d.height)
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
		return false

	} else {
		for i := d.height - 1; i >= 0; i-- {
			if pred[i].next[i] == p.next[0] {
				pred[i].next[i] = p.next[0].next[i]
			}
		}
		for i := d.height - 1; i > 0; i-- {
			if d.header.next[i] == nil || d.header.next[i].key == math.MaxInt64 {
				d.height--
			}
		}
		d.length--
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
		for rand.Float32() > d.p {
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
		d.length++

	} else {
		p.next[0].value = value
	}
}
func (d *Dictionary) Len() int {
	return d.length
}

func (d *Dictionary) String() (str string) {
	str = fmt.Sprintf("Skip List (height: %v)\n", d.height)
	for p := d.header; p != nil; p = p.next[0] {
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
		str += "\n"
	}

	return
}

func doSearches(p float32, n, totalSearches int, wg *sync.WaitGroup) {
	defer wg.Done()
	accLength := 0
	accHeight := 0
	var accSize uint64 = 0
	exps := 20
	q := 1 - p

	for i := 0; i < exps; i++ {

		d := NewDictionary(p)

		for i := 0; i < n; i++ {
			k, v := rand.Int(), ""
			d.Insert(k, v)
		}

		for i := 0; i < totalSearches; i++ {
			k := rand.Int()
			_, length := d.Lookup(k)
			accLength += length
		}
		accHeight += d.height
		accSize += memory.Sizeof(d)
	}

	avgLength := float64(accLength) / float64(totalSearches*exps)
	avgHeight := float64(accHeight) / float64(exps)
	avgSize := float64(accSize) / float64(exps)
	fmt.Printf("\nResults (p=%v):\n * Height: %v\n * Average path length: %.2f\n * SizeOf(d): %v\n",
		p, avgHeight, avgLength, avgSize)
	q1 := float64(1.0 / q)
	eplb := q1 * (math.Log(float64(n)+q1) / math.Log(q1))
	fmt.Printf(" * E[length(searchPath)] <= %.2f\n", eplb)
	fmt.Printf("(%v, %v, %.2f, %.2f, %v)\n", p, avgHeight, avgLength, eplb, avgSize)
}

func main() {

	testsNum := 1
	n := 100000
	totalSearches := 100000
	var wg sync.WaitGroup

	var ps []float32 = []float32{0.05, 0.1, 0.15, 0.2, 0.25, 0.3, 0.35, 0.4,
		0.45, 0.5, 0.55, 0.6, 0.65, 0.7, 0.75, 0.8, 0.85, 0.9, 0.95}

	wg.Add(len(ps))
	for _, p := range ps[:] {

		q := 1 - p
		var pr float64 = 1
		for i := 1; i <= 20 && pr >= 0.01; i++ {
			pr = math.Pow(float64(q), float64(i-1))
			fmt.Printf("P[height(X) => %v] = %.2f\n", i, pr)
		}
		pr = 1
		for i := 1; i <= 20 && pr >= 0.01; i++ {
			pr = float64(n) * math.Pow(float64(q), float64(i-1))
			fmt.Printf("E[# of nodes with height(X) => %v] = %.2f\n", i, pr)
		}

		for i := 0; i < testsNum; i++ {
			go doSearches(p, n, totalSearches, &wg)
		}
	}

	wg.Wait()

}
