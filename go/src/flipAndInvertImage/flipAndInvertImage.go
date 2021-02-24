package flipAndInvertImage

//https://leetcode-cn.com/problems/flipping-an-image/
func flipAndInvertImage(A [][]int) [][]int {

	var left, right, n = 0, 0, len(A[0]) - 1
	for _, row := range A {
		left, right = 0, n

		for left < right {
			if row[left] == row[right] {
				row[left] ^= 1
				row[right] ^= 1
			}

			left++
			right--
		}

		if left == right {
			row[left] ^= 1
		}

	}

	return A
}