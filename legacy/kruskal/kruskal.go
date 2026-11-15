package kruskal

import "sort"

type Edge struct {
	U, V   int
	Weight int
}

func Kruskal(n int, edges []Edge) ([]Edge, int) {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = find(parent[i])
		return parent[i]
	}

	union := func(i, j int) bool {
		rootI := find(i)
		rootJ := find(j)
		if rootI != rootJ {
			if rank[rootI] < rank[rootJ] {
				parent[rootI] = rootJ
			} else if rank[rootI] > rank[rootJ] {
				parent[rootJ] = rootI
			} else {
				parent[rootJ] = rootI
				rank[rootI]++
			}
			return true
		}
		return false
	}

	var mst []Edge
	total := 0
	for _, edge := range edges {
		if union(edge.U, edge.V) {
			mst = append(mst, edge)
			total += edge.Weight
		}
	}

	return mst, total
}
