package union_find

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Auxuliar structure to hold edges and their weights
type EdgeWeight struct {
	origin  int
	destiny int
	weight  int
}

// Implementing sort interface to sort EdgeWeight instances
type byWeight []EdgeWeight

func (s byWeight) Len() int {
	return len(s)
}
func (s byWeight) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byWeight) Less(i, j int) bool {
	return s[i].weight < s[j].weight
}

func initUnionSets(n int) (set []*UnionFindSet) {
	set = make([]*UnionFindSet, n)
	for i := 0; i < n; i++ {
		set[i] = NewUnionFindSet(i)
	}
	return
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

func adjMatrixToPriorityQ(adjMatrix [][]int) []EdgeWeight {
	var vws []EdgeWeight
	for i := 0; i < len(adjMatrix); i++ {
		if adjMatrix[i][i] != 0 {
			fmt.Printf("Vertices can only connect 2 different nodes, ignoring vertex(%v, %v)", i, i)
		}
		for j := i + 1; j < len(adjMatrix[i]); j++ {
			if adjMatrix[i][j] > 0 {
				aux := EdgeWeight{i, j, adjMatrix[i][j]}
				vws = append(vws, aux)

			}
		}
	}

	sort.Sort(byWeight(vws))
	return vws
}

func kruskal(n int, vws []EdgeWeight) (mst []EdgeWeight) {
	set := NewUnionFindSet(n)

	for i := 0; i < len(vws); i++ {
		vw := vws[i]
		u, v, _ := vw.origin, vw.destiny, vw.weight
		if set.Find(u) != set.Find(v) {
			mst = append(mst, vws[i])
			set.Union(u, v)
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

func parseInputCC(input string) (n, m int, edges []EdgeWeight) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	vals := strings.Split(scanner.Text(), " ")
	n, _ = strconv.Atoi(vals[0])
	m, _ = strconv.Atoi(vals[1])

	for i := 0; i < m; i++ {
		scanner.Scan()
		vals := strings.Split(scanner.Text(), " ")
		origin, _ := strconv.Atoi(vals[0])
		destiny, _ := strconv.Atoi(vals[1])
		edges = append(edges, EdgeWeight{origin, destiny, 1})

	}
	return
}

func findConnectedComponents(sets *UnionFindSet, edges []EdgeWeight) (result []int) {
	cc := len(sets.g)
	for i := 0; i < len(edges); i++ {
		u, v := edges[i].origin, edges[i].destiny
		if sets.Find(u) != sets.Find(v) {
			cc--
		}
		sets.Union(u, v)
		result = append(result, cc)

	}
	return
}

func ExampleComponentsConnexes1() {
	input := "4 5\n0 1\n0 2\n1 2\n3 2\n3 1\n"

	n, _, edges := parseInputCC(input)

	sets := NewUnionFindSet(n)

	result := findConnectedComponents(sets, edges)

	for i, r := range result {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%v", r)
	}
	// Output: 3 2 2 1 1
}
func ExampleComponentsConnexes2() {
	input := "100000 4\n17 751\n17 1024\n0 99999\n1024 751\n"

	n, _, edges := parseInputCC(input)

	sets := NewUnionFindSet(n)

	result := findConnectedComponents(sets, edges)
	for i, r := range result {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%v", r)
	}
	// Output: 99999 99998 99997 99997
}
