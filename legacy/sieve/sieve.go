package sieve

func Sieve(n int) []bool {
	if n < 0 || n == int(^uint(0)>>1) {
		panic("sieve: invalid bound")
	}
	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}
	for p := 2; p <= n/p; p++ {
		if !prime[p] {
			continue
		}
		for multiple := p * p; multiple <= n; {
			prime[multiple] = false
			if multiple > n-p {
				break
			}
			multiple += p
		}
	}
	return prime
}

func Primes(n int) []int {
	prime := Sieve(n)
	result := make([]int, 0)
	for i, ok := range prime {
		if ok {
			result = append(result, i)
		}
	}
	return result
}
