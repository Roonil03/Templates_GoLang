package floyd_warshall

type Number interface {
	~int | ~int64 | ~float64
}

func FloydWarshall[W Number](n int, dist [][]W) {
	floydWarshall(n, dist, nil)
}

func FloydWarshallWithInfinity[W Number](n int, dist [][]W, inf W) {
	floydWarshall(n, dist, &inf)
}
func floydWarshall[W Number](n int, dist [][]W, inf *W) {
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
