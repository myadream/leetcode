package reverseLeftWords

//https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func reverseLeftWords(s string, n int) string {
	var l = len(s)

	var ans = []byte(s)
	var res = make([]byte, l)

	for k, v := range ans {
		if k < n {
			res[(l - n - 1) + (k + 1)] = v
		} else {
			res[k - n] = v
		}
	}

	return string(res)
}
