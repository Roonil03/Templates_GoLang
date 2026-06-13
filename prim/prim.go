package prim

type Number interface {
	~int | ~int64 | ~float64
}

type Edge[W Number] struct {
	to     int
	weight W
}

type Item[W Number] struct {
	to     int
	weight W
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
		if i == j || !(pq.data[j].weight < pq.data[i].weight) {
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
		if j2 := j1 + 1; j2 < n && pq.data[j2].weight < pq.data[j1].weight {
			j = j2
		}
		if !(pq.data[j].weight < pq.data[i].weight) {
			break
		}
		pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
		i = j
	}
}

func Prim[W Number](n, start int, adj [][]Edge[W]) W {
	visited := make([]bool, n)
	pq := &PriorityQueue[W]{}
	pq.Push(Item[W]{to: start, weight: 0})
	
	var total W
	for len(pq.data) > 0 {
		curr := pq.Pop()
		u := curr.to
		if visited[u] {
			continue
		}
		visited[u] = true
		total += curr.weight
		for _, edge := range adj[u] {
			if !visited[edge.to] {
				pq.Push(Item[W]{to: edge.to, weight: edge.weight})
			}
		}
	}
	return total
}
