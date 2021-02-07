package equalSubstring

func equalSubstring(s string, t string, maxCost int) int {
	var n = len(t)
	var record = make([]int8, n)

	for k, _ := range s {
		record[k] = abs(int8(t[k] - s[k]))
	}

	var r, l int
	var usedCost, res int8

	for r < n {
		usedCost += record[r]

		for int(usedCost) > maxCost {
			usedCost -= record[l]
			l++
		}

		res = max(res, int8(r - l + 1))
		r++
	}

	return int(res)
}

func max(x, y int8) int8 {
	if x > y {
		return x
	}

	return y
}

func abs(x int8) int8 {
	if x < 0 {
		return ^x + 1
	}

	return x
}