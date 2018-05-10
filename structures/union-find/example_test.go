package structures

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// func canWinLose(node *TrieNode) (cw bool, cl bool) {
// 	if len(node.children) == 0 {
// 		return false, true
// 	}
// 	cw, cl = false, false
// 	for _, child := range node.children {
// 		cw_aux, cl_aux := canWinLose(child)
// 		cw = cw || !cw_aux
// 		cl = cl || !cl_aux
// 	}
// 	return cw, cl
// }
//
// func solveGame(n, k int, trie Trie) string {
//
// 	rootCanWin, rootCanLose := canWinLose(trie.root)
//
// 	if rootCanWin && rootCanLose {
// 		return "First"
// 	} else if rootCanWin {
// 		if (k % 2) == 1 {
// 			return "First"
// 		} else {
// 			return "Second"
// 		}
// 	} else {
// 		return "Second"
// 	}
// }
//
// Auxuliar structure to hold vertices and their weights
type VertexWeight struct {
	origin  int
	destiny int
	weight  int
}

// Implementing sort interface to sort VertexWeight instances
type byWeight []VertexWeight

func (s byWeight) Len() int {
	return len(s)
}
func (s byWeight) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byWeight) Less(i, j int) bool {
	return s[i].weight < s[j].weight
}

// Parsing input. Read number of nodes n. After, each line contains: 'origin destiny weight' of each vertex.
func parseInput(input string) (adjMatrix [][]int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	adjMatrix = make([][]int, n)
	for i := range adjMatrix {
		adjMatrix[i] = make([]int, n)
	}

	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		origin, _ := strconv.Atoi(vals[0])
		destiny, _ := strconv.Atoi(vals[1])
		weight, _ := strconv.Atoi(vals[2])
		adjMatrix[origin][destiny] = weight

	}
	return
}

func adjMatrixToPriorityQ(adjMatrix [][]int) []VertexWeight {
	var vws []VertexWeight
	for i := 0; i < len(adjMatrix); i++ {
		if adjMatrix[i][i] != 0 {
			fmt.Printf("Vertices can only connect 2 different nodes, ignoring vertex(%v, %v)", i, i)
		}
		for j := i + 1; j < len(adjMatrix[i]); j++ {
			if adjMatrix[i][j] > 0 {
				aux := VertexWeight{i, j, adjMatrix[i][j]}
				vws = append(vws, aux)

			}
		}
	}

	sort.Sort(byWeight(vws))
	return vws
}

func kruskal(n int, vws []VertexWeight) (mst []VertexWeight) {
	set := make([]*UnionFindSet, n)
	for i := 0; i < n; i++ {
		set[i] = NewUnionFindSet(i)
	}
	for i := 0; i < len(vws); i++ {
		vw := vws[i]
		u, v, _ := vw.origin, vw.destiny, vw.weight
		if set[u].Find() != set[v].Find() {
			mst = append(mst, vws[i])
			set[u].Union(set[v])
		}
	}
	return
}

func ExampleKruskal1() {
	input := "5\n0 1 3\n0 4 1\n1 2 5\n1 4 4\n2 4 6\n2 3 2\n3 4 7\n"

	adjMatrix := parseInput(input)

	vws := adjMatrixToPriorityQ(adjMatrix)

	mst := kruskal(len(adjMatrix), vws)

	for _, vw := range mst {
		fmt.Printf("%v -> %v [%v] ", vw.origin, vw.destiny, vw.weight)
	}

	// Output: 0 -> 4 [1] 2 -> 3 [2] 0 -> 1 [3] 1 -> 2 [5]

}

//
// func ExampleGame2() {
// 	input := "3 1\na\nb\nc"
//
// 	n, k, trie := parseInput(input)
//
// 	fmt.Printf("%v", solveGame(n, k, trie))
// 	// Output: First
// }
//
// func ExampleGame3() {
// 	input := "1 2\nab"
//
// 	n, k, trie := parseInput(input)
//
// 	fmt.Printf("%v", solveGame(n, k, trie))
// 	// Output: Second
// }
