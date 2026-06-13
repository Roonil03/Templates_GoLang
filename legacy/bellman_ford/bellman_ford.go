package bellman_ford

type Edge struct {
	to     int
	weight int
}

func SPFA(n, start int, adj [][]Edge, inf int) ([]int, bool) {
	dist := make([]int, n)
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
