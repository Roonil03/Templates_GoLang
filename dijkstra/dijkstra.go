package dijkstra

type Number interface {
	~int | ~int64 | ~float64
}

type Edge[W Number] struct {
	to     int
	weight W
}

type Item[W Number] struct {
	node int
	dist W
}

type PriorityQueue[W Number] struct {
	data []Item[W]
}

func (pq *PriorityQueue[W]) Push(x Item[W]) {
	pq.data = append(pq.data, x)
	pq.up(len(pq.data) - 1)
}

func (pq *PriorityQueue[W]) Pop() Item[W] {
	n := len(pq.data) - 1
	pq.data[0], pq.data[n] = pq.data[n], pq.data[0]
	x := pq.data[n]
	pq.data = pq.data[0:n]
	pq.down(0, n)
	return x
}

func (pq *PriorityQueue[W]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !(pq.data[j].dist < pq.data[i].dist) {
			break
		}
		pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
		j = i
	}
}

func (pq *PriorityQueue[W]) down(i0, n int) {
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

func Dijkstra[W Number](n, start int, adj [][]Edge[W], inf W) []W {
	dist := make([]W, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}
	dist[start] = 0
	pq := &PriorityQueue[W]{data: make([]Item[W], 0)}
	pq.Push(Item[W]{start, 0})

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
				pq.Push(Item[W]{v, dist[v]})
			}
		}
	}
	return dist
}
