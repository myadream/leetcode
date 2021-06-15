package numMatrix

type NumMatrix struct {
	nums [][]int
}


func Constructor(matrix [][]int) (nm NumMatrix) {
	row := len(matrix)

	if row == 0 {
		return NumMatrix{}
	}

	col := len(matrix[0])

	nm.nums = make([][]int, row + 1)
	nm.nums[0] = make([]int, col + 1)
	for i := 0; i < row; i++ {
		nm.nums[i + 1] = make([]int, col + 1)
		for j := 0; j < col; j++ {
			nm.nums[i + 1][j + 1] = nm.nums[i + 1][j] + nm.nums[i][j + 1] - nm.nums[i][j] + matrix[i][j]
		}
	}

	return nm
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.nums[row2 + 1][col2 + 1] - this.nums[row1][col2 + 1] - this.nums[row2 + 1][col1] + this.nums[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */