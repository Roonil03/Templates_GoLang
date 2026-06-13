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
}
