package rabin_karp

func RabinKarp[T comparable](text, pattern []T, base uint64, val func(T) uint64) []int {
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
		pHash = pHash*base + val(pattern[i])
		tHash = tHash*base + val(text[i])
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
			tHash = tHash - val(text[i])*h
			tHash = tHash*base + val(text[i+m])
		}
	}
	return matches
}
