package kmp

func ComputePrefix(pattern string) []int {
	n := len(pattern)
	pi := make([]int, n)
	for i := 1; i < n; i++ {
		j := pi[i-1]
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func KMP(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	if m == 0 {
		return nil
	}
	pi := ComputePrefix(pattern)
	var matches []int
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			matches = append(matches, i-m+1)
			j = pi[j-1]
		}
	}
	return matches
}
