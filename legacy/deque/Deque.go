package deque

type Deque struct {
	data       []int
	head, tail int
	size       int
}

func NewDeque(cap int) *Deque {
	return &Deque{data: make([]int, cap), head: 0, tail: 0, size: 0}
}

func (d *Deque) PushFront(val int) {
	d.head = (d.head - 1 + len(d.data)) % len(d.data)
	d.data[d.head] = val
	d.size++
}

func (d *Deque) PushBack(val int) {
	d.data[d.tail] = val
	d.tail = (d.tail + 1) % len(d.data)
	d.size++
}

func (d *Deque) PopFront() int {
	val := d.data[d.head]
	d.head = (d.head + 1) % len(d.data)
	d.size--
	return val
}

func (d *Deque) PopBack() int {
	d.tail = (d.tail - 1 + len(d.data)) % len(d.data)
	val := d.data[d.tail]
	d.size--
	return val
}
