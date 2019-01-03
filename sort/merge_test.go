package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{3, 8, 9, 10, 2, 1, 23}
	MergeSort(a)
	fmt.Println(a)
}
