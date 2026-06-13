package maxflow

type Edge struct {
	to, rev int
	cap, flow int64
}

type MaxFlow struct {
	graph [][]Edge
	level []int
	ptr   []int
}

func NewMaxFlow(n int) *MaxFlow {
	return &MaxFlow{
		graph: make([][]Edge, n),
		level: make([]int, n),
		ptr:   make([]int, n),
	}
}

func (mf *MaxFlow) AddEdge(from, to int, cap int64) {
	mf.graph[from] = append(mf.graph[from], Edge{to, len(mf.graph[to]), cap, 0})
	mf.graph[to] = append(mf.graph[to], Edge{from, len(mf.graph[from]) - 1, 0, 0})
}

func (mf *MaxFlow) bfs(s, t int) bool {
	for i := range mf.level {
		mf.level[i] = -1
	}
	mf.level[s] = 0
	q := []int{s}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, e := range mf.graph[v] {
			if e.cap-e.flow > 0 && mf.level[e.to] == -1 {
				mf.level[e.to] = mf.level[v] + 1
				q = append(q, e.to)
			}
		}
	}
	return mf.level[t] != -1
}

func (mf *MaxFlow) dfs(v, t int, pushed int64) int64 {
	if pushed == 0 {
		return 0
	}
	if v == t {
		return pushed
	}
	for ; mf.ptr[v] < len(mf.graph[v]); mf.ptr[v]++ {
		e := &mf.graph[v][mf.ptr[v]]
		tr := e.to
		if mf.level[v]+1 != mf.level[tr] || e.cap-e.flow == 0 {
			continue
		}
		push := pushed
		if e.cap-e.flow < push {
			push = e.cap - e.flow
		}
		p := mf.dfs(tr, t, push)
		if p == 0 {
			continue
		}
		e.flow += p
		mf.graph[tr][e.rev].flow -= p
		return p
	}
	return 0
}

func (mf *MaxFlow) Dinic(s, t int, inf int64) int64 {
	var flow int64 = 0
	for mf.bfs(s, t) {
		for i := range mf.ptr {
			mf.ptr[i] = 0
		}
		for {
			pushed := mf.dfs(s, t, inf)
			if pushed == 0 {
				break
			}
			flow += pushed
		}
	}
	return flow
}
