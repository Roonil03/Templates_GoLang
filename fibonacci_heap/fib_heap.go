package fibonacci_heap

import "cmp"

type Node[T cmp.Ordered] struct {
	Key                 T
	parent, child       *Node[T]
	left, right         *Node[T]
	degree              int
	mark                bool
}

type FibonacciHeap[T cmp.Ordered] struct {
	minNode *Node[T]
	size    int
}

func NewFibonacciHeap[T cmp.Ordered]() *FibonacciHeap[T] {
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
