package pbds

import (
	"cmp"
	"math/rand"
)

type Node[T cmp.Ordered] struct {
	key         T
	pri, sz     int
	left, right *Node[T]
}

type PBDS[T cmp.Ordered] struct {
	root *Node[T]
}

func NewPBDS[T cmp.Ordered]() *PBDS[T] {
	return &PBDS[T]{}
}

func sz[T cmp.Ordered](n *Node[T]) int {
	if n == nil {
		return 0
	}
	return n.sz
}

func (n *Node[T]) upd() {
	n.sz = 1 + sz(n.left) + sz(n.right)
}

func split[T cmp.Ordered](n *Node[T], k T) (*Node[T], *Node[T]) {
	if n == nil {
		return nil, nil
	}
	if n.key < k {
		l, r := split(n.right, k)
		n.right = l
		n.upd()
		return n, r
	}
	l, r := split(n.left, k)
	n.left = r
	n.upd()
	return l, n
}

func merge[T cmp.Ordered](l, r *Node[T]) *Node[T] {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.pri > r.pri {
		l.right = merge(l.right, r)
		l.upd()
		return l
	}
	r.left = merge(l, r.left)
	r.upd()
	return r
}

func (p *PBDS[T]) Insert(k T) {
	l, r := split(p.root, k)
	p.root = merge(merge(l, &Node[T]{key: k, pri: rand.Int(), sz: 1}), r)
}

func (p *PBDS[T]) OrderOfKey(k T) int {
	curr := p.root
	ans := 0
	for curr != nil {
		if curr.key < k {
			ans += sz(curr.left) + 1
			curr = curr.right
		} else {
			curr = curr.left
		}
	}
	return ans
}

func (p *PBDS[T]) FindByOrder(idx int) (T, bool) {
	curr := p.root
	for curr != nil {
		lsz := sz(curr.left)
		if lsz == idx {
			return curr.key, true
		}
		if idx < lsz {
			curr = curr.left
		} else {
			idx -= lsz + 1
			curr = curr.right
		}
	}
	var zero T
	return zero, false
}
