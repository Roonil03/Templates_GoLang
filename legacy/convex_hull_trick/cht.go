package convex_hull_trick

import (
	"math/big"
	"sort"
)

type Line struct {
	M, C int64
}

func (l Line) Eval(x int64) int64 {
	return l.M*x + l.C
}

type CHT struct {
	lines []Line
}

func NewCHT() *CHT {
	return &CHT{lines: make([]Line, 0)}
}

func Intersect(l1, l2, l3 Line) bool {
	// With descending slopes the middle line is redundant if its intersection
	// with the first is at or beyond its intersection with the third.

	diff := func(a, b int64) *big.Int { return new(big.Int).Sub(big.NewInt(int64(a)), big.NewInt(int64(b))) }
	left := new(big.Int).Mul(diff(l3.C, l2.C), diff(l1.M, l2.M))
	right := new(big.Int).Mul(diff(l2.C, l1.C), diff(l2.M, l3.M))
	return left.Cmp(right) <= 0
}

// Add inserts a line in nonincreasing slope order.
func (c *CHT) Add(m, cVal int64) {
	if n := len(c.lines); n > 0 {
		last := c.lines[n-1]
		if m > last.M {
			panic("slopes must be nonincreasing")
		}
		if m == last.M {
			if last.C <= cVal {
				return
			}
			c.lines = c.lines[:n-1]
		}
	}
	newLine := Line{m, cVal}
	for len(c.lines) >= 2 && Intersect(c.lines[len(c.lines)-2], c.lines[len(c.lines)-1], newLine) {
		c.lines = c.lines[:len(c.lines)-1]
	}
	c.lines = append(c.lines, newLine)
}

// Query returns the minimum at x. Line evaluations must fit the chosen numeric type.
func (c *CHT) Query(x int64) int64 {
	i := sort.Search(len(c.lines)-1, func(i int) bool { return c.lines[i].Eval(x) <= c.lines[i+1].Eval(x) })
	return c.lines[i].Eval(x)
}
