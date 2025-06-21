package ds

type Fenwick2D struct {
	n, m int
	tree [][]int
}

func NewFenwick2D(n, m int) *Fenwick2D {
	b := make([][]int, n+1)
	for i := range b {
		b[i] = make([]int, m+1)
	}
	return &Fenwick2D{n: n, m: m, tree: b}
}

func (f *Fenwick2D) Update(x, y, v int) {
	for i := x; i <= f.n; i += i & -i {
		for j := y; j <= f.m; j += j & -j {
			f.tree[i][j] += v
		}
	}
}

func (f *Fenwick2D) Query(x, y int) int {
	sum := 0
	for i := x; i > 0; i -= i & -i {
		for j := y; j > 0; j -= j & -j {
			sum += f.tree[i][j]
		}
	}
	return sum
}
