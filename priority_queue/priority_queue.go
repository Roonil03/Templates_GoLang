package priority_queue

type PriorityQueue[T ordered] struct{ data []T }

func NewPriorityQueue[T ordered](values ...T) *PriorityQueue[T] {
	h := &PriorityQueue[T]{data: append([]T(nil), values...)}
	for i := len(h.data)/2 - 1; i >= 0; i-- {
		h.down(i, len(h.data))
	}
	return h
}

func (h *PriorityQueue[T]) Push(value T) {
	h.data = append(h.data, value)
	h.up(len(h.data) - 1)
}

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

func (h *PriorityQueue[T]) TryPop() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}
	return h.Pop(), true
}

func (h *PriorityQueue[T]) Peek() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}
	return h.data[0], true
}

func (h *PriorityQueue[T]) Len() int { return len(h.data) }

func (h *PriorityQueue[T]) up(j int) {
	for j > 0 {
		i := (j - 1) / 2
		if compare(h.data[j], h.data[i]) <= 0 {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		j = i
	}
}

func (h *PriorityQueue[T]) down(i, n int) {
	for i < n/2 {
		j := 2*i + 1
		if j+1 < n && compare(h.data[j+1], h.data[j]) > 0 {
			j++
		}
		if compare(h.data[j], h.data[i]) <= 0 {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}
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
