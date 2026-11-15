package monotonic_stack

type MonotonicStack struct{ data []int }

func NewMonotonicStack() *MonotonicStack {
	return &MonotonicStack{}
}

func (s *MonotonicStack) Push(value int) {
	for len(s.data) > 0 && compare(s.data[len(s.data)-1], value) > 0 {
		var zero int
		s.data[len(s.data)-1] = zero
		s.data = s.data[:len(s.data)-1]
	}
	s.data = append(s.data, value)
}

func (s *MonotonicStack) Top() (int, bool) {
	if len(s.data) == 0 {
		var zero int
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

func (s *MonotonicStack) Pop() (int, bool) {
	value, ok := s.Top()
	if ok {
		var zero int
		s.data[len(s.data)-1] = zero
		s.data = s.data[:len(s.data)-1]
	}
	return value, ok
}

func (s *MonotonicStack) Len() int { return len(s.data) }

func NextGreaterIndices(data []int) []int { return nextIndices(data, true) }

func NextSmallerIndices(data []int) []int { return nextIndices(data, false) }

func nextIndices(data []int, greater bool) []int {
	answer := make([]int, len(data))
	stack := make([]int, 0, len(data))
	for i, value := range data {
		answer[i] = -1
		for len(stack) > 0 {
			j := stack[len(stack)-1]
			order := compare(value, data[j])
			if greater && order <= 0 || !greater && order >= 0 {
				break
			}
			answer[j] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answer
}

func compare(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
