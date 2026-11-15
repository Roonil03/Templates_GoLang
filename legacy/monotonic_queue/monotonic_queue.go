// Package monotonic_queue provides a FIFO queue with constant-time minimum queries.
package monotonic_queue

// MonotonicQueue retains a nondecreasing deque of minimum candidates alongside
// its FIFO values. Equal values are retained. Values are integers. The zero value is ready to use. Space is O(n) in queue length;
// Push and Pop take amortized O(1) time, and Min and Len take O(1) time.
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

// NewMonotonicQueue returns an empty queue.
func NewMonotonicQueue() *MonotonicQueue { return &MonotonicQueue{} }

// Push appends value to the FIFO queue and discards larger minimum candidates.
func (q *MonotonicQueue) Push(value int) {
	index := len(q.values)
	q.values = append(q.values, value)
	for len(q.minima) > q.first && compare(q.minima[len(q.minima)-1].value, value) > 0 {
		q.minima[len(q.minima)-1] = entry{}
		q.minima = q.minima[:len(q.minima)-1]
	}
	q.minima = append(q.minima, entry{value: value, index: index})
}

// Pop removes the oldest FIFO value; false means the queue is empty.
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

// Min returns the smallest queued value; false means the queue is empty.
func (q *MonotonicQueue) Min() (int, bool) {
	if q.Len() == 0 {
		var zero int
		return zero, false
	}
	return q.minima[q.first].value, true
}

// Len returns the number of queued FIFO values.
func (q *MonotonicQueue) Len() int { return len(q.values) - q.head }

// SlidingWindowMin returns the minimum in every consecutive window of size k.
// It takes O(n) time and O(k) auxiliary space. It panics unless 1 <= k <= len(data).
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
