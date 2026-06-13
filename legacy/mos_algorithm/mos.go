package mos_algorithm

import "sort"

type Query struct {
	L, R, Idx int
}

func MosAlgorithm(
	arr []int,
	queries []Query,
	blockSize int,
	add func(idx int),
	remove func(idx int),
	getAnswer func() int,
) []int {
	sort.Slice(queries, func(i, j int) bool {
		blockI := queries[i].L / blockSize
		blockJ := queries[j].L / blockSize
		if blockI != blockJ {
			return blockI < blockJ
		}
		if blockI%2 == 1 {
			return queries[i].R < queries[j].R
		}
		return queries[i].R > queries[j].R
	})

	answers := make([]int, len(queries))
	currL := 0
	currR := -1

	for _, q := range queries {
		for currL > q.L {
			currL--
			add(currL)
		}
		for currR < q.R {
			currR++
			add(currR)
		}
		for currL < q.L {
			remove(currL)
			currL++
		}
		for currR > q.R {
			remove(currR)
			currR--
		}
		answers[q.Idx] = getAnswer()
	}

	return answers
}
