package sort

import (
  "testing"
  "fmt"
)

func TestQuickSort(t *testing.T) {
  a := []int{3, 8, 9, 10, 2, 1, 23}
  QuickSort(a)
  fmt.Println(a)
}