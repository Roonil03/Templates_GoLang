package fenwick

type Fenwick[T ordered] struct {
	n    int
	tree []T
	add  func(T, T) T
}

func NewFenwick[T ordered](n int, add func(T, T) T) *Fenwick[T] {
	return &Fenwick[T]{n: n, tree: make([]T, n+1), add: add}
}

func (f *Fenwick[T]) Update(i int, v T) {
	if i <= 0 {
		panic("Fenwick updates use one-based indices")
	}
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

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}
