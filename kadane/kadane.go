package kadane

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Kadane[T Number](arr []T) T {
	if len(arr) == 0 {
		var zero T
		return zero
	}
	maxSoFar := arr[0]
	currMax := arr[0]

	for i := 1; i < len(arr); i++ {
		var zero T
		if currMax < zero {
			currMax = arr[i]
		} else {
			currMax += arr[i]
		}
		if currMax > maxSoFar {
			maxSoFar = currMax
		}
	}
	return maxSoFar
}
