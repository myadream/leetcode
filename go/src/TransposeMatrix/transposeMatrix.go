package TransposeMatrix

func transpose(A [][]int) [][]int {
	var result [][]int
	var i, j int

	for i = 0; i <= len(A); i++ {
		for j = 0; j <= len(A[0]); j++ {
			result[j][i] = A[i][j]
		}
	}


	return result
}