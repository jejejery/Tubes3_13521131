package algorithm

func min3Int(x int, y int, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	} else {
		return z
	}
}

func MinimumDist(s, t string) int {
	m := len(s)
	n := len(t)

	// initialize the distance matrix
	d := make([][]int, m+1)
	for i := range d {
		d[i] = make([]int, n+1)
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}

	// fill the distance matrix
	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min3Int(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+1)
			}
		}
	}

	return d[m][n]

}

func similarityPercentage(text string, pattern string) float32 {
	distance := MinimumDist(text, pattern)
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
