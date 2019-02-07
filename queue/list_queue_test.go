package queue

import (
	"testing"
)

func TestListQueue(t *testing.T) {
	q := NewListQueue(6)
	q.Enqueue(2)
	q.Enqueue(2)
	q.Enqueue(2)
	q.Enqueue(2)
	q.Enqueue(2)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	if q.Dequeue() != nil {
		t.Errorf("expected dequeue nil, but it is not a empty queue")
	}

	q.Enqueue(3)
	q.Enqueue(6)
	if q.Dequeue() != 3 {
		t.Errorf("expected 3, but %d", q.Dequeue())
	}
	if q.Dequeue() != 6 {
		t.Errorf("expected 6, but %d", q.Dequeue())
	}
}
