package segment_tree

import "cmp"

type SegmentTree[T cmp.Ordered] struct {
	n        int
	data     []T
	identity T
	merge    func(T, T) T
}

func NewSegmentTree[T cmp.Ordered](arr []T, identity T, merge func(T, T) T) *SegmentTree[T] {
	n := len(arr)
	size := 1
	for size < n {
		size <<= 1
	}
	data := make([]T, 2*size)
	for i := 0; i < 2*size; i++ {
		data[i] = identity
	}
	for i := 0; i < n; i++ {
		data[size+i] = arr[i]
	}
	for i := size - 1; i > 0; i-- {
		data[i] = merge(data[2*i], data[2*i+1])
	}
	return &SegmentTree[T]{n: size, data: data, merge: merge, identity: identity}
}

func (st *SegmentTree[T]) Update(i int, v T) {
	pos := i + st.n
	st.data[pos] = v
	for pos > 1 {
		pos >>= 1
		st.data[pos] = st.merge(st.data[2*pos], st.data[2*pos+1])
	}
}

func (st *SegmentTree[T]) Query(l, r int) T {
	resL, resR := st.identity, st.identity
	l += st.n
	r += st.n
	for l <= r {
		if l&1 == 1 {
			resL = st.merge(resL, st.data[l])
			l++
		}
		if r&1 == 0 {
			resR = st.merge(st.data[r], resR)
			r--
		}
		l >>= 1
		r >>= 1
	}
	return st.merge(resL, resR)
}
