package structures

var SON_SYMBOL = '├'
var LAST_SON_SYMBOL = '└'

type TrieNode struct {
	children  map[rune]*TrieNode
	endOfWord bool
}

func (tn TrieNode) Words(currentWord string) []string {
	words := []string{}
	if tn.endOfWord {
		return append(words, currentWord)
	}
	if len(tn.children) > 0 {
		for key, child := range tn.children {
			words = append(words, child.Words(currentWord+string(key))...)
		}
	}
	return words
}

func (tn TrieNode) toString(prefix string) string {
	var text string = ""
	var nprefix string

	idx := 0
	for key, child := range tn.children {

		if idx < len(tn.children)-1 {
			text += prefix + " ├─ " + string(key) + "\n"
			nprefix = prefix + " │ "
		} else {
			text += prefix + " └─ " + string(key) + "\n"
			nprefix = prefix + "   "
		}
		text += child.toString(nprefix)
		idx++
	}
	return text
}
func (tn TrieNode) String() string {
	return tn.toString("")
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	tn := TrieNode{make(map[rune]*TrieNode), false}
	return Trie{&tn}
}

func (t Trie) String() string {
	return "Trie:\n" + t.root.String()
}

func (t Trie) Words() []string {
	return t.root.Words("")
}

func (t Trie) Insert(key string) {
	currentNode := t.root
	for _, letter := range key {
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
	for _, letter := range key {
		if val, ok := currentNode.children[letter]; ok {
			currentNode = val
		} else {
			return false
		}
	}
	return currentNode.endOfWord
}
