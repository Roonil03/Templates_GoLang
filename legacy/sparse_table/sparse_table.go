package sparse_table

type SparseTable struct {
	st  [][]int
	log []int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NewSparseTable(arr []int) *SparseTable {
	n := len(arr)
	if n == 0 {
		return &SparseTable{}
	}
	log := make([]int, n+1)
	log[1] = 0
	for i := 2; i <= n; i++ {
		log[i] = log[i/2] + 1
	}
	logN := log[n] + 1
	st := make([][]int, logN)
	for i := range st {
		st[i] = make([]int, n)
	}
	copy(st[0], arr)
	for i := 1; i < logN; i++ {
		for j := 0; j+(1<<i) <= n; j++ {
			st[i][j] = min(st[i-1][j], st[i-1][j+(1<<(i-1))])
		}
	}
	return &SparseTable{st: st, log: log}
}

func (s *SparseTable) Query(l, r int) int {
	j := s.log[r-l+1]
	return min(s.st[j][l], s.st[j][r-(1<<j)+1])
}
