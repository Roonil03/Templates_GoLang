package suffix_array

import (
	"cmp"
	"sort"
)

type SuffixArray[T cmp.Ordered] struct {
	SA  []int
	LCP []int
}

func NewSuffixArray[T cmp.Ordered](s []T) *SuffixArray[T] {
	n := len(s)
	if n == 0 {
		return &SuffixArray[T]{}
	}
	sa := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		sa[i] = i
	}
	sort.Slice(sa, func(i, j int) bool {
		return s[sa[i]] < s[sa[j]]
	})
	classes := 0
	c[sa[0]] = 0
	for i := 1; i < n; i++ {
		if s[sa[i]] != s[sa[i-1]] {
			classes++
		}
		c[sa[i]] = classes
	}
	
	cn := make([]int, n)
	for h := 1; h < n; h <<= 1 {
		sort.Slice(sa, func(i, j int) bool {
			if c[sa[i]] != c[sa[j]] {
				return c[sa[i]] < c[sa[j]]
			}
			nxtI := sa[i] + h
			nxtJ := sa[j] + h
			if nxtI < n && nxtJ < n {
				return c[nxtI] < c[nxtJ]
			}
			return nxtI > nxtJ 
		})
		classes = 0
		cn[sa[0]] = 0
		for i := 1; i < n; i++ {
			pI, pJ := sa[i-1], sa[i]
			nxtI := pI + h
			nxtJ := pJ + h
			diff := c[pI] != c[pJ]
			if !diff {
				if nxtI < n && nxtJ < n {
					diff = c[nxtI] != c[nxtJ]
				} else {
					diff = nxtI != nxtJ
				}
			}
			if diff {
				classes++
			}
			cn[sa[i]] = classes
		}
		copy(c, cn)
		if classes == n-1 {
			break
		}
	}

	lcp := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		rank[sa[i]] = i
	}
	k := 0
	for i := 0; i < n; i++ {
		if rank[i] == n-1 {
			k = 0
			continue
		}
		j := sa[rank[i]+1]
		for i+k < n && j+k < n && s[i+k] == s[j+k] {
			k++
		}
		lcp[rank[i]] = k
		if k > 0 {
			k--
		}
	}

	return &SuffixArray[T]{SA: sa, LCP: lcp}
}
