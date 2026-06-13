package kahn

func Kahn(n int, edges [][]int) ([]int, bool) {
	inDegree := make([]int, n)
	adj := make([][]int, n)

	for u, neighbors := range edges {
		for _, v := range neighbors {
			adj[u] = append(adj[u], v)
			inDegree[v]++
		}
	}

	q := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	var order []int
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		order = append(order, u)

		for _, v := range adj[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(order) == n {
		return order, false
	}
	return nil, true
}
