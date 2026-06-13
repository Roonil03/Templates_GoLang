package pbds

import "math/rand"

type Node struct {
	key         int
	pri, sz     int
	left, right *Node
}

type PBDS struct {
	root *Node
}

func NewPBDS() *PBDS {
	return &PBDS{}
}

func sz(n *Node) int {
	if n == nil {
		return 0
	}
	return n.sz
}

func (n *Node) upd() {
	n.sz = 1 + sz(n.left) + sz(n.right)
}

func split(n *Node, k int) (*Node, *Node) {
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

func merge(l, r *Node) *Node {
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

func (p *PBDS) Insert(k int) {
	l, r := split(p.root, k)
	p.root = merge(merge(l, &Node{key: k, pri: rand.Int(), sz: 1}), r)
}

func (p *PBDS) OrderOfKey(k int) int {
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

func (p *PBDS) FindByOrder(idx int) (int, bool) {
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
	return 0, false
}
