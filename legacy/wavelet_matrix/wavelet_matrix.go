package wavelet_matrix

import (
	"math/bits"

	"sort"
)

type WaveletMatrix struct {
	values []int
	prefix [][]int
	zeros  []int
	n      int
}

func NewWaveletMatrix(data []int) *WaveletMatrix {
	values := append([]int(nil), data...)
	sort.Ints(values)
	unique := values[:0]
	for _, v := range values {
		if len(unique) == 0 || unique[len(unique)-1] != v {
			unique = append(unique, v)
		}
	}
	values = unique
	w := &WaveletMatrix{values: values, n: len(data)}
	if len(values) < 2 {
		return w
	}
	height := bits.Len(uint(len(values) - 1))
	w.prefix = make([][]int, height)
	w.zeros = make([]int, height)
	current, next := make([]int, len(data)), make([]int, len(data))
	for i, v := range data {
		current[i] = w.lowerBound(v)
	}
	for level := 0; level < height; level++ {
		bit := height - 1 - level
		prefix := make([]int, len(data)+1)
		for i, v := range current {
			prefix[i+1] = prefix[i] + ((v >> bit) & 1)
		}
		zeros := len(data) - prefix[len(data)]
		w.prefix[level], w.zeros[level] = prefix, zeros
		a, b := 0, zeros
		for _, v := range current {
			if (v>>bit)&1 == 0 {
				next[a] = v
				a++
			} else {
				next[b] = v
				b++
			}
		}
		current, next = next, current
	}
	return w
}

func (w *WaveletMatrix) Len() int { return w.n }

func (w *WaveletMatrix) check(l, r int) {
	if l < 0 || l > r || r > w.n {
		panic("wavelet_matrix: invalid range")
	}
}
func (w *WaveletMatrix) lowerBound(value int) int {
	return sort.Search(len(w.values), func(i int) bool { return w.values[i] >= value })
}

func (w *WaveletMatrix) Kth(l, r, k int) (int, bool) {
	w.check(l, r)
	if k < 0 || k >= r-l {
		var zero int
		return zero, false
	}
	rank := 0
	for level, prefix := range w.prefix {
		a, b := prefix[l], prefix[r]
		zeros := r - l - (b - a)
		if k < zeros {
			l, r = l-a, r-b
		} else {
			k -= zeros
			rank |= 1 << uint(len(w.prefix)-1-level)
			l, r = w.zeros[level]+a, w.zeros[level]+b
		}
	}
	return w.values[rank], true
}

func (w *WaveletMatrix) countLessRank(l, r, rank int) int {
	if rank <= 0 {
		return 0
	}
	if rank >= len(w.values) {
		return r - l
	}
	count := 0
	for level, prefix := range w.prefix {
		a, b := prefix[l], prefix[r]
		if (rank>>uint(len(w.prefix)-1-level))&1 == 0 {
			l, r = l-a, r-b
		} else {
			count += r - l - (b - a)
			l, r = w.zeros[level]+a, w.zeros[level]+b
		}
	}
	return count
}

func (w *WaveletMatrix) CountLess(l, r int, value int) int {
	w.check(l, r)
	return w.countLessRank(l, r, w.lowerBound(value))
}

func (w *WaveletMatrix) Count(l, r int, value int) int {
	w.check(l, r)
	rank := w.lowerBound(value)
	if rank == len(w.values) || w.values[rank] != value {
		return 0
	}
	return w.countLessRank(l, r, rank+1) - w.countLessRank(l, r, rank)
}

func (w *WaveletMatrix) RangeFreq(l, r int, low, high int) int {
	w.check(l, r)
	if low >= high {
		return 0
	}
	return w.countLessRank(l, r, w.lowerBound(high)) - w.countLessRank(l, r, w.lowerBound(low))
}
