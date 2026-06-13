package mod_arithmetic

func AddMod(a, b, mod int64) int64 {
	a %= mod
	b %= mod
	res := a + b
	if res >= mod {
		res -= mod
	}
	return res
}

func SubMod(a, b, mod int64) int64 {
	a %= mod
	b %= mod
	res := a - b
	if res < 0 {
		res += mod
	}
	return res
}

func MulMod(a, b, mod int64) int64 {
	a %= mod
	b %= mod
	return (a * b) % mod
}

func PowerMod(base, exp, mod int64) int64 {
	var res int64 = 1
	base %= mod
	for exp > 0 {
		if exp&1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return res
}
