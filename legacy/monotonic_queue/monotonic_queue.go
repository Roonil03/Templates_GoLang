package monotonic_queue

type MonotonicQueue struct {
	values []int
	head   int
	minima []entry
	first  int
}
type entry struct {
	value int
	index int
}

func NewMonotonicQueue() *MonotonicQueue { return &MonotonicQueue{} }

func (q *MonotonicQueue) Push(value int) {
	index := len(q.values)
	q.values = append(q.values, value)
	for len(q.minima) > q.first && compare(q.minima[len(q.minima)-1].value, value) > 0 {
		q.minima[len(q.minima)-1] = entry{}
		q.minima = q.minima[:len(q.minima)-1]
	}
	q.minima = append(q.minima, entry{value: value, index: index})
}

func (q *MonotonicQueue) Pop() (int, bool) {
	if q.Len() == 0 {
		var zero int
		return zero, false
	}
	value := q.values[q.head]
	var zero int
	q.values[q.head] = zero
	if q.minima[q.first].index == q.head {
		q.minima[q.first] = entry{}
		q.first++
	}
	q.head++
	if q.head == len(q.values) {
		q.values = nil
		q.minima = nil
		q.head = 0
		q.first = 0
	} else if q.head >= len(q.values)-q.head {
		q.values = append([]int(nil), q.values[q.head:]...)
		q.minima = append([]entry(nil), q.minima[q.first:]...)
		for i := range q.minima {
			q.minima[i].index -= q.head
		}
		q.head = 0
		q.first = 0
	} else if q.first > 0 && q.first >= len(q.minima)-q.first {
		q.minima = append([]entry(nil), q.minima[q.first:]...)
		q.first = 0
	}
	return value, true
}

func (q *MonotonicQueue) Min() (int, bool) {
	if q.Len() == 0 {
		var zero int
		return zero, false
	}
	return q.minima[q.first].value, true
}

func (q *MonotonicQueue) Len() int { return len(q.values) - q.head }

func SlidingWindowMin(data []int, k int) []int {
	if k < 1 || k > len(data) {
		panic("monotonic_queue: invalid window size")
	}
	q := NewMonotonicQueue()
	answer := make([]int, 0, len(data)-k+1)
	for i, value := range data {
		if q.Len() == k {
			q.Pop()
		}
		q.Push(value)
		if i >= k-1 {
			minimum, _ := q.Min()
			answer = append(answer, minimum)
		}
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
