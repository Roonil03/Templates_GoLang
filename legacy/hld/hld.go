package hld

type HLD struct {
	n                                int
	adj                              [][]int
	parent, depth, heavy, head, pos  []int
	timer                            int
}

func NewHLD(n int) *HLD {
	return &HLD{
		n:      n,
		adj:    make([][]int, n),
		parent: make([]int, n),
		depth:  make([]int, n),
		heavy:  make([]int, n),
		head:   make([]int, n),
		pos:    make([]int, n),
	}
}

func (h *HLD) AddEdge(u, v int) {
	h.adj[u] = append(h.adj[u], v)
	h.adj[v] = append(h.adj[v], u)
}

func (h *HLD) dfsSize(v, p int) int {
	sz := 1
	maxSub := 0
	h.heavy[v] = -1
	for _, to := range h.adj[v] {
		if to != p {
			h.parent[to] = v
			h.depth[to] = h.depth[v] + 1
			subSz := h.dfsSize(to, v)
			sz += subSz
			if subSz > maxSub {
				maxSub = subSz
				h.heavy[v] = to
			}
		}
	}
	return sz
}

func (h *HLD) dfsHLD(v, p, hHead int) {
	h.head[v] = hHead
	h.pos[v] = h.timer
	h.timer++
	if h.heavy[v] != -1 {
		h.dfsHLD(h.heavy[v], v, hHead)
	}
	for _, to := range h.adj[v] {
		if to != p && to != h.heavy[v] {
			h.dfsHLD(to, v, to)
		}
	}
}

func (h *HLD) Build(root int) {
	h.timer = 0
	h.dfsSize(root, -1)
	h.dfsHLD(root, -1, root)
}

func (h *HLD) QueryPath(u, v int, queryFunc func(l, r int)) {
	for h.head[u] != h.head[v] {
		if h.depth[h.head[u]] < h.depth[h.head[v]] {
			u, v = v, u
		}
		queryFunc(h.pos[h.head[u]], h.pos[u])
		u = h.parent[h.head[u]]
	}
	if h.depth[u] > h.depth[v] {
		u, v = v, u
	}
	queryFunc(h.pos[u], h.pos[v])
}
