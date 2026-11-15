package persistent_segment_tree

type Node[T ordered] struct {
	left, right *Node[T]
	val         T
}

type PersistentSegmentTree[T ordered] struct {
	roots []*Node[T]
	size  int
	agg   func(T, T) T
}

func NewPersistentSegmentTree[T ordered](size int, agg func(T, T) T) *PersistentSegmentTree[T] {
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
	if newNode.left != nil {
		lv = newNode.left.val
	}
	if newNode.right != nil {
		rv = newNode.right.val
	}

	if newNode.left != nil && newNode.right != nil {
		newNode.val = pst.agg(lv, rv)
	} else if newNode.left != nil {
		newNode.val = lv
	} else if newNode.right != nil {
		newNode.val = rv
	}

	return newNode
}

func (pst *PersistentSegmentTree[T]) Root(version int) *Node[T] { return pst.roots[version] }

func (pst *PersistentSegmentTree[T]) Query(root *Node[T], l, r int) (T, bool) {
	return pst.query(root, 0, pst.size-1, l, r)
}

func (pst *PersistentSegmentTree[T]) query(node *Node[T], lo, hi, l, r int) (T, bool) {
	if node == nil || hi < l || lo > r || lo > hi || l > r {
		var zero T
		return zero, false
	}
	if l <= lo && hi <= r {
		return node.val, true
	}
	mid := lo + (hi-lo)/2
	a, okA := pst.query(node.left, lo, mid, l, r)
	b, okB := pst.query(node.right, mid+1, hi, l, r)
	if !okA {
		return b, okB
	}
	if !okB {
		return a, true
	}
	return pst.agg(a, b), true
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}
