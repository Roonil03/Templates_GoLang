package fenwick2d

type Fenwick2D[T ordered] struct {
	n, m int
	tree [][]T
	add  func(T, T) T
}

func NewFenwick2D[T ordered](n, m int, add func(T, T) T) *Fenwick2D[T] {
	b := make([][]T, n+1)
	for i := range b {
		b[i] = make([]T, m+1)
	}
	return &Fenwick2D[T]{n: n, m: m, tree: b, add: add}
}

func (f *Fenwick2D[T]) Update(x, y int, v T) {
	if x <= 0 || y <= 0 {
		panic("Fenwick updates use one-based indices")
	}
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

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}
