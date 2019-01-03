package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	a := []int{3, 8, 9, 10, 2, 1, 23}
	InsertionSort(a)
	fmt.Println(a)
}
