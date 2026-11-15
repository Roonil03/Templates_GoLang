package floyd_warshall

// FloydWarshall updates a dense matrix in place. All intermediate sums must fit the weight type.
// For an integer sentinel indicating missing edges, use FloydWarshallWithInfinity.
func FloydWarshall(n int, dist [][]int) {
	floydWarshall(n, dist, nil)
}

// FloydWarshallWithInfinity skips missing paths marked with inf.
func FloydWarshallWithInfinity(n int, dist [][]int, inf int) {
	floydWarshall(n, dist, &inf)
}
func floydWarshall(n int, dist [][]int, inf *int) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if inf != nil && dist[i][k] == *inf {
				continue
			}
			for j := 0; j < n; j++ {
				if inf != nil && dist[k][j] == *inf {
					continue
				}
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
}
