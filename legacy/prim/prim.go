package prim

type Edge struct {
	to     int
	weight int
}

type Item struct {
	to     int
	weight int
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
		if i == j || !(pq.data[j].weight < pq.data[i].weight) {
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

func Prim(n, start int, adj [][]Edge) int {
	visited := make([]bool, n)
	pq := &PriorityQueue{}
	pq.Push(Item{to: start, weight: 0})
	
	total := 0
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
				pq.Push(Item{to: edge.to, weight: edge.weight})
			}
		}
	}
	return total
}
