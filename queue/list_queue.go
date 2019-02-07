package queue

type listNode struct {
	val  interface{}
	next *listNode
}

type ListQueue struct {
	head     *listNode
	tail     *listNode
	capacity int
	length   int
}

func NewListQueue(n int) *ListQueue {
	return &ListQueue{
		capacity: n,
		length:   0,
	}
}

func (q *ListQueue) Len() int {
	return q.length
}

func (q *ListQueue) Enqueue(x interface{}) bool {
	if q.Len() == q.capacity {
		return false
	}

	node := &listNode{val: x}
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}

	q.length++

	return true
}

func (q *ListQueue) Dequeue() interface{} {
	if q.Len() == 0 {
		return nil
	}

	v := q.head.val
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.length--

	return v
}
