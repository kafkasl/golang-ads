## Golang Advanced Data Structures


### Tests

Run all the tests, examples and performance evaluations with:

`./bin/run_tests_and_examples.sh`

### Documentation

Detailed documentation and explanation of each data structure and the performance tests is available in:

`./doc/golang_ads.pdf`

### Available Structures

#### Trie

* Location: `./structures/trie`
* Description:  Implements Trie Structure using as key any Go Rune (and no alphabet size required)
* Examples: Trie examples are the public examples of problem http://codeforces.com/contest/456/problem/D
* You can find the actual submission code in `./submissions/submission456D.go`


#### Patricia trie

* Location: `./structures/patricia-trie`
* Description:  Implements Patricia trie Structure as described in Handbook of  Data Structures and Applications book.


#### Union-find Set with Path Compression / with Rank

* Location: `./structures/union-find/`
* Description:  Implements union-find sets. Two variants are available: one using path compression using vectors (most used and efficient); another one using rank for set-union and with pointers.
* Examples: Union-find examples are the public examples of problem https://jutge.org/problems/P94041_ca/statement and the implementation of Kruskal's minimum spanning tree algorithm
* You can find the submission code in `./submissions/submissionConnectedComponents.go`


#### Skip list

* Location: `./structures/trie`
* Description:  Implements a map using the skip list structure (instead of a search tree)
* Examples: `./evaluation/skip-list_performance.go` does a complexity evaluation (both time and space) of the implemented structure. Base package example just tests String() method.  


#### Frequent Itemsets

* Location  `./structures/frequent-itemsets`
* Description:  uses item lists, trie, and patricia tries to represent a frequent itemset
* Examples: `./evaluation/fis_size-tests.go` creates represents each dataset under `./data` and compares the total memory size of each one of them.
