package monotonic_stack

type MonotonicStack[T ordered] struct{ data []T }

func NewMonotonicStack[T ordered]() *MonotonicStack[T] {
	return &MonotonicStack[T]{}
}

func (s *MonotonicStack[T]) Push(value T) {
	for len(s.data) > 0 && compare(s.data[len(s.data)-1], value) > 0 {
		var zero T
		s.data[len(s.data)-1] = zero
		s.data = s.data[:len(s.data)-1]
	}
	s.data = append(s.data, value)
}

func (s *MonotonicStack[T]) Top() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

func (s *MonotonicStack[T]) Pop() (T, bool) {
	value, ok := s.Top()
	if ok {
		var zero T
		s.data[len(s.data)-1] = zero
		s.data = s.data[:len(s.data)-1]
	}
	return value, ok
}

func (s *MonotonicStack[T]) Len() int { return len(s.data) }

func NextGreaterIndices[T ordered](data []T) []int { return nextIndices(data, true) }

func NextSmallerIndices[T ordered](data []T) []int { return nextIndices(data, false) }

func nextIndices[T ordered](data []T, greater bool) []int {
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

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

func compare[T ordered](x, y T) int {
	if x < y || (x != x && y == y) {
		return -1
	}
	if x > y || (x == x && y != y) {
		return 1
	}
	return 0
}
