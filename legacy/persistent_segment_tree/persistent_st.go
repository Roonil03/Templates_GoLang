package persistent_segment_tree

type Node struct {
	left, right *Node
	val         int64
}

type PersistentSegmentTree struct {
	roots []*Node
	size  int
	agg   func(int64, int64) int64
}

func NewPersistentSegmentTree(size int, agg func(int64, int64) int64) *PersistentSegmentTree {
	return &PersistentSegmentTree{
		roots: make([]*Node, 0),
		size:  size,
		agg:   agg,
	}
}

func (pst *PersistentSegmentTree) AddRoot(root *Node) {
	pst.roots = append(pst.roots, root)
}

func (pst *PersistentSegmentTree) Update(node *Node, l, r, idx int, val int64) *Node {
	if l == r {
		return &Node{val: val}
	}
	mid := l + (r-l)/2
	newNode := &Node{}
	if node != nil {
		newNode.left = node.left
		newNode.right = node.right
	}
	if idx <= mid {
		var ln *Node
		if node != nil {
			ln = node.left
		}
		newNode.left = pst.Update(ln, l, mid, idx, val)
	} else {
		var rn *Node
		if node != nil {
			rn = node.right
		}
		newNode.right = pst.Update(rn, mid+1, r, idx, val)
	}
	var lv, rv int64
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

// Root returns a previously saved root by its zero-based version number.
func (pst *PersistentSegmentTree) Root(version int) *Node { return pst.roots[version] }

// Query aggregates assigned leaves in the inclusive range [l,r]. Empty ranges return false.
func (pst *PersistentSegmentTree) Query(root *Node, l, r int) (int64, bool) {
	return pst.query(root, 0, pst.size-1, l, r)
}

func (pst *PersistentSegmentTree) query(node *Node, lo, hi, l, r int) (int64, bool) {
	if node == nil || hi < l || lo > r || lo > hi || l > r {
		var zero int64
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
