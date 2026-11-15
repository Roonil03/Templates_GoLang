package mod_arithmetic

import "math/bits"

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func residue[T Integer](x, mod T) T {
	if mod <= 0 {
		panic("modulus must be positive")
	}
	x %= mod
	if x < 0 {
		x += mod
	}
	return x
}

func AddMod[T Integer](a, b, mod T) T {
	a, b = residue(a, mod), residue(b, mod)
	if a >= mod-b {
		return a - (mod - b)
	}
	return a + b
}

func SubMod[T Integer](a, b, mod T) T {
	a, b = residue(a, mod), residue(b, mod)
	if a < b {
		return mod - (b - a)
	}
	return a - b
}

func MulMod[T Integer](a, b, mod T) T {
	hi, lo := bits.Mul64(uint64(residue(a, mod)), uint64(residue(b, mod)))
	_, rem := bits.Div64(hi, lo, uint64(mod))
	return T(rem)
}

func PowerMod[T Integer](base, exp, mod T) T {
	if exp < 0 {
		panic("exponent must be nonnegative")
	}
	base = residue(base, mod)
	result := T(1) % mod
	for exp > 0 {
		if exp&1 == 1 {
			result = MulMod(result, base, mod)
		}
		base = MulMod(base, base, mod)
		exp >>= 1
	}
	return result
}
