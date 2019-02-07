package app

import (
	"fmt"
	"testing"
)

func TestTopK(t *testing.T) {
	a := []int{2, 3, 8, 5, 6, 7, 1, 0, 9}
	fmt.Println("answer: ", TopK(a, 9))
}
