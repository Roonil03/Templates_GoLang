package rope

type Rope[T any] struct {
	weight int
	left   *Rope[T]
	right  *Rope[T]
	data   []T
}

func NewRope[T any](s []T) *Rope[T] {
	return &Rope[T]{weight: len(s), data: s}
}

func Concat[T any](a, b *Rope[T]) *Rope[T] {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return &Rope[T]{weight: a.Len(), left: a, right: b}
}

func (r *Rope[T]) Len() int {
	if r == nil {
		return 0
	}
	if r.left == nil && r.right == nil {
		return r.weight
	}
	return r.weight + r.right.Len()
}

func (r *Rope[T]) Index(i int) T {
	if r.left == nil && r.right == nil {
		return r.data[i]
	}
	if i < r.weight {
		return r.left.Index(i)
	}
	return r.right.Index(i - r.weight)
}

func (r *Rope[T]) Split(i int) (*Rope[T], *Rope[T]) {
	if r == nil {
		if i != 0 {
			panic("split index out of range")
		}
		return nil, nil
	}
	if r.left == nil && r.right == nil {
		return NewRope(r.data[:i]), NewRope(r.data[i:])
	}
	if i < r.weight {
		l, r2 := r.left.Split(i)
		return l, Concat(r2, r.right)
	}
	l2, r2 := r.right.Split(i - r.weight)
	return Concat(r.left, l2), r2
}
