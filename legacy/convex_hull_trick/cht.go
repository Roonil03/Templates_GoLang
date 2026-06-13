package convex_hull_trick

type Line struct {
	M, C int64
}

func (l Line) Eval(x int64) int64 {
	return l.M*x + l.C
}

type CHT struct {
	lines []Line
	ptr   int
}

func NewCHT() *CHT {
	return &CHT{lines: make([]Line, 0), ptr: 0}
}

func Intersect(l1, l2, l3 Line) bool {
	return (float64(l3.C)-float64(l1.C))*float64(l1.M-l2.M) <= (float64(l2.C)-float64(l1.C))*float64(l1.M-l3.M)
}

func (c *CHT) Add(m, cVal int64) {
	newLine := Line{m, cVal}
	for len(c.lines) >= 2 && Intersect(c.lines[len(c.lines)-2], c.lines[len(c.lines)-1], newLine) {
		c.lines = c.lines[:len(c.lines)-1]
	}
	c.lines = append(c.lines, newLine)
}

func (c *CHT) Query(x int64) int64 {
	if c.ptr >= len(c.lines) {
		c.ptr = len(c.lines) - 1
	}
	for c.ptr < len(c.lines)-1 && c.lines[c.ptr].Eval(x) > c.lines[c.ptr+1].Eval(x) {
		c.ptr++
	}
	return c.lines[c.ptr].Eval(x)
}
