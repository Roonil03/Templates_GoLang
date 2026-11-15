// Package monotonic_stack provides a nondecreasing stack and next-element queries.
package monotonic_stack

import "cmp"

// MonotonicStack keeps values in nondecreasing order from bottom to top.
// Push discards larger values; equal values are retained. Values use cmp.Compare
// ordering, including NaNs. The zero value is ready to use. Storage is O(n).
type MonotonicStack[T cmp.Ordered] struct{ data []T }

// NewMonotonicStack returns an empty stack.
func NewMonotonicStack[T cmp.Ordered]() *MonotonicStack[T] {
	return &MonotonicStack[T]{}
}

// Push removes larger values from the top, then appends value.
// Across a sequence of pushes and pops, each push takes amortized O(1) time.
func (s *MonotonicStack[T]) Push(value T) {
	for len(s.data) > 0 && cmp.Compare(s.data[len(s.data)-1], value) > 0 {
		var zero T
		s.data[len(s.data)-1] = zero
		s.data = s.data[:len(s.data)-1]
	}
	s.data = append(s.data, value)
}

// Top returns the top value and whether one exists, in O(1) time.
func (s *MonotonicStack[T]) Top() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// Pop removes and returns the top value in O(1) time; false means empty.
func (s *MonotonicStack[T]) Pop() (T, bool) {
	value, ok := s.Top()
	if ok {
		var zero T
		s.data[len(s.data)-1] = zero
		s.data = s.data[:len(s.data)-1]
	}
	return value, ok
}

// Len returns the number of retained values in O(1) time.
func (s *MonotonicStack[T]) Len() int { return len(s.data) }

// NextGreaterIndices returns, for each position, the index of the first strictly
// greater value to its right, or -1 if none exists. It takes O(n) time and space.
func NextGreaterIndices[T cmp.Ordered](data []T) []int { return nextIndices(data, true) }

// NextSmallerIndices returns, for each position, the index of the first strictly
// smaller value to its right, or -1 if none exists. It takes O(n) time and space.
func NextSmallerIndices[T cmp.Ordered](data []T) []int { return nextIndices(data, false) }

func nextIndices[T cmp.Ordered](data []T, greater bool) []int {
	answer := make([]int, len(data))
	stack := make([]int, 0, len(data))
	for i, value := range data {
		answer[i] = -1
		for len(stack) > 0 {
			j := stack[len(stack)-1]
			order := cmp.Compare(value, data[j])
			if greater && order <= 0 || !greater && order >= 0 {
				break
			}
			answer[j] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answer
}
