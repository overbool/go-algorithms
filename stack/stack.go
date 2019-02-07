package stack

import "container/list"

type Stack struct {
	data *list.List
}

func New() *Stack {
	return &Stack{
		data: list.New(),
	}
}

func (s *Stack) Len() int {
	return s.data.Len()
}

func (s *Stack) IsEmpty() bool {
	return s.data.Len() == 0
}

func (s *Stack) Push(x interface{}) {
	s.data.PushBack(x)
}

func (s *Stack) Pop() interface{} {
	if s.data.Len() == 0 {
		return nil
	} else {
		res := s.data.Back().Value
		s.data.Remove(s.data.Back())
		return res
	}
}

func (s *Stack) Top() interface{} {
	if s.data.Len() == 0 {
		return nil
	} else {
		return s.data.Back().Value
	}
}
