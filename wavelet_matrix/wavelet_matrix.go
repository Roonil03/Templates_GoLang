// Package wavelet_matrix provides static range order statistics in a wavelet matrix.
package wavelet_matrix

import (
	"math/bits"
	"sort"
)

// WaveletMatrix indexes an immutable sequence in O(n log(sigma+1)) space.
// Construction takes O(n log n) time; queries take O(log(sigma+1)) time,
// where sigma is the distinct value count. Values use cmp.Compare ordering,
// including its ordering of NaNs. The zero value is an empty matrix.
type WaveletMatrix[T ordered] struct {
	values []T
	prefix [][]int
	zeros  []int
	n      int
}

// NewWaveletMatrix copies data; subsequent changes to data have no effect.
func NewWaveletMatrix[T ordered](data []T) *WaveletMatrix[T] {
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
	w := &WaveletMatrix[T]{values: values, n: len(data)}
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
func (w *WaveletMatrix[T]) Len() int { return w.n }

func (w *WaveletMatrix[T]) check(l, r int) {
	if l < 0 || l > r || r > w.n {
		panic("wavelet_matrix: invalid range")
	}
}
func (w *WaveletMatrix[T]) lowerBound(value T) int {
	return sort.Search(len(w.values), func(i int) bool { return compare(w.values[i], value) >= 0 })
}

// Kth returns the zero-based kth smallest value in data[l:r].
// It returns false if k is outside [0,r-l). It panics unless 0 <= l <= r <= Len().
func (w *WaveletMatrix[T]) Kth(l, r, k int) (T, bool) {
	w.check(l, r)
	if k < 0 || k >= r-l {
		var zero T
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

func (w *WaveletMatrix[T]) countLessRank(l, r, rank int) int {
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
func (w *WaveletMatrix[T]) CountLess(l, r int, value T) int {
	w.check(l, r)
	return w.countLessRank(l, r, w.lowerBound(value))
}

// Count counts occurrences of value in data[l:r], using cmp.Compare equality.
// It panics unless 0 <= l <= r <= Len().
func (w *WaveletMatrix[T]) Count(l, r int, value T) int {
	w.check(l, r)
	rank := w.lowerBound(value)
	if rank == len(w.values) || compare(w.values[rank], value) != 0 {
		return 0
	}
	return w.countLessRank(l, r, rank+1) - w.countLessRank(l, r, rank)
}

// RangeFreq counts values in [low,high) within data[l:r]; low >= high returns 0.
// It panics unless 0 <= l <= r <= Len().
func (w *WaveletMatrix[T]) RangeFreq(l, r int, low, high T) int {
	w.check(l, r)
	if compare(low, high) >= 0 {
		return 0
	}
	return w.countLessRank(l, r, w.lowerBound(high)) - w.countLessRank(l, r, w.lowerBound(low))
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
