package sort

func SelectionSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}

			if min != i {
				a[i], a[min] = a[min], a[i]
			}
		}
	}
}
