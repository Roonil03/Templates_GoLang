package heap

import "cmp"

type PriorityQueue[T cmp.Ordered] struct {
	data []T
}

func NewPriorityQueue[T cmp.Ordered]() *PriorityQueue[T] {
	return &PriorityQueue[T]{data: make([]T, 0)}
}

func (h *PriorityQueue[T]) Push(x T) {
	h.data = append(h.data, x)
	h.up(len(h.data) - 1)
}

func (h *PriorityQueue[T]) Pop() T {
	n := len(h.data) - 1
	h.data[0], h.data[n] = h.data[n], h.data[0]
	x := h.data[n]
	h.data = h.data[0:n]
	h.down(0, n)
	return x
}

func (h *PriorityQueue[T]) Len() int {
	return len(h.data)
}

func (h *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2 
		if i == j || !(h.data[j] < h.data[i]) { 
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		j = i
	}
}

func (h *PriorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { 
			break
		}
		j := j1 
		if j2 := j1 + 1; j2 < n && h.data[j2] < h.data[j1] {
			j = j2 
		}
		if !(h.data[j] < h.data[i]) {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}
	return i > i0
}
