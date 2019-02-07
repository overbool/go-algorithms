package app

func TopK(a []int, k int) int {
	return findK(a, 0, len(a)-1, k)
}

func findK(a []int, start, end, k int) int {
	//if start == end {
	//	return a[start]
	//}
	p := partition(a, start, end)
	if p+1 == k {
		return a[p]
	} else if p+1 < k {
		return findK(a, p+1, end, k)
	} else {
		return findK(a, start, p-1, k)
	}
}

// from big to small
func partition(a []int, start, end int) int {
	pivot := a[start]
	i := start
	j := end
	for i < j {
		for i < j && a[j] <= pivot {
			j--
		}
		a[i] = a[j]

		for i < j && a[i] >= pivot {
			i++
		}
		a[j] = a[i]
	}

	a[i] = pivot
	return i
}
