package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/OneOfOne/go-utils/memory"
	. "github.com/kafkasl/golang-ads/structures/frequent-itemsets"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
		for _, val := range vals[:] {
			el, _ := strconv.Atoi(val)
			line = append(line, uint(el))
		}
		dataset = append(dataset, &line)

	}
	return
}

func processDataset(filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	f := strings.Split(filename, "/")
	file := f[len(f)-1]

	p := message.NewPrinter(language.English)

	dataset := readDataset(filename)
	fmt.Printf("Processing %v\n", filename)
	il := NewItemList(&dataset, 3)

	p.Printf("[%v] Size of Itemlist: %v\n", file, memory.Sizeof(il))

	t := NewTrie()
	for _, tx := range il.Transactions()[:] {
		strTx := ""
		for _, e := range (*tx)[:] {
			strTx += ToCharStr(e)
		}
		t.Insert(strTx)
	}

	p.Printf("[%v] Size of Trie: %v\n", file, memory.Sizeof(t))

	pt := NewFISPatriciaTrie(t)
	p.Printf("[%v] Size of Patricia Trie: %v\n", file, memory.Sizeof(pt))

}

func main() {
	datasets := []string{"../data/chess.dat",
		"../data/pumsb.dat",
		"../data/pumsb_star.dat",
		"../data/mushroom.dat",
		"../data/connect.dat"}

	var wg sync.WaitGroup
	wg.Add(len(datasets))

	for _, filename := range datasets[:] {
		go processDataset(filename, &wg)
	}

	wg.Wait()

}
