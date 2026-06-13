package maxflow

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Edge[T Number] struct {
	to, rev int
	cap, flow T
}

type MaxFlow[T Number] struct {
	graph [][]Edge[T]
	level []int
	ptr   []int
}

func NewMaxFlow[T Number](n int) *MaxFlow[T] {
	return &MaxFlow[T]{
		graph: make([][]Edge[T], n),
		level: make([]int, n),
		ptr:   make([]int, n),
	}
}

func (mf *MaxFlow[T]) AddEdge(from, to int, cap T) {
	mf.graph[from] = append(mf.graph[from], Edge[T]{to, len(mf.graph[to]), cap, 0})
	mf.graph[to] = append(mf.graph[to], Edge[T]{from, len(mf.graph[from]) - 1, 0, 0})
}

func (mf *MaxFlow[T]) bfs(s, t int) bool {
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

func (mf *MaxFlow[T]) dfs(v, t int, pushed T) T {
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

func (mf *MaxFlow[T]) Dinic(s, t int, inf T) T {
	var flow T = 0
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
