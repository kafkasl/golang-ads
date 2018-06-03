package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/OneOfOne/go-utils/memory"
)

func readDataset(filename string) (dataset []*[]uint) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		var line []uint
		for _, val := range vals {
			el, _ := strconv.Atoi(val)
			line = append(line, uint(el))
		}
		dataset = append(dataset, &line)

	}
	return
}
func main() {

	datasets := []string{"../../data/chess.dat",
        "../../data/connect.dat",
		"../../data/pumsb.dat",
        "../../data/pumsb_star.dat",
        "../../data/mushroom.dat",
		"../../data/T10I4D100K.dat",
        "../../data/T40I10D100K.dat"}

	for _, filename := range datasets {
		dataset := readDataset(filename)
		fmt.Printf("Processing %v\n", filename)
		il := NewItemList(&dataset, 3)
		fmt.Printf("Size of Itemlist: %v\n", memory.Sizeof(il))

		t := NewTrie()
		for _, tx := range il.txs {
			strTx := ""
			for _, e := range *tx {
				strTx += ToCharStr(e)
			}
			t.Insert(strTx)
		}

		// fmt.Printf("%v", t)
		fmt.Printf("Size of Trie: %v\n", memory.Sizeof(t))

		pt := NewFISPatriciaTrie(t)
		fmt.Printf("Size of Patricia Trie: %v\n", memory.Sizeof(pt))

	}

}
