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

//
// func (il ItemList) Frequencies() []uint {
// 	var is ItemSlice
// 	for _, e := range il.v {
// 		is = append(is, *e)
// 	}
// 	sort.Sort(sort.Reverse(is))
//
//     frequencies :=
// }
