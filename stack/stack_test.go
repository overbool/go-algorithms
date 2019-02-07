package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()
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
