package floyd_warshall

type Number interface {
	~int | ~int64 | ~float64
}

func FloydWarshall[W Number](n int, dist [][]W) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
}
