// Package heap provides binary min heaps and priority queues.
package heap

// PriorityQueue is a binary min heap. The zero value is ready to use.
// Values use cmp.Compare ordering, including its ordering of NaNs.
// Push and Pop take O(log n) time; Len and Peek take O(1) time.
// Storage is O(n).
type PriorityQueue[T ordered] struct{ data []T }

// NewPriorityQueue returns an empty min-priority queue.
func NewPriorityQueue[T ordered]() *PriorityQueue[T] { return &PriorityQueue[T]{} }

// MinHeap explicitly names a min-priority queue and exposes its operations.
// Its zero value is ready to use.
type MinHeap[T ordered] struct{ PriorityQueue[T] }

// NewMinHeap copies values and heapifies them in O(n) time.
func NewMinHeap[T ordered](values ...T) *MinHeap[T] {
	h := &MinHeap[T]{PriorityQueue: PriorityQueue[T]{data: append([]T(nil), values...)}}
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.down(i, len(h.data))
	}
	return h
}

// Push adds value to the heap.
func (h *PriorityQueue[T]) Push(value T) {
	h.data = append(h.data, value)
	h.up(len(h.data) - 1)
}

// Pop removes and returns the smallest value. It panics if the heap is empty.
func (h *PriorityQueue[T]) Pop() T {
	if len(h.data) == 0 {
		panic("heap: Pop on empty heap")
	}
	n := len(h.data) - 1
	value := h.data[0]
	h.data[0] = h.data[n]
	var zero T
	h.data[n] = zero
	h.data = h.data[:n]
	h.down(0, n)
	return value
}

// TryPop removes and returns the smallest value; false means the heap is empty.
func (h *PriorityQueue[T]) TryPop() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}
	return h.Pop(), true
}

// Peek returns the smallest value without removing it; false means empty.
func (h *PriorityQueue[T]) Peek() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}
	return h.data[0], true
}

// Len returns the number of stored values.
func (h *PriorityQueue[T]) Len() int { return len(h.data) }

func (h *PriorityQueue[T]) up(j int) {
	for j > 0 {
		i := (j - 1) / 2
		if compare(h.data[j], h.data[i]) >= 0 {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		j = i
	}
}

func (h *PriorityQueue[T]) down(i, n int) {
	for i < n/2 {
		j := 2*i + 1
		if j+1 < n && compare(h.data[j+1], h.data[j]) < 0 {
			j++
		}
		if compare(h.data[j], h.data[i]) >= 0 {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}
	// For odd n, i==n/2 has no children; for even n, the loop includes
	// the final parent with its single left child.
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
