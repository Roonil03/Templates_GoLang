package lazy_segment_tree

type LazySegmentTree[V any, L any] struct {
	tree []V
	lazy []L
	hasL []bool
	size int
	agg  func(V, V) V
	upd  func(V, L, int) V
	comp func(L, L) L
}

func min(a, b int) int { if a < b { return a }; return b }
func max(a, b int) int { if a > b { return a }; return b }

func NewLazySegmentTree[V any, L any](size int, agg func(V, V) V, upd func(V, L, int) V, comp func(L, L) L) *LazySegmentTree[V, L] {
	return &LazySegmentTree[V, L]{
		tree: make([]V, 4*size),
		lazy: make([]L, 4*size),
		hasL: make([]bool, 4*size),
		size: size,
		agg:  agg,
		upd:  upd,
		comp: comp,
	}
}

func (st *LazySegmentTree[V, L]) push(v, tl, tr int) {
	if st.hasL[v] {
		tm := tl + (tr-tl)/2
		st.lazy[v*2] = st.comp(st.lazy[v*2], st.lazy[v])
		st.hasL[v*2] = true
		st.tree[v*2] = st.upd(st.tree[v*2], st.lazy[v], tm-tl+1)

		st.lazy[v*2+1] = st.comp(st.lazy[v*2+1], st.lazy[v])
		st.hasL[v*2+1] = true
		st.tree[v*2+1] = st.upd(st.tree[v*2+1], st.lazy[v], tr-tm)

		st.hasL[v] = false
	}
}

func (st *LazySegmentTree[V, L]) Update(v, tl, tr, l, r int, val L) {
	if l > r {
		return
	}
	if l == tl && r == tr {
		st.lazy[v] = st.comp(st.lazy[v], val)
		st.hasL[v] = true
		st.tree[v] = st.upd(st.tree[v], val, tr-tl+1)
		return
	}
	st.push(v, tl, tr)
	tm := tl + (tr-tl)/2
	st.Update(v*2, tl, tm, l, min(r, tm), val)
	st.Update(v*2+1, tm+1, tr, max(l, tm+1), r, val)
	st.tree[v] = st.agg(st.tree[v*2], st.tree[v*2+1])
}
