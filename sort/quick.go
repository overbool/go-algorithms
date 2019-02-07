package sort

func QuickSort(a []int) {
	if len(a) < 2 {
		return
	}

	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, start, end int) {
	if start >= end {
		return
	}

	p := partition(a, start, end)
	quickSort(a, start, p-1)
	quickSort(a, p+1, end)
}

func partition(a []int, start, end int) int {
	// select the first num as pivot
	pivot := a[start]
	i := start
	j := end
	for i < j {
		for i < j && a[j] >= pivot {
			j--
		}
		a[i] = a[j]

		for i < j && a[i] <= pivot {
			i++
		}
		a[j] = a[i]
	}

	a[i] = pivot
	return i
}
