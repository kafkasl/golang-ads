package union_find

import "fmt"

type UnionFinSetwRP struct {
	parent *UnionFinSetwRP
	rank   int
	value  int
}

func NewUnionFinSetwRP(value int) *UnionFinSetwRP {
	uf := UnionFinSetwRP{nil, 0, value}
	uf.parent = &uf
	return &uf
}

func (uf *UnionFinSetwRP) Find() *UnionFinSetwRP {
	if uf.parent == uf {
		return uf
	}
	return uf.parent.Find()

}

func (uf *UnionFinSetwRP) Union(other *UnionFinSetwRP) *UnionFinSetwRP {
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

func (uf *UnionFinSetwRP) String() string {
	if uf.parent == uf {
		return fmt.Sprintf("[%v]: %v", uf.rank, uf.value)
	}
	str := fmt.Sprintf("[%v]: %v -> %v", uf.rank, uf.value, uf.parent)
	return str
}
