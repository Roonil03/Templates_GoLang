package euler_tour

type EulerTour struct {
	n         int
	adj       [][]int
	tin, tout []int
	timer     int
}

func NewEulerTour(n int) *EulerTour {
	return &EulerTour{n: n, adj: make([][]int, n), tin: make([]int, n), tout: make([]int, n)}
}

func (e *EulerTour) AddEdge(u, v int) {
	e.adj[u] = append(e.adj[u], v)
	e.adj[v] = append(e.adj[v], u)
}

func (e *EulerTour) dfs(u, p int) {
	e.timer++
	e.tin[u] = e.timer
	for _, v := range e.adj[u] {
		if v != p {
			e.dfs(v, u)
		}
	}
	e.tout[u] = e.timer
}

func (e *EulerTour) Build(root int) {
	e.timer = 0
	e.dfs(root, -1)
}

// IsAncestor reports whether u lies on the path from the root to v in the built tree.
func (e *EulerTour) IsAncestor(u, v int) bool {
	return e.tin[u] > 0 && e.tin[u] <= e.tin[v] && e.tout[v] <= e.tout[u]
}
