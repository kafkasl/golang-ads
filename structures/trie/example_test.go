package structures

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

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

func parseInput(input string) (n, k int, trie Trie) {
	scanner := bufio.NewScanner(strings.NewReader(input))
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

func ExampleGame1() {
	input := "2 3\na\nb"

	n, k, trie := parseInput(input)

	fmt.Printf("%v", solveGame(n, k, trie))
	// Output: First
}

func ExampleGame2() {
	input := "3 1\na\nb\nc"

	n, k, trie := parseInput(input)

	fmt.Printf("%v", solveGame(n, k, trie))
	// Output: First
}

func ExampleGame3() {
	input := "1 2\nab"

	n, k, trie := parseInput(input)

	fmt.Printf("%v", solveGame(n, k, trie))
	// Output: Second
}
