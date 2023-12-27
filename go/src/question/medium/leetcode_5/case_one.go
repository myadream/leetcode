package template

import "math"

func CaseOne(s string) string {
	if len(s) <= 0 || s == "" {
		return ""
	}

	var start, end = 0, 0
	for i := 0; i <= len(s); i++ {
		l := expandAroundCenter(s, i, i)
		r := expandAroundCenter(s, i, i+1)

		strlen := math.Max()

		if strlen > end-start {
			start = i - (strlen-1)/2
			end = i + strlen/2
		}
	}

	return s[start:end]
}

func expandAroundCenter(s string, start, end int) (int, int) {

}
