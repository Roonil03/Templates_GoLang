// Package euler_path checks Euler-trail feasibility and finds its endpoints.
package euler_path

// Endpoints returns suitable start/end vertices for a trail using every edge
// once. Directed edges are ordered pairs; undirected loops and parallel edges
// are allowed. Only vertices incident to edges must be connected (weakly for
// directed graphs). Empty graphs return (0,0,true), or (-1,-1,true) when n=0.
// Invalid vertex IDs panic. Time and space are O(V+E).
func Endpoints(n int, edges [][2]int, directed bool) (start, end int, ok bool) {
	adj := make([][]int, n)
	in, out := make([]int, n), make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		out[u]++
		if directed {
			in[v]++
		} else {
			out[v]++
		}
	}
	if len(edges) == 0 {
		if n == 0 {
			return -1, -1, true
		}
		return 0, 0, true
	}
	start, end = -1, -1
	for v := 0; v < n; v++ {
		if directed {
			switch out[v] - in[v] {
			case 1:
				if start != -1 {
					return -1, -1, false
				}
				start = v
			case -1:
				if end != -1 {
					return -1, -1, false
				}
				end = v
			case 0:
			default:
				return -1, -1, false
			}
		} else if out[v]%2 != 0 {
			if start == -1 {
				start = v
			} else if end == -1 {
				end = v
			} else {
				return -1, -1, false
			}
		}
	}
	if (start == -1) != (end == -1) {
		return -1, -1, false
	}
	if start == -1 {
		start = edges[0][0]
		end = start
	}
	seen := make([]bool, n)
	seen[start] = true
	stack := []int{start}
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, v := range adj[u] {
			if !seen[v] {
				seen[v] = true
				stack = append(stack, v)
			}
		}
	}
	for v := 0; v < n; v++ {
		if len(adj[v]) > 0 && !seen[v] {
			return -1, -1, false
		}
	}
	return start, end, true
}
