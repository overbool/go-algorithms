package sort

func InsertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		value := a[i]
		var j int
		for j = i - 1; j >= 0 && a[j] > value; j-- {
			a[j+1] = a[j]
		}

		a[j+1] = value
	}
}
