// Package xor_basis maintains a linear basis of unsigned integers under XOR.
// Insert, Contains, and MaxXOR take O(w) time; a basis uses O(w) space, where
// w <= 64 is the integer width. Merge takes O(w squared) time.
package xor_basis

// Basis represents every XOR combination of its inserted values. Its zero
// value is an empty usable basis. Copying a Basis produces an independent copy.
type Basis struct {
	pivots [64]uint64
	rank   int
}

// NewBasis returns an empty basis.
func NewBasis() *Basis { return &Basis{} }

// Insert adds value and returns whether it increased the rank. Zero and values
// already represented by the basis do not increase its rank.
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

// Contains reports whether value is an XOR combination of inserted values.
// The empty combination represents zero even in an empty basis.
func (b *Basis) Contains(value uint64) bool {
	x := uint64(value)
	for bit := 63; bit >= 0; bit-- {
		if x&(uint64(1)<<bit) != 0 {
			x ^= b.pivots[bit]
		}
	}
	return x == 0
}

// MaxXOR returns the largest unsigned value of seed XOR any represented value.
// A seed of zero returns the largest representable value.
func (b *Basis) MaxXOR(seed uint64) uint64 {
	x := uint64(seed)
	for bit := 63; bit >= 0; bit-- {
		if candidate := x ^ b.pivots[bit]; candidate > x {
			x = candidate
		}
	}
	return x
}

// Rank returns the number of linearly independent inserted values.
func (b *Basis) Rank() int { return b.rank }

// Merge adds every value represented by other to b. It does not modify other.
// A basis can be merged with itself.
func (b *Basis) Merge(other *Basis) {
	for _, pivot := range other.pivots {
		b.Insert(pivot)
	}
}
