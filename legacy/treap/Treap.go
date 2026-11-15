package treap

import "math/rand"

type Treap struct {
	key, pri, sz int
	left, right  *Treap
}

func NewTreap(key int) *Treap {
	return &Treap{key: key, pri: rand.Int(), sz: 1}
}

func (t *Treap) recalc() {
	t.sz = 1
	if t.left != nil {
		t.sz += t.left.sz
	}
	if t.right != nil {
		t.sz += t.right.sz
	}
}

func split(t *Treap, key int, includeEqual bool) (*Treap, *Treap) {
	if t == nil {
		return nil, nil
	}
	if key < t.key || (!includeEqual && key == t.key) {
		l, r := split(t.left, key, includeEqual)
		t.left = r
		t.recalc()
		return l, t
	}
	l, r := split(t.right, key, includeEqual)
	t.right = l
	t.recalc()
	return t, r
}

func merge(a, b *Treap) *Treap {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.pri < b.pri {
		a.right = merge(a.right, b)
		a.recalc()
		return a
	}
	b.left = merge(a, b.left)
	b.recalc()
	return b
}

func (t *Treap) Insert(key int) *Treap {
	n := NewTreap(key)
	a, b := split(t, key, true)
	return merge(merge(a, n), b)
}

func (t *Treap) Delete(key int) *Treap {
	a, b := split(t, key, false)
	_, c := split(b, key, true)
	return merge(a, c)
}

func (t *Treap) Contains(key int) bool {
	for t != nil {
		if key == t.key {
			return true
		}
		if key < t.key {
			t = t.left
		} else {
			t = t.right
		}
	}
	return false
}
