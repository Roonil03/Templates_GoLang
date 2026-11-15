// Package wavelet_matrix provides static range order statistics in a wavelet matrix.
package wavelet_matrix

import (
	"math/bits"

	"sort"
)

// WaveletMatrix indexes an immutable sequence in O(n log(sigma+1)) space.
// Construction takes O(n log n) time; queries take O(log(sigma+1)) time,
// where sigma is the distinct value count. Values are integers. The zero value is an empty matrix.
type WaveletMatrix struct {
	values []int
	prefix [][]int
	zeros  []int
	n      int
}

// NewWaveletMatrix copies data; subsequent changes to data have no effect.
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

// Len returns the original sequence length in O(1) time.
func (w *WaveletMatrix) Len() int { return w.n }

func (w *WaveletMatrix) check(l, r int) {
	if l < 0 || l > r || r > w.n {
		panic("wavelet_matrix: invalid range")
	}
}
func (w *WaveletMatrix) lowerBound(value int) int {
	return sort.Search(len(w.values), func(i int) bool { return w.values[i] >= value })
}

// Kth returns the zero-based kth smallest value in data[l:r].
// It returns false if k is outside [0,r-l). It panics unless 0 <= l <= r <= Len().
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

// CountLess counts values strictly below value in data[l:r].
// It panics unless 0 <= l <= r <= Len().
func (w *WaveletMatrix) CountLess(l, r int, value int) int {
	w.check(l, r)
	return w.countLessRank(l, r, w.lowerBound(value))
}

// Count counts occurrences of value in data[l:r].
// It panics unless 0 <= l <= r <= Len().
func (w *WaveletMatrix) Count(l, r int, value int) int {
	w.check(l, r)
	rank := w.lowerBound(value)
	if rank == len(w.values) || w.values[rank] != value {
		return 0
	}
	return w.countLessRank(l, r, rank+1) - w.countLessRank(l, r, rank)
}

// RangeFreq counts values in [low,high) within data[l:r]; low >= high returns 0.
// It panics unless 0 <= l <= r <= Len().
func (w *WaveletMatrix) RangeFreq(l, r int, low, high int) int {
	w.check(l, r)
	if low >= high {
		return 0
	}
	return w.countLessRank(l, r, w.lowerBound(high)) - w.countLessRank(l, r, w.lowerBound(low))
}
