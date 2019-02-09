package string

// SearchByBF searches pattern in str using brute force.
// it will return the first matching index.
func SearchByBF(str string, pattern string) int {
	if len(str) == 0 || len(pattern) == 0 || len(str) < len(pattern) {
		return -1
	}

	for i := 0; i <= len(str)-len(pattern); i++ {
		sub := str[i : i+len(pattern)]
		if sub == pattern {
			return i
		}
	}

	return -1
}
