package min_heap

type PriorityQueue struct{ data []int }

func NewPriorityQueue() *PriorityQueue { return &PriorityQueue{} }

type MinHeap struct{ PriorityQueue }

func NewMinHeap(values ...int) *MinHeap {
	h := &MinHeap{PriorityQueue: PriorityQueue{data: append([]int(nil), values...)}}
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.down(i, len(h.data))
	}
	return h
}

func (h *PriorityQueue) Push(value int) {
	h.data = append(h.data, value)
	h.up(len(h.data) - 1)
}

func (h *PriorityQueue) Pop() int {
	if len(h.data) == 0 {
		panic("heap: Pop on empty heap")
	}
	n := len(h.data) - 1
	value := h.data[0]
	h.data[0] = h.data[n]
	var zero int
	h.data[n] = zero
	h.data = h.data[:n]
	h.down(0, n)
	return value
}

func (h *PriorityQueue) TryPop() (int, bool) {
	if h.Len() == 0 {
		var zero int
		return zero, false
	}
	return h.Pop(), true
}

func (h *PriorityQueue) Peek() (int, bool) {
	if h.Len() == 0 {
		var zero int
		return zero, false
	}
	return h.data[0], true
}

func (h *PriorityQueue) Len() int { return len(h.data) }

func (h *PriorityQueue) up(j int) {
	for j > 0 {
		i := (j - 1) / 2
		if compare(h.data[j], h.data[i]) >= 0 {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		j = i
	}
}

func (h *PriorityQueue) down(i, n int) {
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
}

func compare(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
