// Package lca answers lowest-common-ancestor queries in an undirected forest.
// Vertices are zero-based. Add edges before Build; inputs must form a forest.
package lca

import "math/bits"

// LCA stores binary-lifting ancestors. Build costs O(V log V) time and space;
// Query and KthAncestor take O(log V) time.
type LCA struct {
	adj              [][]int
	up               [][]int
	depth, component []int
}

func NewLCA(n int) *LCA         { return &LCA{adj: make([][]int, n)} }
func (l *LCA) AddEdge(u, v int) { l.adj[u] = append(l.adj[u], v); l.adj[v] = append(l.adj[v], u) }

// Build roots the first component at root and other components at their first
// unvisited vertex. On an empty forest root is ignored. Rebuilding is supported.
func (l *LCA) Build(root int) {
	n := len(l.adj)
	levels := bits.Len(uint(n))
	if levels == 0 {
		levels = 1
	}
	l.up = make([][]int, levels)
	for k := range l.up {
		l.up[k] = make([]int, n)
		for v := range l.up[k] {
			l.up[k][v] = -1
		}
	}
	l.depth = make([]int, n)
	l.component = make([]int, n)
	for v := range l.component {
		l.component[v] = -1
	}
	if n == 0 {
		return
	}
	roots := []int{root}
	for v := 0; v < n; v++ {
		roots = append(roots, v)
	}
	for _, r := range roots {
		if l.component[r] != -1 {
			continue
		}
		l.component[r] = r
		stack := []int{r}
		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, w := range l.adj[v] {
				if l.component[w] != -1 {
					continue
				}
				l.component[w] = r
				l.depth[w] = l.depth[v] + 1
				l.up[0][w] = v
				stack = append(stack, w)
			}
		}
	}
	for k := 1; k < levels; k++ {
		for v := 0; v < n; v++ {
			if p := l.up[k-1][v]; p != -1 {
				l.up[k][v] = l.up[k-1][p]
			}
		}
	}
}

// KthAncestor returns -1 when k is negative or exceeds v's depth.
func (l *LCA) KthAncestor(v, k int) int {
	if k < 0 || k > l.depth[v] {
		return -1
	}
	for j := 0; k > 0; j++ {
		if k&1 != 0 {
			v = l.up[j][v]
		}
		k >>= 1
	}
	return v
}

// Query returns -1 for vertices in different components. Build must run first.
func (l *LCA) Query(u, v int) int {
	if l.component[u] != l.component[v] {
		return -1
	}
	if l.depth[u] < l.depth[v] {
		u, v = v, u
	}
	u = l.KthAncestor(u, l.depth[u]-l.depth[v])
	if u == v {
		return u
	}
	for k := len(l.up) - 1; k >= 0; k-- {
		if l.up[k][u] != l.up[k][v] {
			u, v = l.up[k][u], l.up[k][v]
		}
	}
	return l.up[0][u]
}
