package sort

import (
  "testing"
  "fmt"
)

func TestBubbleSort(t *testing.T) {
  a := []int{3, 8, 9, 10, 2, 1, 23}
  BubbleSort(a)
  fmt.Println(a)
}
