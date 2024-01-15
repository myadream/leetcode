package leetcode_2085

import "math"

func caseOne(words1 []string, words2 []string) int {

	w1, w2 := len(words1), len(words2)

	var dp1 = make(map[string]uint8, w1)
	var dp2 = make(map[string]uint8, w2)

	for i := w1 - 1; i >= 0; i-- {
		dp1[words1[i]] += 1
	}

	for i := w2 - 1; i >= 0; i-- {
		dp2[words2[i]] += 1
	}

	sum := 0
	for key, num := range dp2 {
		if dp1[key] == 1 && num == 1 {
			sum += 1
		}
	}

	return sum
}

func caseTwo(words1 []string, words2 []string) int {

	if len(words1) > len(words2) {
		return caseTwo(words2, words1)
	}

	var m = make(map[string]uint8, len(words1))
	for i := len(words1) - 1; i >= 0; i-- {
		m[words1[i]] += 1
	}

	sum := 0
	for i := len(words2) - 1; i >= 0; i-- {
		if _, ok := m[words2[i]]; !ok {
			continue
		}

		if m[words2[i]] == 1 {
			sum += 1
			m[words2[i]] = math.MaxUint8
		} else if m[words2[i]] == math.MaxUint8 {
			sum -= 1
			m[words2[i]] = 0
		}
	}

	return sum
}
