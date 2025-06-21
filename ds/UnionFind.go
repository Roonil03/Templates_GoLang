package ds

type DSU struct {
	parent, rank []int
}

func NewDSU(n int) *DSU {
	p, r := make([]int, n), make([]int, n)
	for i := range p {
		p[i] = i
	}
	return &DSU{parent: p, rank: r}
}

func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Union(x, y int) {
	x, y = d.Find(x), d.Find(y)
	if x == y {
		return
	}
	if d.rank[x] < d.rank[y] {
		d.parent[x] = y
	} else {
		d.parent[y] = x
		if d.rank[x] == d.rank[y] {
			d.rank[x]++
		}
	}
}
