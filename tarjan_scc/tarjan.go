package tarjan_scc

type TarjanSCC[T any] struct {
	graph [][]int
	dfn   []int
	low   []int
	stack []int
	inStk []bool
	timer int
	sccs  [][]int
	nodes []T
}

func NewTarjanSCC[T any](n int, nodes []T) *TarjanSCC[T] {
	return &TarjanSCC[T]{
		graph: make([][]int, n),
		dfn:   make([]int, n),
		low:   make([]int, n),
		inStk: make([]bool, n),
		nodes: nodes,
	}
}

func (t *TarjanSCC[T]) AddEdge(u, v int) {
	t.graph[u] = append(t.graph[u], v)
}

func (t *TarjanSCC[T]) Build() [][]int {
	n := len(t.graph)
	for i := 0; i < n; i++ {
		t.dfn[i] = -1
		t.low[i] = -1
	}
	t.timer = 0
	t.sccs = nil
	t.stack = make([]int, 0, n)

	for i := 0; i < n; i++ {
		if t.dfn[i] == -1 {
			t.dfs(i)
		}
	}
	return t.sccs
}

func (t *TarjanSCC[T]) dfs(u int) {
	t.dfn[u] = t.timer
	t.low[u] = t.timer
	t.timer++
	t.stack = append(t.stack, u)
	t.inStk[u] = true

	for _, v := range t.graph[u] {
		if t.dfn[v] == -1 {
			t.dfs(v)
			if t.low[v] < t.low[u] {
				t.low[u] = t.low[v]
			}
		} else if t.inStk[v] {
			if t.dfn[v] < t.low[u] {
				t.low[u] = t.dfn[v]
			}
		}
	}

	if t.low[u] == t.dfn[u] {
		var scc []int
		for {
			v := t.stack[len(t.stack)-1]
			t.stack = t.stack[:len(t.stack)-1]
			t.inStk[v] = false
			scc = append(scc, v)
			if u == v {
				break
			}
		}
		t.sccs = append(t.sccs, scc)
	}
}
