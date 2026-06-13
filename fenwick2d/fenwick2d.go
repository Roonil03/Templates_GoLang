package fenwick2d

import "cmp"

type Fenwick2D[T cmp.Ordered] struct {
	n, m int
	tree [][]T
	add  func(T, T) T
}

func NewFenwick2D[T cmp.Ordered](n, m int, add func(T, T) T) *Fenwick2D[T] {
	b := make([][]T, n+1)
	for i := range b {
		b[i] = make([]T, m+1)
	}
	return &Fenwick2D[T]{n: n, m: m, tree: b, add: add}
}

func (f *Fenwick2D[T]) Update(x, y int, v T) {
	for i := x; i <= f.n; i += i & -i {
		for j := y; j <= f.m; j += j & -j {
			f.tree[i][j] = f.add(f.tree[i][j], v)
		}
	}
}

func (f *Fenwick2D[T]) Query(x, y int, zero T) T {
	sum := zero
	for i := x; i > 0; i -= i & -i {
		for j := y; j > 0; j -= j & -j {
			sum = f.add(sum, f.tree[i][j])
		}
	}
	return sum
}
