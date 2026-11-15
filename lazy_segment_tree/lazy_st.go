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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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

// apply composes a newer update after any pending update at this node.
func (st *LazySegmentTree[V, L]) apply(v, length int, tag L) {
	st.tree[v] = st.upd(st.tree[v], tag, length)
	if st.hasL[v] {
		st.lazy[v] = st.comp(st.lazy[v], tag)
	} else {
		st.lazy[v] = tag
	}
	st.hasL[v] = true
}
func (st *LazySegmentTree[V, L]) push(v, tl, tr int) {
	if !st.hasL[v] || tl == tr {
		return
	}
	mid := tl + (tr-tl)/2
	st.apply(v*2, mid-tl+1, st.lazy[v])
	st.apply(v*2+1, tr-mid, st.lazy[v])
	st.hasL[v] = false
}

func (st *LazySegmentTree[V, L]) Update(v, tl, tr, l, r int, val L) {
	if l > r {
		return
	}
	if l == tl && r == tr {
		st.apply(v, tr-tl+1, val)
		return
	}
	st.push(v, tl, tr)
	tm := tl + (tr-tl)/2
	st.Update(v*2, tl, tm, l, min(r, tm), val)
	st.Update(v*2+1, tm+1, tr, max(l, tm+1), r, val)
	st.tree[v] = st.agg(st.tree[v*2], st.tree[v*2+1])
}

// Query returns the aggregate of the inclusive range [l,r] in the zero-initialized tree.
func (st *LazySegmentTree[V, L]) Query(l, r int) V {
	if l < 0 || l > r || r >= st.size {
		panic("invalid query range")
	}
	return st.query(1, 0, st.size-1, l, r)
}

func (st *LazySegmentTree[V, L]) query(v, tl, tr, l, r int) V {
	if l == tl && r == tr {
		return st.tree[v]
	}
	st.push(v, tl, tr)
	mid := tl + (tr-tl)/2
	if r <= mid {
		return st.query(v*2, tl, mid, l, r)
	}
	if l > mid {
		return st.query(v*2+1, mid+1, tr, l, r)
	}
	return st.agg(st.query(v*2, tl, mid, l, mid), st.query(v*2+1, mid+1, tr, mid+1, r))
}
