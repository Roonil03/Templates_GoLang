package kadane

func Kadane(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxSoFar := arr[0]
	currMax := arr[0]

	for i := 1; i < len(arr); i++ {
		if currMax < 0 {
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
