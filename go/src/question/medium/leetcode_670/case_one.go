package template

import "strconv"

func caseOne(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if s[i] >= s[j] {
				continue
			}

			s[i], s[j] = s[j], s[i]
			t, _ := strconv.Atoi(string(s))
			num = max(num, t)

			s[i], s[j] = s[j], s[i]
		}

	}

	return num
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func caseTwo(num int) int {
	s := []byte(strconv.Itoa(num))
	n := len(s)

	maxIdx, idx1, idx2 := n-1, -1, -1

	for i := n - 2; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else if s[i] < s[maxIdx] {
			idx1, idx2 = i, maxIdx
		}
	}

	if idx1 >= 0 {
		s[idx1], s[idx2] = s[idx2], s[idx1]
		num, _ = strconv.Atoi(string(s))
	}

	return num
}
