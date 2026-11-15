package xor_basis

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Basis[T Unsigned] struct {
	pivots [64]uint64
	rank   int
}

func NewBasis[T Unsigned]() *Basis[T] { return &Basis[T]{} }

func (b *Basis[T]) Insert(value T) bool {
	x := uint64(value)
	for bit := 63; bit >= 0; bit-- {
		if x&(uint64(1)<<bit) == 0 {
			continue
		}
		if b.pivots[bit] == 0 {
			b.pivots[bit] = x
			b.rank++
			return true
		}
		x ^= b.pivots[bit]
	}
	return false
}

func (b *Basis[T]) Contains(value T) bool {
	x := uint64(value)
	for bit := 63; bit >= 0; bit-- {
		if x&(uint64(1)<<bit) != 0 {
			x ^= b.pivots[bit]
		}
	}
	return x == 0
}

func (b *Basis[T]) MaxXOR(seed T) T {
	x := uint64(seed)
	for bit := 63; bit >= 0; bit-- {
		if candidate := x ^ b.pivots[bit]; candidate > x {
			x = candidate
		}
	}
	return T(x)
}

func (b *Basis[T]) Rank() int { return b.rank }

func (b *Basis[T]) Merge(other *Basis[T]) {
	for _, pivot := range other.pivots {
		b.Insert(T(pivot))
	}
}
