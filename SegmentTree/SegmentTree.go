package segmenttree

type SegmentTree struct {
	n    int
	data []int
}

func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	size := 1
	for size < n {
		size <<= 1
	}
	data := make([]int, 2*size)
	for i := 0; i < n; i++ {
		data[size+i] = arr[i]
	}
	for i := size - 1; i > 0; i-- {
		data[i] = data[2*i] + data[2*i+1]
	}
	return &SegmentTree{n: size, data: data}
}

func (st *SegmentTree) Update(i, v int) {
	pos := i + st.n
	st.data[pos] = v
	for pos > 1 {
		pos >>= 1
		st.data[pos] = st.data[2*pos] + st.data[2*pos+1]
	}
}

func (st *SegmentTree) Query(l, r int) int {
	res := 0
	l += st.n
	r += st.n
	for l <= r {
		if l&1 == 1 {
			res += st.data[l]
			l++
		}
		if r&1 == 0 {
			res += st.data[r]
			r--
		}
		l >>= 1
		r >>= 1
	}
	return res
}
