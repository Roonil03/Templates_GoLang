package trie

type Node struct {
	next [26]*Node
	term bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if curr.next[c] == nil {
			curr.next[c] = &Node{}
		}
		curr = curr.next[c]
	}
	curr.term = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if curr.next[c] == nil {
			return false
		}
		curr = curr.next[c]
	}
	return curr.term
}
