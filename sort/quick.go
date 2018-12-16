package sort

func QuickSort(a []int) {
  if len(a) <= 1 {
    return
  }
  pivot := partition(a)
  QuickSort(a[:pivot-1])
  QuickSort(a[pivot+1:])
}

func partition(a []int) int {
  pivot := a[0]
  low := 0
  high := len(a) - 1
  for low < high {
    for low < high && a[high] >= pivot {
      high--
    }
    a[low] = a[high]

    for low < high && a[low] <= pivot {
      low++
    }
    a[high] = a[low]
  }

  a[low] = pivot
  return low
}
