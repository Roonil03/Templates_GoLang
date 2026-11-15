// Package chinese_remainder solves simultaneous congruences with possibly
// non-coprime positive moduli. Intermediate arithmetic uses math/big to avoid
// overflow; the final period must fit in a positive int64.
package chinese_remainder

import "math/big"

// Solve returns the smallest nonnegative solution and its period (the LCM of
// the moduli). The result represents every solution as remainder + k*modulus.
// Negative remainders are accepted. Empty inputs return (0, 1, true).
// ok is false for mismatched lengths, nonpositive moduli, inconsistent
// congruences, or a period exceeding MaxInt64. Inputs are not modified.
// It performs O(k log M) Euclidean steps for k congruences and largest running
// period M, using bounded-size big integers for overflow-safe intermediate work.
func Solve(remainders, moduli []int64) (remainder, modulus int64, ok bool) {
	if len(remainders) != len(moduli) {
		return 0, 0, false
	}
	r, m := big.NewInt(0), big.NewInt(1)
	for i, value := range moduli {
		if value <= 0 {
			return 0, 0, false
		}
		n := big.NewInt(value)
		a := new(big.Int).Mod(big.NewInt(remainders[i]), n)
		g, inverse := new(big.Int), new(big.Int)
		g.GCD(inverse, nil, m, n)
		difference := new(big.Int).Sub(a, r)
		if new(big.Int).Mod(difference, g).Sign() != 0 {
			return 0, 0, false
		}
		reduced := new(big.Int).Quo(n, g)
		period := new(big.Int).Mul(m, reduced)
		if !period.IsInt64() {
			return 0, 0, false
		}
		step := new(big.Int).Quo(difference, g)
		step.Mul(step, inverse).Mod(step, reduced)
		r.Add(r, new(big.Int).Mul(m, step)).Mod(r, period)
		m = period
	}
	return r.Int64(), m.Int64(), true
}
