package xor_basis

type Basis struct {
	pivots [64]uint64
	rank   int
}

func NewBasis() *Basis { return &Basis{} }

func (b *Basis) Insert(value uint64) bool {
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

func (b *Basis) Contains(value uint64) bool {
	x := uint64(value)
	for bit := 63; bit >= 0; bit-- {
		if x&(uint64(1)<<bit) != 0 {
			x ^= b.pivots[bit]
		}
	}
	return x == 0
}

func (b *Basis) MaxXOR(seed uint64) uint64 {
	x := uint64(seed)
	for bit := 63; bit >= 0; bit-- {
		if candidate := x ^ b.pivots[bit]; candidate > x {
			x = candidate
		}
	}
	return x
}

func (b *Basis) Rank() int { return b.rank }

func (b *Basis) Merge(other *Basis) {
	for _, pivot := range other.pivots {
		b.Insert(pivot)
	}
}
