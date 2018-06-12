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
