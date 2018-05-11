package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Auxuliar structure to hold edges and their weights
type EdgeWeight struct {
	origin  int
	destiny int
	weight  int
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
	if otherParent.rank > ufParent.rank {
		ufParent.parent = otherParent
		return otherParent
	} else {
		otherParent.parent = ufParent
		if otherParent.rank == ufParent.rank {
			ufParent.rank++
		}
		return ufParent
	}

}

func initUnionSets(n int) (set []*UnionFindSet) {
	set = make([]*UnionFindSet, n)
	for i := 0; i < n; i++ {
		set[i] = NewUnionFindSet(i)
	}
	return
}

func findConnectedComponents(sets []*UnionFindSet, edges []EdgeWeight) (result []int) {
	cc := len(sets)
	for i := 0; i < len(edges); i++ {
		u, v := edges[i].origin, edges[i].destiny
		if sets[u].Find() != sets[v].Find() {
			cc--
		}
		sets[u].Union(sets[v])
		result = append(result, cc)

	}
	return
}

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(vals[0])
		m, _ := strconv.Atoi(vals[1])
		var edges []EdgeWeight

		for i := 0; i < m; i++ {
			scanner.Scan()
			vals := strings.Split(scanner.Text(), " ")
			origin, _ := strconv.Atoi(vals[0])
			destiny, _ := strconv.Atoi(vals[1])
			edges = append(edges, EdgeWeight{origin, destiny, 1})

		}

		sets := initUnionSets(n)

		result := findConnectedComponents(sets, edges)

		for i, r := range result {
			if i > 0 {
				fmt.Printf(" ")
			}
			fmt.Printf("%v", r)
		}
		fmt.Printf("\n")
		scanner.Scan()
	}

}
