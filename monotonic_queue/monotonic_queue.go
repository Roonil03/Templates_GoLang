// Package monotonic_queue provides a FIFO queue with constant-time minimum queries.
package monotonic_queue

import "cmp"

// MonotonicQueue retains a nondecreasing deque of minimum candidates alongside
// its FIFO values. Equal values are retained. Values use cmp.Compare ordering,
// including NaNs. The zero value is ready to use. Space is O(n) in queue length;
// Push and Pop take amortized O(1) time, and Min and Len take O(1) time.
type MonotonicQueue[T cmp.Ordered] struct {
	values []T
	head   int
	minima []entry[T]
	first  int
}
type entry[T cmp.Ordered] struct {
	value T
	index int
}

// NewMonotonicQueue returns an empty queue.
func NewMonotonicQueue[T cmp.Ordered]() *MonotonicQueue[T] { return &MonotonicQueue[T]{} }

// Push appends value to the FIFO queue and discards larger minimum candidates.
func (q *MonotonicQueue[T]) Push(value T) {
	index := len(q.values)
	q.values = append(q.values, value)
	for len(q.minima) > q.first && cmp.Compare(q.minima[len(q.minima)-1].value, value) > 0 {
		q.minima[len(q.minima)-1] = entry[T]{}
		q.minima = q.minima[:len(q.minima)-1]
	}
	q.minima = append(q.minima, entry[T]{value: value, index: index})
}

// Pop removes the oldest FIFO value; false means the queue is empty.
func (q *MonotonicQueue[T]) Pop() (T, bool) {
	if q.Len() == 0 {
		var zero T
		return zero, false
	}
	value := q.values[q.head]
	var zero T
	q.values[q.head] = zero
	if q.minima[q.first].index == q.head {
		q.minima[q.first] = entry[T]{}
		q.first++
	}
	q.head++
	if q.head == len(q.values) {
		q.values = nil
		q.minima = nil
		q.head = 0
		q.first = 0
	} else if q.head >= len(q.values)-q.head {
		q.values = append([]T(nil), q.values[q.head:]...)
		q.minima = append([]entry[T](nil), q.minima[q.first:]...)
		for i := range q.minima {
			q.minima[i].index -= q.head
		}
		q.head = 0
		q.first = 0
	} else if q.first > 0 && q.first >= len(q.minima)-q.first {
		q.minima = append([]entry[T](nil), q.minima[q.first:]...)
		q.first = 0
	}
	return value, true
}

// Min returns the smallest queued value; false means the queue is empty.
func (q *MonotonicQueue[T]) Min() (T, bool) {
	if q.Len() == 0 {
		var zero T
		return zero, false
	}
	return q.minima[q.first].value, true
}

// Len returns the number of queued FIFO values.
func (q *MonotonicQueue[T]) Len() int { return len(q.values) - q.head }

// SlidingWindowMin returns the minimum in every consecutive window of size k.
// It takes O(n) time and O(k) auxiliary space. It panics unless 1 <= k <= len(data).
func SlidingWindowMin[T cmp.Ordered](data []T, k int) []T {
	if k < 1 || k > len(data) {
		panic("monotonic_queue: invalid window size")
	}
	q := NewMonotonicQueue[T]()
	answer := make([]T, 0, len(data)-k+1)
	for i, value := range data {
		if q.Len() == k {
			q.Pop()
		}
		q.Push(value)
		if i >= k-1 {
			minimum, _ := q.Min()
			answer = append(answer, minimum)
		}
	}
	return answer
}
