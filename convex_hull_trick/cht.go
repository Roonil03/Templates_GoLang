package convex_hull_trick

import (
	"math"
	"math/big"
	"sort"
)

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
}

func NewCHT[T SignedNumeric]() *CHT[T] {
	return &CHT[T]{lines: make([]Line[T], 0)}
}

func Intersect[T SignedNumeric](l1, l2, l3 Line[T]) bool {
	if T(1)/T(2) != 0 {
		diff := func(a, b T) *big.Rat {
			x := new(big.Rat).SetFloat64(float64(a))
			y := new(big.Rat).SetFloat64(float64(b))
			if x == nil || y == nil {
				panic("line coefficients must be finite")
			}
			return x.Sub(x, y)
		}
		left := new(big.Rat).Mul(diff(l3.C, l2.C), diff(l1.M, l2.M))
		right := new(big.Rat).Mul(diff(l2.C, l1.C), diff(l2.M, l3.M))
		return left.Cmp(right) <= 0
	}
	diff := func(a, b T) *big.Int { return new(big.Int).Sub(big.NewInt(int64(a)), big.NewInt(int64(b))) }
	left := new(big.Int).Mul(diff(l3.C, l2.C), diff(l1.M, l2.M))
	right := new(big.Int).Mul(diff(l2.C, l1.C), diff(l2.M, l3.M))
	return left.Cmp(right) <= 0
}

func (c *CHT[T]) Add(m, cVal T) {
	if T(1)/T(2) != 0 && (math.IsInf(float64(m), 0) || math.IsNaN(float64(m)) || math.IsInf(float64(cVal), 0) || math.IsNaN(float64(cVal))) {
		panic("line coefficients must be finite")
	}
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
	newLine := Line[T]{m, cVal}
	for len(c.lines) >= 2 && Intersect(c.lines[len(c.lines)-2], c.lines[len(c.lines)-1], newLine) {
		c.lines = c.lines[:len(c.lines)-1]
	}
	c.lines = append(c.lines, newLine)
}

func (c *CHT[T]) Query(x T) T {
	i := sort.Search(len(c.lines)-1, func(i int) bool { return c.lines[i].Eval(x) <= c.lines[i+1].Eval(x) })
	return c.lines[i].Eval(x)
}
