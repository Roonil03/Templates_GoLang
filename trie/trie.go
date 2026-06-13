package trie

type Node[T any] struct {
	next  [26]*Node[T]
	value T
	term  bool
}

type Trie[T any] struct {
	root *Node[T]
}

func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{root: &Node[T]{}}
}

func (t *Trie[T]) Insert(word string, val T) {
	curr := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if curr.next[c] == nil {
			curr.next[c] = &Node[T]{}
		}
		curr = curr.next[c]
	}
	curr.value = val
	curr.term = true
}

func (t *Trie[T]) Search(word string) (T, bool) {
	curr := t.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if curr.next[c] == nil {
			var zero T
			return zero, false
		}
		curr = curr.next[c]
	}
	return curr.value, curr.term
}
