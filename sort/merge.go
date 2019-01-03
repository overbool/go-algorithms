package sort

func MergeSort(a []int) {
	if len(a) < 2 {
		return
	}

	mid := len(a) / 2
	MergeSort(a[:mid])
	MergeSort(a[mid:])

	if a[mid-1] <= a[mid] {
		return
	}

	s := make([]int, mid)
	copy(s, a[:mid])

	left := 0
	right := mid

	for i := 0; ; i++ {
		if s[left] <= a[right] {
			a[i] = s[left]
			left++
			if left == mid {
				break
			}
		} else {
			a[i] = a[right]
			right++
			if right == len(a) {
				copy(a[i+1:], s[left:mid])
				break
			}
		}
	}
}
