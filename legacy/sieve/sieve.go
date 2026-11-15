// Package sieve generates primes with the sieve of Eratosthenes.
// Bounds are inclusive; construction takes O(n log log n) time and O(n) space.
package sieve

// Sieve returns a table whose entry i reports whether i is prime, for 0 <= i <= n.
// It panics if n is negative or n+1 cannot be represented by int. The caller
// must provide a bound small enough that the table fits in available memory.
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

// Primes returns all primes at most n in ascending order. Its bounds and panic
// conditions are the same as Sieve's. No primes produces an empty slice.
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
