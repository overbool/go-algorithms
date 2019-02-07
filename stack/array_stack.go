package stack

type ArrayStack struct {
	data []interface{}
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
	}
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *ArrayStack) Len() int {
	return len(s.data)
}

func (s *ArrayStack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return res
}

func (s *ArrayStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.data[len(s.data)-1]
}
