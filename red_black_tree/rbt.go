package red_black_tree

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type Node[T ordered] struct {
	Key                 T
	left, right, parent *Node[T]
	color               Color
}

// RedBlackTree orders keys using cmp.Compare, including its ordering of NaNs.
type RedBlackTree[T ordered] struct {
	root *Node[T]
}

func NewRedBlackTree[T ordered]() *RedBlackTree[T] {
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
		if compare(node.Key, x.Key) < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}
	node.parent = y
	if y == nil {
		t.root = node
	} else if compare(node.Key, y.Key) < 0 {
		y.left = node
	} else {
		y.right = node
	}
	t.balanceInsert(node)
}
func (t *RedBlackTree[T]) rightRotate(x *Node[T]) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

func (t *RedBlackTree[T]) balanceInsert(n *Node[T]) {
	for n.parent != nil && n.parent.color == Red {
		p := n.parent
		g := p.parent
		left := p == g.left
		uncle := g.left
		if left {
			uncle = g.right
		}
		if uncle != nil && uncle.color == Red {
			p.color, uncle.color, g.color = Black, Black, Red
			n = g
			continue
		}
		if left {
			if n == p.right {
				t.leftRotate(p)
				p = n
			}
			t.rightRotate(g)
		} else {
			if n == p.left {
				t.rightRotate(p)
				p = n
			}
			t.leftRotate(g)
		}
		p.color, g.color = Black, Red
	}
	t.root.color = Black
}

// Contains reports whether key has been inserted.
func (t *RedBlackTree[T]) Contains(key T) bool {
	for n := t.root; n != nil; {
		order := compare(key, n.Key)
		if order == 0 {
			return true
		}
		if order < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}
	return false
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

func compare[T ordered](x, y T) int {
	if x < y || (x != x && y == y) {
		return -1
	}
	if x > y || (x == x && y != y) {
		return 1
	}
	return 0
}
