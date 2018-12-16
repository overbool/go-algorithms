package sort

func BubbleSort(a []int) {
  length := len(a)
  done := false
  for i := 0; i < length && done; i++ {
    done = false
    for j := 0; j < length - i - 1; j++ {
      if a[j] > a[j+1] {
        a[j], a[j+1] = a[j+1], a[j]
        done = true
      }
    }
  }
}
