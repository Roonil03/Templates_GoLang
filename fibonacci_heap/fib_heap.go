package fibonacci_heap

type Node[T ordered] struct {
	Key                 T
	parent, child       *Node[T]
	left, right         *Node[T]
	degree              int
	mark                bool
}

type FibonacciHeap[T ordered] struct {
	minNode *Node[T]
	size    int
}

func NewFibonacciHeap[T ordered]() *FibonacciHeap[T] {
	return &FibonacciHeap[T]{}
}

func (h *FibonacciHeap[T]) Insert(key T) *Node[T] {
	node := &Node[T]{Key: key}
	node.left = node
	node.right = node
	if h.minNode == nil {
		h.minNode = node
	} else {
		node.right = h.minNode.right
		node.left = h.minNode
		h.minNode.right.left = node
		h.minNode.right = node
		if key < h.minNode.Key {
			h.minNode = node
		}
	}
	h.size++
	return node
}

func (h *FibonacciHeap[T]) GetMin() *Node[T] {
	return h.minNode
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}
