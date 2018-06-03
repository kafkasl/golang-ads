package frequent_itemsets

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

type OrderingItems struct {
	list  *[]uint
	items *map[uint]*Item
}

func (oi OrderingItems) Len() int { return len(*oi.list) }
func (oi OrderingItems) Less(i, j int) bool {
	ki, kj := (*oi.list)[i], (*oi.list)[j]
	if (*oi.items)[ki].count == (*oi.items)[kj].count {
		return ki > kj
	}
	return (*oi.items)[ki].count < (*oi.items)[kj].count
}

func (oi OrderingItems) Swap(i, j int) { (*oi.list)[i], (*oi.list)[j] = (*oi.list)[j], (*oi.list)[i] }

func orderTx(list *[]uint, itemlist *map[uint]*Item) {
	oi := OrderingItems{list, itemlist}

	sort.Sort(sort.Reverse(oi))

	return
}

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
	v   map[uint]*Item
	txs []*[]uint
}

// Data structures methods
func NewItemList(dataset *[]*[]uint, minSupport uint) *ItemList {
	var v map[uint]*Item = make(map[uint]*Item)
	var txs []*[]uint = make([]*[]uint, len(*dataset))

	for i, tx := range *dataset {
		txs[i] = tx
		for _, l := range *tx {
			if item, ok := v[l]; ok {
				item.count++
				item.Insert(tx)
			} else {
				v[l] = NewItem(l, tx)
			}

		}
	}

	for _, tx := range txs {
		var newTx []uint
		for _, l := range *tx {
			if v[l].count >= minSupport {
				newTx = append(newTx, l)

			}
			orderTx(&newTx, &v)
			*tx = newTx
		}
	}

	for _, l := range v {
		if l.count < minSupport {
			delete(v, l.element)
		}
	}

	// for _, tx := range txs {
	// 	fmt.Print("\nTX:")
	// 	for _, e := range *tx {
	// 		fmt.Printf("%v ", ToCharStr(e))
	// 	}
	// }

	return &ItemList{v, txs}
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
