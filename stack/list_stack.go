package stack

type listNode struct {
	next *listNode
	val  interface{}
}

type ListStack struct {
	head *listNode
}

func NewListStack() *ListStack {
	return &ListStack{
		head: &listNode{},
	}
}

func (s *ListStack) IsEmpty() bool {
	return s.head == nil
}

func (s *ListStack) Push(x interface{}) {
	s.head = &listNode{
		next: s.head,
		val:  x,
	}
}

func (s *ListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	res := s.head.val
	s.head = s.head.next

	return res
}

func (s *ListStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.head.val
}
