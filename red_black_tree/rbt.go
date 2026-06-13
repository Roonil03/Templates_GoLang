package red_black_tree

import "cmp"

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type Node[T cmp.Ordered] struct {
	Key                 T
	left, right, parent *Node[T]
	color               Color
}

type RedBlackTree[T cmp.Ordered] struct {
	root *Node[T]
}

func NewRedBlackTree[T cmp.Ordered]() *RedBlackTree[T] {
	return &RedBlackTree[T]{}
}

func (t *RedBlackTree[T]) leftRotate(x *Node[T]) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *RedBlackTree[T]) Insert(key T) {
	node := &Node[T]{Key: key, color: Red}
	var y *Node[T]
	x := t.root
	for x != nil {
		y = x
		if node.Key < x.Key {
			x = x.left
		} else {
			x = x.right
		}
	}
	node.parent = y
	if y == nil {
		t.root = node
	} else if node.Key < y.Key {
		y.left = node
	} else {
		y.right = node
	}
}
