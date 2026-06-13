package gaussian_elimination

import "math"

func GaussianElimination(a [][]float64, b []float64) []float64 {
	n := len(a)
	for i := 0; i < n; i++ {
		pivot := i
		for j := i + 1; j < n; j++ {
			if math.Abs(a[j][i]) > math.Abs(a[pivot][i]) {
				pivot = j
			}
		}
		a[i], a[pivot] = a[pivot], a[i]
		b[i], b[pivot] = b[pivot], b[i]

		for j := i + 1; j < n; j++ {
			factor := a[j][i] / a[i][i]
			for k := i; k < n; k++ {
				a[j][k] -= factor * a[i][k]
			}
			b[j] -= factor * b[i]
		}
	}

	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := float64(0)
		for j := i + 1; j < n; j++ {
			sum += a[i][j] * x[j]
		}
		x[i] = (b[i] - sum) / a[i][i]
	}
	return x
}
