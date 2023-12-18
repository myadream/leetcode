package template

import (
	"strconv"
)

func CaseOne(num1 string, num2 string) string {
	sum := 0
	total := ""
	for l, r := len(num1), len(num2); l > 0 || r > 0 || sum != 0; {
		if l > 0 {
			l--
			sum += int(num1[l] - '0')
		}

		if r > 0 {
			r--
			sum += int(num2[r] - '0')
		}

		total = strconv.Itoa(sum%10) + total
		sum /= 10
	}

	return total
}
