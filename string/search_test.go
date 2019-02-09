package string

import (
	"testing"
)

func TestSearchByBF(t *testing.T) {
	str := "abcd88fbi"
	pattern := "fb"
	idx := SearchByBF(str, pattern)

	if idx != 6 {
		t.Errorf("expected %d, but %d", 6, idx)
	}
}
