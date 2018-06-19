package union_find

import "fmt"

type UnionFindSetwRP struct {
	parent *UnionFindSetwRP
	rank   int
	value  int
}

func NewUnionFindSetwRP(value int) *UnionFindSetwRP {
	uf := UnionFindSetwRP{nil, 0, value}
	uf.parent = &uf
	return &uf
}

func (uf *UnionFindSetwRP) Find() *UnionFindSetwRP {
	if uf.parent == uf {
		return uf
	}
	return uf.parent.Find()

}

func (uf *UnionFindSetwRP) Union(other *UnionFindSetwRP) *UnionFindSetwRP {
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

func (uf *UnionFindSetwRP) String() string {
	if uf.parent == uf {
		return fmt.Sprintf("[%v]: %v", uf.rank, uf.value)
	}
	str := fmt.Sprintf("[%v]: %v -> %v", uf.rank, uf.value, uf.parent)
	return str
}
