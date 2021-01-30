package diagonalSum

//https://leetcode-cn.com/problems/matrix-diagonal-sum/
func diagonalSum(mat [][]int) int {

	var n = len(mat)
	var res, mid = 0, n / 2
	if len(mat[0]) == 1 {
		return mat[0][0]
	}

	for i := 0; i < n; i++ {
		res += mat[i][i] + mat[i][n - 1 - i]
	}

	return res - mat[mid][mid] * (n & 1)
}