package leetcode_49

import "sort"

func caseOne(strs []string) [][]string {

	m := make(map[[26]byte][]string)
	for i := len(strs) - 1; i >= 0; i-- {
		sum := [26]byte{}
		for j := len(strs[i]) - 1; j >= 0; j-- {
			sum[strs[i][j]-'a']++
		}

		m[sum] = append(m[sum], strs[i])
	}

	ans := make([][]string, 0)
	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

func caseTwo(strs []string) [][]string {

	var m = make(map[string][]string)
	for i := len(strs) - 1; i >= 0; i-- {
		key := []byte(strs[i])
		sort.Slice(key, func(i, j int) bool { return key[i] < key[j] })

		m[string(key)] = append(m[string(key)], strs[i])
	}

	ans := make([][]string, 0)
	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}
