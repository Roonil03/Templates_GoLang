package wavelet_tree

import (
	"sort"
)

type WaveletTree struct {
	values []int
	root   *node
	n      int
}

type node struct {
	lo, hi      int
	prefix      []int
	left, right *node
}

func NewWaveletTree(data []int) *WaveletTree {
	values := append([]int(nil), data...)
	sort.Ints(values)
	unique := values[:0]
	for _, v := range values {
		if len(unique) == 0 || unique[len(unique)-1] != v {
			unique = append(unique, v)
		}
	}
	values = unique
	w := &WaveletTree{values: values, n: len(data)}
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

func (w *WaveletTree) Len() int { return w.n }

func (w *WaveletTree) check(l, r int) {
	if l < 0 || l > r || r > w.n {
		panic("wavelet_tree: invalid range")
	}
}

func (w *WaveletTree) lowerBound(value int) int {
	return sort.Search(len(w.values), func(i int) bool { return w.values[i] >= value })
}

func (w *WaveletTree) Kth(l, r, k int) (int, bool) {
	w.check(l, r)
	if k < 0 || k >= r-l {
		var zero int
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

func (w *WaveletTree) CountLess(l, r int, value int) int {
	w.check(l, r)
	return countLess(w.root, l, r, w.lowerBound(value))
}

func (w *WaveletTree) Count(l, r int, value int) int {
	w.check(l, r)
	rank := w.lowerBound(value)
	if rank == len(w.values) || w.values[rank] != value {
		return 0
	}
	return countLess(w.root, l, r, rank+1) - countLess(w.root, l, r, rank)
}

func (w *WaveletTree) RangeFreq(l, r int, low, high int) int {
	w.check(l, r)
	if low >= high {
		return 0
	}
	return countLess(w.root, l, r, w.lowerBound(high)) - countLess(w.root, l, r, w.lowerBound(low))
}
