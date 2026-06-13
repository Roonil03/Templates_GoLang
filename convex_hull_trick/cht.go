package convex_hull_trick

type SignedNumeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Line[T SignedNumeric] struct {
	M, C T
}

func (l Line[T]) Eval(x T) T {
	return l.M*x + l.C
}

type CHT[T SignedNumeric] struct {
	lines []Line[T]
	ptr   int
}

func NewCHT[T SignedNumeric]() *CHT[T] {
	return &CHT[T]{lines: make([]Line[T], 0), ptr: 0}
}

func Intersect[T SignedNumeric](l1, l2, l3 Line[T]) bool {
	return float64(l3.C-l1.C)*float64(l1.M-l2.M) <= float64(l2.C-l1.C)*float64(l1.M-l3.M)
}

func (c *CHT[T]) Add(m, cVal T) {
	newLine := Line[T]{m, cVal}
	for len(c.lines) >= 2 && Intersect(c.lines[len(c.lines)-2], c.lines[len(c.lines)-1], newLine) {
		c.lines = c.lines[:len(c.lines)-1]
	}
	c.lines = append(c.lines, newLine)
}

func (c *CHT[T]) Query(x T) T {
	if c.ptr >= len(c.lines) {
		c.ptr = len(c.lines) - 1
	}
	for c.ptr < len(c.lines)-1 && c.lines[c.ptr].Eval(x) > c.lines[c.ptr+1].Eval(x) {
		c.ptr++
	}
	return c.lines[c.ptr].Eval(x)
}
