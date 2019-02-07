package stack

import "testing"

func TestListStack(t *testing.T) {
	s := NewListStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for i := 3; i > 0; i-- {
		item := s.Pop()
		if item != i {
			t.Errorf("expected %d, but %d", i, item)
		}
	}
}
