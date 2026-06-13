package fibonacci_heap

type Node struct {
	Key                 int64
	parent, child       *Node
	left, right         *Node
	degree              int
	mark                bool
}

type FibonacciHeap struct {
	minNode *Node
	size    int
}

func NewFibonacciHeap() *FibonacciHeap {
	return &FibonacciHeap{}
}

func (h *FibonacciHeap) Insert(key int64) *Node {
	node := &Node{Key: key}
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

func (h *FibonacciHeap) GetMin() *Node {
	return h.minNode
}
