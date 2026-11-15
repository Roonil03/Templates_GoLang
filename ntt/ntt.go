package ntt

const (
	Modulus int64 = 998244353
	MaxLength = 1 << 23
)

type Coefficient interface{ ~int | ~int32 | ~int64 }

func Transform[T Coefficient](a []T, inverse bool) {
	n := len(a)
	if n == 0 {
		return
	}
	if n > MaxLength || n&(n-1) != 0 {
		panic("ntt: invalid transform length")
	}
	for i, value := range a {
		x := int64(value) % Modulus
		if x < 0 {
			x += Modulus
		}
		a[i] = T(x)
	}
	for i, j := 1, 0; i < n; i++ {
		bit := n >> 1
		for j&bit != 0 {
			j ^= bit
			bit >>= 1
		}
		j ^= bit
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	root := int64(3)
	if inverse {
		root = 332748118
	}
	for length := 2; length <= n; length <<= 1 {
		step := power(root, (Modulus-1)/int64(length))
		for start := 0; start < n; start += length {
			w := int64(1)
			for j := 0; j < length/2; j++ {
				u := int64(a[start+j])
				v := int64(a[start+j+length/2]) * w % Modulus
				sum, difference := u+v, u-v
				if sum >= Modulus {
					sum -= Modulus
				}
				if difference < 0 {
					difference += Modulus
				}
				a[start+j], a[start+j+length/2] = T(sum), T(difference)
				w = w * step % Modulus
			}
		}
	}
	if inverse {
		scale := power(int64(n), Modulus-2)
		for i, value := range a {
			a[i] = T(int64(value) * scale % Modulus)
		}
	}
}

func Convolve[T Coefficient](a, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return []T{}
	}
	if len(a) > MaxLength || len(b) > MaxLength || len(a) > MaxLength-len(b)+1 {
		panic("ntt: convolution exceeds maximum length")
	}
	size := len(a) + len(b) - 1
	n := 1
	for n < size {
		n <<= 1
	}
	left, right := make([]T, n), make([]T, n)
	copy(left, a)
	copy(right, b)
	Transform(left, false)
	Transform(right, false)
	for i := range left {
		left[i] = T(int64(left[i]) * int64(right[i]) % Modulus)
	}
	Transform(left, true)
	return left[:size]
}

func power(base, exponent int64) int64 {
	result := int64(1)
	for exponent > 0 {
		if exponent&1 != 0 {
			result = result * base % Modulus
		}
		base = base * base % Modulus
		exponent >>= 1
	}
	return result
}
