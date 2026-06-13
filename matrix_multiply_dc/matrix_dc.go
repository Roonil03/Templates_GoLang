package matrix_multiply_dc

type Numeric interface {
	~int | ~int64 | ~float32 | ~float64
}

type Matrix[T Numeric] struct {
	Data []T
	Rows int
	Cols int
}

func MultiplyDC[T Numeric](a, b *Matrix[T]) *Matrix[T] {
	res := &Matrix[T]{
		Data: make([]T, a.Rows*b.Cols),
		Rows: a.Rows,
		Cols: b.Cols,
	}
	multiplyBlock(a, b, res, 0, 0, 0, 0, 0, 0, a.Rows)
	return res
}

func multiplyBlock[T Numeric](a, b, c *Matrix[T], ar, ac, br, bc, cr, cc, size int) {
	if size == 1 {
		c.Data[(cr)*c.Cols+(cc)] += a.Data[(ar)*a.Cols+(ac)] * b.Data[(br)*b.Cols+(bc)]
		return
	}
	half := size / 2
	multiplyBlock(a, b, c, ar, ac, br, bc, cr, cc, half)
	multiplyBlock(a, b, c, ar, ac+half, br+half, bc, cr, cc, half)

	multiplyBlock(a, b, c, ar, ac, br, bc+half, cr, cc+half, half)
	multiplyBlock(a, b, c, ar, ac+half, br+half, bc+half, cr, cc+half, half)

	multiplyBlock(a, b, c, ar+half, ac, br, bc, cr+half, cc, half)
	multiplyBlock(a, b, c, ar+half, ac+half, br+half, bc, cr+half, cc, half)

	multiplyBlock(a, b, c, ar+half, ac, br, bc+half, cr+half, cc+half, half)
	multiplyBlock(a, b, c, ar+half, ac+half, br+half, bc+half, cr+half, cc+half, half)
}
