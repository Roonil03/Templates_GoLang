package mod_arithmetic

import "math/bits"

// residue normalizes x for a positive modulus.
func residue(x, mod int64) int64 {
	if mod <= 0 {
		panic("modulus must be positive")
	}
	x %= mod
	if x < 0 {
		x += mod
	}
	return x
}

func AddMod(a, b, mod int64) int64 {
	a, b = residue(a, mod), residue(b, mod)
	if a >= mod-b {
		return a - (mod - b)
	}
	return a + b
}

func SubMod(a, b, mod int64) int64 {
	a, b = residue(a, mod), residue(b, mod)
	if a < b {
		return mod - (b - a)
	}
	return a - b
}

func MulMod(a, b, mod int64) int64 {
	hi, lo := bits.Mul64(uint64(residue(a, mod)), uint64(residue(b, mod)))
	_, rem := bits.Div64(hi, lo, uint64(mod))
	return int64(rem)
}

func PowerMod(base, exp, mod int64) int64 {
	if exp < 0 {
		panic("exponent must be nonnegative")
	}
	base = residue(base, mod)
	result := int64(1) % mod
	for exp > 0 {
		if exp&1 == 1 {
			result = MulMod(result, base, mod)
		}
		base = MulMod(base, base, mod)
		exp >>= 1
	}
	return result
}
