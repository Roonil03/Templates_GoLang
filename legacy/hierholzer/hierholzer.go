// Package hierholzer constructs Euler trails in directed or undirected graphs.
package hierholzer

import "github.com/Roonil03/Templates_GoLang/legacy/euler_path"

// Trail returns the vertex sequence of an Euler trail, or (nil,false) if none
// exists. Vertices are zero-based; invalid IDs panic. Loops and parallel edges
// are distinct edges. Input edges are not modified. An edgeless graph returns
// [0] when n>0 and an empty slice when n=0. Time and space are O(V+E).
func Trail(n int, edges [][2]int, directed bool) ([]int, bool) {
	start, _, ok := euler_path.Endpoints(n, edges, directed)
	if !ok {
		return nil, false
	}
	if start == -1 {
		return []int{}, true
	}
	type arc struct{ to, id int }
	adj := make([][]arc, n)
	for id, e := range edges {
		adj[e[0]] = append(adj[e[0]], arc{e[1], id})
		if !directed {
			adj[e[1]] = append(adj[e[1]], arc{e[0], id})
		}
	}
	used := make([]bool, len(edges))
	next := make([]int, n)
	stack := []int{start}
	path := make([]int, 0, len(edges)+1)
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		for next[u] < len(adj[u]) && used[adj[u][next[u]].id] {
			next[u]++
		}
		if next[u] == len(adj[u]) {
			path = append(path, u)
			stack = stack[:len(stack)-1]
		} else {
			e := adj[u][next[u]]
			next[u]++
			used[e.id] = true
			stack = append(stack, e.to)
		}
	}
	if len(path) != len(edges)+1 {
		return nil, false
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path, true
}
