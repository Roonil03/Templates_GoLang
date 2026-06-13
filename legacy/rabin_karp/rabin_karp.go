package rabin_karp

func RabinKarp(text, pattern string, base uint64) []int {
	n := len(text)
	m := len(pattern)
	if m == 0 || n < m {
		return nil
	}

	var pHash, tHash uint64
	var h uint64 = 1

	for i := 0; i < m-1; i++ {
		h *= base
	}

	for i := 0; i < m; i++ {
		pHash = pHash*base + uint64(pattern[i])
		tHash = tHash*base + uint64(text[i])
	}

	var matches []int
	for i := 0; i <= n-m; i++ {
		if pHash == tHash {
			match := true
			for j := 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				matches = append(matches, i)
			}
		}
		if i < n-m {
			tHash = tHash - uint64(text[i])*h
			tHash = tHash*base + uint64(text[i+m])
		}
	}
	return matches
}
