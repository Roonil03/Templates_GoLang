// Package dominator_tree computes immediate dominators in directed flow graphs.
package dominator_tree

// Build returns immediate dominators using the Lengauer-Tarjan algorithm with
// path compression and simple linking. Vertices are zero-based. The root is its
// own immediate dominator; unreachable vertices have -1. Parallel edges and loops
// are allowed. Invalid IDs panic; root is ignored when n=0. Time is
// O((V+E) log V), space O(V+E). The input edge list is preserved.
func Build(n int, edges [][2]int, root int) []int {
	graph := make([][]int, n)
	for _, e := range edges {
		if e[1] < 0 || e[1] >= n {
			panic("vertex out of range")
		}
		graph[e[0]] = append(graph[e[0]], e[1])
	}
	result := make([]int, n)
	for v := range result {
		result[v] = -1
	}
	if n == 0 {
		return result
	}
	number := make([]int, n)
	vertex := make([]int, n+1)
	parent := make([]int, n+1)
	count := 0
	var dfs func(int)
	dfs = func(v int) {
		count++
		number[v] = count
		vertex[count] = v
		for _, w := range graph[v] {
			if number[w] == 0 {
				dfs(w)
				parent[number[w]] = number[v]
			}
		}
	}
	dfs(root)
	pred := make([][]int, count+1)
	for u := range graph {
		if number[u] == 0 {
			continue
		}
		for _, v := range graph[u] {
			pred[number[v]] = append(pred[number[v]], number[u])
		}
	}
	semi, label, ancestor, idom := make([]int, count+1), make([]int, count+1), make([]int, count+1), make([]int, count+1)
	bucket := make([][]int, count+1)
	for v := 1; v <= count; v++ {
		semi[v] = v
		label[v] = v
	}
	var compress func(int)
	compress = func(v int) {
		a := ancestor[v]
		if ancestor[a] != 0 {
			compress(a)
			if semi[label[a]] < semi[label[v]] {
				label[v] = label[a]
			}
			ancestor[v] = ancestor[a]
		}
	}
	eval := func(v int) int {
		if ancestor[v] == 0 {
			return label[v]
		}
		compress(v)
		return label[v]
	}
	for w := count; w >= 2; w-- {
		for _, v := range pred[w] {
			u := eval(v)
			if semi[u] < semi[w] {
				semi[w] = semi[u]
			}
		}
		bucket[semi[w]] = append(bucket[semi[w]], w)
		ancestor[w] = parent[w]
		p := parent[w]
		for _, v := range bucket[p] {
			u := eval(v)
			if semi[u] < semi[v] {
				idom[v] = u
			} else {
				idom[v] = p
			}
		}
		bucket[p] = nil
	}
	for w := 2; w <= count; w++ {
		if idom[w] != semi[w] {
			idom[w] = idom[idom[w]]
		}
		result[vertex[w]] = vertex[idom[w]]
	}
	result[root] = root
	return result
}
