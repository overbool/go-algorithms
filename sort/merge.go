package sort

// O(nlogn)
// space: O(n)
func MergeSort(a []int) {
	if len(a) < 2 {
		return
	}

	mergeSort(a, 0, len(a)-1)
}

func mergeSort(a []int, start, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)

	merge(a, start, mid, end)
}

func merge(a []int, start, mid, end int) {
	tmp := make([]int, end-start+1)

	i := start
	j := mid + 1
	k := 0

	for ; i <= mid && j <= end; k++ {
		if a[i] < a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp[k] = a[i]
		k++
	}

	for ; j <= end; j++ {
		tmp[k] = a[j]
		k++
	}

	copy(a[start:end+1], tmp)
}
