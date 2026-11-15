// Package smallest_prime_factor builds a linear-time smallest-prime-factor
// table in O(n) time and space, supporting O(log x) prime factorization.
package smallest_prime_factor

// SmallestPrimeFactors returns spf for 0 <= x <= n, where spf[x] is the
// smallest prime divisor of x. Entries 0 and 1 are zero. It panics for negative
// n or when n+1 overflows int. The table must fit in available memory.
func SmallestPrimeFactors(n int) []int {
	if n < 0 || n == int(^uint(0)>>1) {
		panic("smallest_prime_factor: invalid bound")
	}
	spf := make([]int, n+1)
	primes := make([]int, 0)
	for x := 2; x <= n; x++ {
		if spf[x] == 0 {
			spf[x] = x
			primes = append(primes, x)
		}
		for _, p := range primes {
			if p > spf[x] || p > n/x {
				break
			}
			spf[p*x] = p
		}
	}
	return spf
}

// Factorize returns x's prime factors in nondecreasing order, with repetition.
// It returns an empty slice for x=1. It panics unless 1 <= x < len(spf).
// spf must be an unmodified table from SmallestPrimeFactors.
func Factorize(x int, spf []int) []int {
	if x < 1 || x >= len(spf) {
		panic("smallest_prime_factor: value outside table")
	}
	factors := make([]int, 0)
	for x > 1 {
		p := spf[x]
		if p < 2 || p > x || x%p != 0 {
			panic("smallest_prime_factor: invalid table")
		}
		factors = append(factors, p)
		x /= p
	}
	return factors
}
