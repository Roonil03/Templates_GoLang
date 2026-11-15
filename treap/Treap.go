package treap

import (
	"math/rand"
)

type Treap[K ordered] struct {
	key         K
	pri, sz     int
	left, right *Treap[K]
}

func NewTreap[K ordered](key K) *Treap[K] {
	return &Treap[K]{key: key, pri: rand.Int(), sz: 1}
}

func (t *Treap[K]) recalc() {
	t.sz = 1
	if t.left != nil {
		t.sz += t.left.sz
	}
	if t.right != nil {
		t.sz += t.right.sz
	}
}

func splitStrict[K ordered](t *Treap[K], key K) (*Treap[K], *Treap[K]) {
	if t == nil {
		return nil, nil
	}
	if compare(t.key, key) < 0 {
		l, r := splitStrict(t.right, key)
		t.right = l
		t.recalc()
		return t, r
	}
	l, r := splitStrict(t.left, key)
	t.left = r
	t.recalc()
	return l, t
}

func splitEq[K ordered](t *Treap[K], key K) (*Treap[K], *Treap[K]) {
	if t == nil {
		return nil, nil
	}
	if compare(key, t.key) < 0 {
		l, r := splitEq(t.left, key)
		t.left = r
		t.recalc()
		return l, t
	}
	l, r := splitEq(t.right, key)
	t.right = l
	t.recalc()
	return t, r
}

func merge[K ordered](a, b *Treap[K]) *Treap[K] {
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

func (t *Treap[K]) Insert(key K) *Treap[K] {
	n := NewTreap(key)
	a, b := splitEq(t, key)
	return merge(merge(a, n), b)
}

func (t *Treap[K]) Delete(key K) *Treap[K] {
	a, b := splitStrict(t, key)
	b1, c := splitEq(b, key)
	_ = b1
	return merge(a, c)
}

func (t *Treap[K]) Contains(key K) bool {
	for t != nil {
		order := compare(key, t.key)
		if order == 0 {
			return true
		}
		if order < 0 {
			t = t.left
		} else {
			t = t.right
		}
	}
	return false
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

func compare[T ordered](x, y T) int {
	if x < y || (x != x && y == y) {
		return -1
	}
	if x > y || (x == x && y != y) {
		return 1
	}
	return 0
}
