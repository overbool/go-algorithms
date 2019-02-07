package stack

type Interface interface {
	IsEmpty() bool
	Push(interface{})
	Pop() interface{}
	Top() interface{}
}
