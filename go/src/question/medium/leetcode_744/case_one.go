package leetcode_744

func CaseOne(letters []byte, target byte) byte {
	res := letters[0]

	for _, word := range letters {
		if word <= target {
			continue
		}

		if word < res || res <= target {
			res = word
		}
	}

	return res
}

func CaseTwo(letters []byte, target byte) byte {
	l, r, mid := 0, len(letters)-1, 0
	res := letters[0]

	for l <= r {
		mid = l + ((r - l) >> 1)
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1

			if letters[mid] < res || res == letters[0] {
				res = letters[mid]
			}
		}
	}

	return res
}

func CaseThree(letters []byte, target byte) byte {
	r := len(letters)

	for l := 0; l < r; {
		mid := l + ((r - l) >> 1)
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if r == len(letters) {
		return letters[0]
	}

	return letters[r]
}
