package structures

import "fmt"

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

func (uf *UnionFindSet) String() string {
	if uf.parent == uf {
		return fmt.Sprintf("[%v]: %v", uf.rank, uf.value)
	}
	str := fmt.Sprintf("[%v]: %v -> %v", uf.rank, uf.value, uf.parent)
	return str
}
