package bellman_ford

type Number interface {
	~int | ~int64 | ~float64
}

type Edge[W Number] struct {
	to     int
	weight W
}

func SPFA[W Number](n, start int, adj [][]Edge[W], inf W) ([]W, bool) {
	dist := make([]W, n)
	inQueue := make([]bool, n)
	count := make([]int, n)
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
			v := edge.to
			weight := edge.weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				if !inQueue[v] {
					q = append(q, v)
					inQueue[v] = true
					count[v]++
					if count[v] >= n {
						return nil, true 
					}
				}
			}
		}
	}
	return dist, false
}
