// Package wavelet_tree provides a static range order-statistics tree.
package wavelet_tree

import (
	"sort"
)

// WaveletTree indexes an immutable sequence. Construction takes O(n log n)
// time and O(n log(sigma+1)) space, where sigma is the distinct value count.
// Queries take O(log(sigma+1)) time. Values use cmp.Compare ordering,
// including its ordering of NaNs. The zero value is an empty tree.
type WaveletTree[T ordered] struct {
	values []T
	root   *node
	n      int
}

type node struct {
	lo, hi      int
	prefix      []int
	left, right *node
}

// NewWaveletTree copies data; later changes to data do not affect the tree.
func NewWaveletTree[T ordered](data []T) *WaveletTree[T] {
	values := append(data[:0:0], data...)
	sort.Slice(values, func(i, j int) bool { return compare(values[i], values[j]) < 0 })
	unique := 0
	for _, value := range values {
		if unique == 0 || compare(values[unique-1], value) != 0 {
			values[unique] = value
			unique++
		}
	}
	var zero T
	for i := unique; i < len(values); i++ {
		values[i] = zero
	}
	values = values[:unique]
	w := &WaveletTree[T]{values: values, n: len(data)}
	if len(data) == 0 {
		return w
	}
	ranks := make([]int, len(data))
	for i, v := range data {
		ranks[i] = w.lowerBound(v)
	}
	w.root = build(ranks, 0, len(values))
	return w
}

func build(data []int, lo, hi int) *node {
	p := &node{lo: lo, hi: hi}
	if hi-lo == 1 {
		return p
	}
	mid := lo + (hi-lo)/2
	p.prefix = make([]int, len(data)+1)
	left, right := make([]int, 0, len(data)), make([]int, 0, len(data))
	for i, v := range data {
		p.prefix[i+1] = p.prefix[i]
		if v < mid {
			left = append(left, v)
			p.prefix[i+1]++
		} else {
			right = append(right, v)
		}
	}
	p.left, p.right = build(left, lo, mid), build(right, mid, hi)
	return p
}

// Len returns the length of the original sequence in O(1) time.
func (w *WaveletTree[T]) Len() int { return w.n }

func (w *WaveletTree[T]) check(l, r int) {
	if l < 0 || l > r || r > w.n {
		panic("wavelet_tree: invalid range")
	}
}

func (w *WaveletTree[T]) lowerBound(value T) int {
	return sort.Search(len(w.values), func(i int) bool { return compare(w.values[i], value) >= 0 })
}

// Kth returns the zero-based kth smallest value in data[l:r].
// It returns false if k is outside [0,r-l). It panics unless 0 <= l <= r <= Len().
func (w *WaveletTree[T]) Kth(l, r, k int) (T, bool) {
	w.check(l, r)
	if k < 0 || k >= r-l {
		var zero T
		return zero, false
	}
	p := w.root
	for p.hi-p.lo > 1 {
		a, b := p.prefix[l], p.prefix[r]
		if k < b-a {
			l, r, p = a, b, p.left
		} else {
			k -= b - a
			l, r, p = l-a, r-b, p.right
		}
	}
	return w.values[p.lo], true
}

func countLess(p *node, l, r, rank int) int {
	if p == nil || l == r || rank <= p.lo {
		return 0
	}
	if rank >= p.hi {
		return r - l
	}
	a, b := p.prefix[l], p.prefix[r]
	return countLess(p.left, a, b, rank) + countLess(p.right, l-a, r-b, rank)
}

// CountLess counts values strictly below value in data[l:r].
// It panics unless 0 <= l <= r <= Len().
func (w *WaveletTree[T]) CountLess(l, r int, value T) int {
	w.check(l, r)
	return countLess(w.root, l, r, w.lowerBound(value))
}

// Count counts occurrences of value in data[l:r], using cmp.Compare equality.
// It panics unless 0 <= l <= r <= Len().
func (w *WaveletTree[T]) Count(l, r int, value T) int {
	w.check(l, r)
	rank := w.lowerBound(value)
	if rank == len(w.values) || compare(w.values[rank], value) != 0 {
		return 0
	}
	return countLess(w.root, l, r, rank+1) - countLess(w.root, l, r, rank)
}

// RangeFreq counts values in [low,high) within data[l:r]; low >= high returns 0.
// It panics unless 0 <= l <= r <= Len().
func (w *WaveletTree[T]) RangeFreq(l, r int, low, high T) int {
	w.check(l, r)
	if compare(low, high) >= 0 {
		return 0
	}
	return countLess(w.root, l, r, w.lowerBound(high)) - countLess(w.root, l, r, w.lowerBound(low))
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
