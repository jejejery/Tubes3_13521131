package algorithm

func minInt(x int, y int, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	} else {
		return z
	}
}
func Distance(text string, pattern string) int {
	return MinimumDist(0, 0, text, pattern)
}

func MinimumDist(idx1 int, idx2 int, text string, pattern string) int {
	if idx1 == len(text) && idx2 == len(pattern) {
		return (len(text) - idx1 + len(pattern) - idx2)
	}

	if text[idx1] == pattern[idx2] {
		return MinimumDist((idx1 + 1), (idx2 + 1), text, pattern)
	} else {
		return 1 + (minInt(MinimumDist(idx1, idx2+1, text, pattern), MinimumDist(idx1+1, idx2, text, pattern), MinimumDist(idx1+1, idx2+2, text, pattern)))
	}
}

func similarityPercentage(text string, pattern string) float32 {
	distance := Distance(text, pattern)
	lenText := len(text)
	lenPattern := len(pattern)

	var changePercentage float32
	if lenPattern < lenText {
		changePercentage = (float32(distance) / float32(lenText) * 100)
	} else {
		changePercentage = (float32(distance) / float32(lenText) * 100)
	}
	return (100 - changePercentage)
}
