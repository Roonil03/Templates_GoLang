package ds

type Fenwick struct {
	n    int
	tree []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{n: n, tree: make([]int, n+1)}
}

func (f *Fenwick) Update(i, v int) {
	for ; i <= f.n; i += i & -i {
		f.tree[i] += v
	}
}

func (f *Fenwick) Query(i int) int {
	sum := 0
	for ; i > 0; i -= i & -i {
		sum += f.tree[i]
	}
	return sum
}
