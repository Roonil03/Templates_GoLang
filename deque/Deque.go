package deque

type Deque[T any] struct {
	data       []T
	head, tail int
	size       int
}

func NewDeque[T any](cap int) *Deque[T] {
	return &Deque[T]{data: make([]T, cap), head: 0, tail: 0, size: 0}
}

func (d *Deque[T]) grow() {
	if d.size < len(d.data) {
		return
	}
	data := make([]T, 2*d.size+1)
	n := copy(data, d.data[d.head:])
	copy(data[n:], d.data[:d.head])
	d.data, d.head, d.tail = data, 0, d.size
}

func (d *Deque[T]) PushFront(val T) {
	d.grow()
	d.head = (d.head - 1 + len(d.data)) % len(d.data)
	d.data[d.head] = val
	d.size++
}

func (d *Deque[T]) PushBack(val T) {
	d.grow()
	d.data[d.tail] = val
	d.tail = (d.tail + 1) % len(d.data)
	d.size++
}

func (d *Deque[T]) PopFront() T {
	if d.size == 0 {
		panic("pop from empty deque")
	}
	val := d.data[d.head]
	var zero T
	d.data[d.head] = zero
	d.head = (d.head + 1) % len(d.data)
	d.size--
	return val
}

func (d *Deque[T]) PopBack() T {
	if d.size == 0 {
		panic("pop from empty deque")
	}
	d.tail = (d.tail - 1 + len(d.data)) % len(d.data)
	val := d.data[d.tail]
	var zero T
	d.data[d.tail] = zero
	d.size--
	return val
}
