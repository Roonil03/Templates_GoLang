package red_black_tree

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type Node struct {
	Key                 int64
	left, right, parent *Node
	color               Color
}

type RedBlackTree struct {
	root *Node
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (t *RedBlackTree) leftRotate(x *Node) {
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

func (t *RedBlackTree) Insert(key int64) {
	node := &Node{Key: key, color: Red}
	var y *Node
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
	t.balanceInsert(node)
}
func (t *RedBlackTree) rightRotate(x *Node) {
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

func (t *RedBlackTree) balanceInsert(n *Node) {
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
func (t *RedBlackTree) Contains(key int64) bool {
	for n := t.root; n != nil; {
		if key == n.Key {
			return true
		}
		if key < n.Key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return false
}
