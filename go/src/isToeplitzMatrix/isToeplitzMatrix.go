package isToeplitzMatrix


//https://leetcode-cn.com/problems/toeplitz-matrix/
func isToeplitzMatrix(matrix [][]int) bool {
	var r = len(matrix)
	var c = len(matrix[0])

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][j] != matrix[i - 1][j - 1] {
				return false
			}
		}
	}

	return true
}
