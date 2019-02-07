package queue

type ArrayQueue struct {
	data     []interface{}
	capacity int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		data:     make([]interface{}, 0, n),
		capacity: n,
	}
}

func (q *ArrayQueue) Len() int {
	return len(q.data)
}

func (q *ArrayQueue) Enqueue(x interface{}) bool {
	if q.capacity == q.Len() {
		return false
	}

	q.data = append(q.data, x)
	return true
}

func (q *ArrayQueue) Dequeue() interface{} {
	if q.Len() == 0 {
		return nil
	}

	res := q.data[0]

	if q.Len() == 1 {
		q.data = make([]interface{}, 0, q.capacity)
	} else {
		q.data = q.data[1:]
	}

	return res
}
