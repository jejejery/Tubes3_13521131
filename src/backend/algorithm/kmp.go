package algorithm

func KMPMatch(text string, pattern string) (int, float32) {
	n := len(text)
	m := len(pattern)

	b := calculateBorder(pattern)
	i := 0
	j := 0

	for i < n {
		if (pattern[j] == text[i]) {
			if (j == m - 1) {
				return i - m + 1, similarityPercentage(text, pattern)
			}
			i++
			j++
		} else if j > 0 {
			j = b[j-1]
		} else {
			i++
		}
	}
	return -1, similarityPercentage(text, pattern)
}

func calculateBorder(pattern string) []int {
	var b [1000]int
	m := len(pattern)
	j := 0
	i := 1
	k := m-1

	for i < m {
		if pattern[j] == pattern[i] { // chars match
			b[i] = j + 1
			j++
			i++
		} else if j > 0 {
			j = b[j-1]
		} else {
			b[i] = 0
			i++
		}
	}

	return b[:k]
}