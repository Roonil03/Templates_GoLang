package dijkstra

type Edge struct {
	to     int
	weight int
}

type Item struct {
	node int
	dist int
}

type PriorityQueue struct {
	data []Item
}

func (pq *PriorityQueue) Push(x Item) {
	pq.data = append(pq.data, x)
	pq.up(len(pq.data) - 1)
}

func (pq *PriorityQueue) Pop() Item {
	n := len(pq.data) - 1
	pq.data[0], pq.data[n] = pq.data[n], pq.data[0]
	x := pq.data[n]
	pq.data = pq.data[0:n]
	pq.down(0, n)
	return x
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !(pq.data[j].dist < pq.data[i].dist) {
			break
		}
		pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
		j = i
	}
}

func (pq *PriorityQueue) down(i0, n int) {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && pq.data[j2].dist < pq.data[j1].dist {
			j = j2
		}
		if !(pq.data[j].dist < pq.data[i].dist) {
			break
		}
		pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
		i = j
	}
}

func Dijkstra(n, start int, adj [][]Edge, inf int) []int {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[start] = 0
	pq := &PriorityQueue{data: make([]Item, 0)}
	pq.Push(Item{start, 0})

	for len(pq.data) > 0 {
		curr := pq.Pop()
		u := curr.node
		d := curr.dist
		if d > dist[u] {
			continue
		}
		for _, edge := range adj[u] {
			v := edge.to
			weight := edge.weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				pq.Push(Item{v, dist[v]})
			}
		}
	}
	return dist
}
