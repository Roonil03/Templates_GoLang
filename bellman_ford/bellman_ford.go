package bellman_ford

type Number interface {
	~int | ~int64 | ~float64
}

type Edge[W Number] struct {
	To     int
	Weight W
}

func SPFA[W Number](n, start int, adj [][]Edge[W], inf W) ([]W, bool) {
	dist := make([]W, n)
	inQueue := make([]bool, n)
	pathLen := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[start] = 0
	q := []int{start}
	inQueue[start] = true

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		inQueue[u] = false

		for _, edge := range adj[u] {
			v := edge.To
			weight := edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				pathLen[v] = pathLen[u] + 1
				if pathLen[v] >= n {
					return nil, true
				}
				if !inQueue[v] {
					q = append(q, v)
					inQueue[v] = true
				}
			}
		}
	}
	return dist, false
}
