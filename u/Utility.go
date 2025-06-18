package main

func max32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func maxf32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
func maxf64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func minf32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
func minf64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func abs32(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}
func abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
func absf32(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}
func absf64(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
