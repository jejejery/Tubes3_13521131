package main

func BMMatch(text string, pattern string) int {
	last := buildLast(pattern)
	n := len(text)
	m := len(pattern)
	i := (m - 1)
	if i > (n - 1) {
		return -1
	}
	j := (m - 1)

	for ok := true; ok; ok = (i <= n-1) {
		if pattern[j] == text[i] {
			if j == 0 {
				return i
			} else {
				i--
				j--
			}
		} else {
			lo := last[text[i]]
			if j < 1+lo {
				i = (i + m - j)
			} else {
				i = (i + m - (1 + lo))
			}
			j = m - 1

		}
	}
	return -1

}

func buildLast(pattern string) []int {
	var last [128]int
	for i := 0; i < 128; i++ {
		last[i] = -1
	}
	for j := 0; j < len(pattern); j++ {
		last[pattern[j]] = j
	}
	return last[:128]

}
