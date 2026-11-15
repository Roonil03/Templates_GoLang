package kruskal

import (
	"sort"
)

type Edge[W ordered] struct {
	U, V   int
	Weight W
}

func Kruskal[W ordered](n int, edges []Edge[W]) []Edge[W] {
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

	var mst []Edge[W]
	for _, edge := range edges {
		if union(edge.U, edge.V) {
			mst = append(mst, edge)
		}
	}

	return mst
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}
