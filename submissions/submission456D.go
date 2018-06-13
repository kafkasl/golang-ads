package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	tn := TrieNode{make(map[rune]*TrieNode), false}
	return Trie{&tn}
}

func (t Trie) Insert(key string) {
	currentNode := t.root
	for _, letter := range key[:] {
		if val, ok := currentNode.children[letter]; ok {
			currentNode = val
		} else {
			nn := &TrieNode{make(map[rune]*TrieNode), false}
			currentNode.children[letter] = nn
			currentNode = nn
		}
	}
	currentNode.endOfWord = true
}

func (t Trie) Search(key string) bool {
	currentNode := t.root
	for _, letter := range key[:] {
		if val, ok := currentNode.children[letter]; ok {
			currentNode = val
		} else {
			return false
		}
	}
	return currentNode.endOfWord
}

func canWinLose(node *TrieNode) (cw bool, cl bool) {
	if len(node.children) == 0 {
		return false, true
	}
	cw, cl = false, false
	for _, child := range node.children {
		cw_aux, cl_aux := canWinLose(child)
		cw = cw || !cw_aux
		cl = cl || !cl_aux
	}
	return cw, cl
}

func solveGame(n, k int, trie Trie) string {

	rootCanWin, rootCanLose := canWinLose(trie.root)

	if rootCanWin && rootCanLose {
		return "First"
	} else if rootCanWin {
		if (k % 2) == 1 {
			return "First"
		} else {
			return "Second"
		}
	} else {
		return "Second"
	}
}

func parseInput() (n, k int, trie Trie) {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	scanner.Scan()
	vals := strings.Split(scanner.Text(), " ")
	n, _ = strconv.Atoi(vals[0])
	k, _ = strconv.Atoi(vals[1])
	trie = NewTrie()
	for scanner.Scan() {
		trie.Insert(scanner.Text())
	}
	return
}

func main() {
	n, k, trie := parseInput()
	fmt.Println(solveGame(n, k, trie))
}
