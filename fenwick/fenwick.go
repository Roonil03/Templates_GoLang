package fenwick

import "cmp"

type Fenwick[T cmp.Ordered] struct {
	n    int
	tree []T
	add  func(T, T) T
}

func NewFenwick[T cmp.Ordered](n int, add func(T, T) T) *Fenwick[T] {
	return &Fenwick[T]{n: n, tree: make([]T, n+1), add: add}
}

func (f *Fenwick[T]) Update(i int, v T) {
	for ; i <= f.n; i += i & -i {
		f.tree[i] = f.add(f.tree[i], v)
	}
}

func (f *Fenwick[T]) Query(i int, zero T) T {
	sum := zero
	for ; i > 0; i -= i & -i {
		sum = f.add(sum, f.tree[i])
	}
	return sum
}
