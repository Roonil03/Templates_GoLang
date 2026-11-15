package centroid_decomposition

type CentroidDecomposition struct{ adj [][]int }

func NewCentroidDecomposition(n int) *CentroidDecomposition {
	return &CentroidDecomposition{make([][]int, n)}
}
func (c *CentroidDecomposition) AddEdge(u, v int) {
	c.adj[u] = append(c.adj[u], v)
	c.adj[v] = append(c.adj[v], u)
}

func (c *CentroidDecomposition) Build() []int {
	n := len(c.adj)
	parent, size := make([]int, n), make([]int, n)
	removed := make([]bool, n)
	for v := range parent {
		parent[v] = -1
	}
	var measure func(int, int) int
	measure = func(v, p int) int {
		size[v] = 1
		for _, w := range c.adj[v] {
			if w != p && !removed[w] {
				size[v] += measure(w, v)
			}
		}
		return size[v]
	}
	var center func(int, int, int) int
	center = func(v, p, total int) int {
		for _, w := range c.adj[v] {
			if w != p && !removed[w] && size[w] > total/2 {
				return center(w, v, total)
			}
		}
		return v
	}
	var decompose func(int, int)
	decompose = func(v, p int) {
		r := center(v, -1, measure(v, -1))
		parent[r] = p
		removed[r] = true
		for _, w := range c.adj[r] {
			if !removed[w] {
				decompose(w, r)
			}
		}
	}
	for v := 0; v < n; v++ {
		if !removed[v] {
			decompose(v, -1)
		}
	}
	return parent
}
