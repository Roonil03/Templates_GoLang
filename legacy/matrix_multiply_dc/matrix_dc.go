package matrix_multiply_dc

type Matrix struct {
	Data []int64
	Rows int
	Cols int
}

func MultiplyDC(a, b *Matrix) *Matrix {
	if a.Cols != b.Rows {
		panic("incompatible matrix dimensions")
	}
	c := &Matrix{Rows: a.Rows, Cols: b.Cols, Data: make([]int64, a.Rows*b.Cols)}
	multiplyBlock(a, b, c, 0, a.Rows, 0, a.Cols, 0, b.Cols)
	return c
}

func multiplyBlock(a, b, c *Matrix, r, endR, k, endK, col, endC int) {
	rows, inner, cols := endR-r, endK-k, endC-col
	if rows == 0 || inner == 0 || cols == 0 {
		return
	}
	switch {
	case rows >= inner && rows >= cols && rows > 16:
		mid := r + rows/2
		multiplyBlock(a, b, c, r, mid, k, endK, col, endC)
		multiplyBlock(a, b, c, mid, endR, k, endK, col, endC)
	case inner >= cols && inner > 16:
		mid := k + inner/2
		multiplyBlock(a, b, c, r, endR, k, mid, col, endC)
		multiplyBlock(a, b, c, r, endR, mid, endK, col, endC)
	case cols > 16:
		mid := col + cols/2
		multiplyBlock(a, b, c, r, endR, k, endK, col, mid)
		multiplyBlock(a, b, c, r, endR, k, endK, mid, endC)
	default:
		for i := r; i < endR; i++ {
			for j := k; j < endK; j++ {
				for x := col; x < endC; x++ {
					c.Data[i*c.Cols+x] += a.Data[i*a.Cols+j] * b.Data[j*b.Cols+x]
				}
			}
		}
	}
}
