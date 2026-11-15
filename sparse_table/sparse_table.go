package sparse_table

type SparseTable[T any] struct {
	st  [][]T
	log []int
	op  func(T, T) T
}

// NewSparseTable builds an immutable range-query table. The operation must be
// associative and idempotent, such as minimum, maximum, or greatest common divisor.
// Construction takes O(n log n) time and space, assuming op takes O(1) time.
func NewSparseTable[T any](arr []T, op func(T, T) T) *SparseTable[T] {
	n := len(arr)
	if n == 0 {
		return &SparseTable[T]{op: op}
	}
	log := make([]int, n+1)
	log[1] = 0
	for i := 2; i <= n; i++ {
		log[i] = log[i/2] + 1
	}
	logN := log[n] + 1
	st := make([][]T, logN)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], arr)
	for i := 1; i < logN; i++ {
		for j := 0; j+(1<<i) <= n; j++ {
			st[i][j] = op(st[i-1][j], st[i-1][j+(1<<(i-1))])
		}
	}
	return &SparseTable[T]{st: st, log: log, op: op}
}

// Query combines the inclusive range [l,r] in O(1) calls to op. It uses overlapping
// blocks and requires the associative, idempotent operation supplied at construction.
// The range must satisfy 0 <= l <= r < len(arr).
func (s *SparseTable[T]) Query(l, r int) T {
	j := s.log[r-l+1]
	return s.op(s.st[j][l], s.st[j][r-(1<<j)+1])
}
