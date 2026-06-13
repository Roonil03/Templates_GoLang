package persistent_segment_tree

import "cmp"

type Node[T cmp.Ordered] struct {
	left, right *Node[T]
	val         T
}

type PersistentSegmentTree[T cmp.Ordered] struct {
	roots []*Node[T]
	size  int
	agg   func(T, T) T
}

func NewPersistentSegmentTree[T cmp.Ordered](size int, agg func(T, T) T) *PersistentSegmentTree[T] {
	return &PersistentSegmentTree[T]{
		roots: make([]*Node[T], 0),
		size:  size,
		agg:   agg,
	}
}

func (pst *PersistentSegmentTree[T]) AddRoot(root *Node[T]) {
	pst.roots = append(pst.roots, root)
}

func (pst *PersistentSegmentTree[T]) Update(node *Node[T], l, r, idx int, val T) *Node[T] {
	if l == r {
		return &Node[T]{val: val}
	}
	mid := l + (r-l)/2
	newNode := &Node[T]{}
	if node != nil {
		newNode.left = node.left
		newNode.right = node.right
	}
	if idx <= mid {
		var ln *Node[T]
		if node != nil {
			ln = node.left
		}
		newNode.left = pst.Update(ln, l, mid, idx, val)
	} else {
		var rn *Node[T]
		if node != nil {
			rn = node.right
		}
		newNode.right = pst.Update(rn, mid+1, r, idx, val)
	}
	
	var lv, rv T
	if newNode.left != nil { lv = newNode.left.val }
	if newNode.right != nil { rv = newNode.right.val }
	
	if newNode.left != nil && newNode.right != nil {
		newNode.val = pst.agg(lv, rv)
	} else if newNode.left != nil {
		newNode.val = lv
	} else if newNode.right != nil {
		newNode.val = rv
	}
	
	return newNode
}
