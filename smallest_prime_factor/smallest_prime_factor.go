package smallest_prime_factor

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
