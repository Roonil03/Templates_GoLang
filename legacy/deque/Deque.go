package deque

type Deque struct {
	data       []int
	head, tail int
	size       int
}

func NewDeque(cap int) *Deque {
	return &Deque{data: make([]int, cap), head: 0, tail: 0, size: 0}
}

func (d *Deque) grow() {
	if d.size < len(d.data) {
		return
	}
	data := make([]int, 2*d.size+1)
	n := copy(data, d.data[d.head:])
	copy(data[n:], d.data[:d.head])
	d.data, d.head, d.tail = data, 0, d.size
}

func (d *Deque) PushFront(val int) {
	d.grow()
	d.head = (d.head - 1 + len(d.data)) % len(d.data)
	d.data[d.head] = val
	d.size++
}

func (d *Deque) PushBack(val int) {
	d.grow()
	d.data[d.tail] = val
	d.tail = (d.tail + 1) % len(d.data)
	d.size++
}

func (d *Deque) PopFront() int {
	if d.size == 0 {
		panic("pop from empty deque")
	}
	val := d.data[d.head]
	var zero int
	d.data[d.head] = zero
	d.head = (d.head + 1) % len(d.data)
	d.size--
	return val
}

func (d *Deque) PopBack() int {
	if d.size == 0 {
		panic("pop from empty deque")
	}
	d.tail = (d.tail - 1 + len(d.data)) % len(d.data)
	val := d.data[d.tail]
	var zero int
	d.data[d.tail] = zero
	d.size--
	return val
}
