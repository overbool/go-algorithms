package queue

type Interface interface {
	Enqueue(interface{})
	Dequeue() interface{}
}
