// Package hopcroft_karp computes maximum bipartite matchings.
package hopcroft_karp

// HopcroftKarp uses zero-based, separate left and right vertex sets.
// MaximumMatching takes O((V+E) sqrt(V)) time and O(V+E) space.
type HopcroftKarp struct {
	adj        [][]int
	rightCount int
}

func NewHopcroftKarp(left, right int) *HopcroftKarp { return &HopcroftKarp{make([][]int, left), right} }
func (h *HopcroftKarp) AddEdge(left, right int) {
	if right < 0 || right >= h.rightCount {
		panic("right vertex out of range")
	}
	h.adj[left] = append(h.adj[left], right)
}

// MaximumMatching recomputes and returns both partner arrays (-1 for unmatched)
// and the matching cardinality. Duplicate edges are allowed.
func (h *HopcroftKarp) MaximumMatching() (left, right []int, size int) {
	left = make([]int, len(h.adj))
	right = make([]int, h.rightCount)
	for i := range left {
		left[i] = -1
	}
	for i := range right {
		right[i] = -1
	}
	dist := make([]int, len(left))
	inf := len(left) + 1
	bfs := func() int {
		q := make([]int, 0, len(left))
		for u := range left {
			if left[u] == -1 {
				dist[u] = 0
				q = append(q, u)
			} else {
				dist[u] = inf
			}
		}
		shortest := inf
		for head := 0; head < len(q); head++ {
			u := q[head]
			if dist[u] >= shortest {
				continue
			}
			for _, v := range h.adj[u] {
				w := right[v]
				if w == -1 {
					shortest = dist[u] + 1
				} else if dist[w] == inf {
					dist[w] = dist[u] + 1
					q = append(q, w)
				}
			}
		}
		return shortest
	}
	var dfs func(int, int) bool
	dfs = func(u, shortest int) bool {
		for _, v := range h.adj[u] {
			w := right[v]
			if (w == -1 && dist[u]+1 == shortest) || (w != -1 && dist[w] == dist[u]+1 && dfs(w, shortest)) {
				left[u] = v
				right[v] = u
				return true
			}
		}
		dist[u] = inf
		return false
	}
	for {
		shortest := bfs()
		if shortest == inf {
			break
		}
		for u := range left {
			if left[u] == -1 && dfs(u, shortest) {
				size++
			}
		}
	}
	return
}
