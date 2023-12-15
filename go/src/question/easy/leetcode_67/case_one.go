package leetcode_67

import (
	"strconv"
)

func CaseOne(a string, b string) string {
	al := len(a)
	bl := len(b)

	sum := 0
	ans := ""
	for al > 0 || bl > 0 {
		if al > 0 {
			al--
			sum += int(a[al] - '0')
		}

		if bl > 0 {
			bl--
			sum += int(b[bl] - '0')
		}

		ans = strconv.Itoa(sum%2) + ans
		sum /= 2
	}

	if sum == 1 {
		ans = "1" + ans
	}

	return ans
}
