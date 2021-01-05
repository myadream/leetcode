package TransposeMatrix

import (
	"testing"
)

func TestTranspose(t *testing.T) {
	var successResult, target [][]int
	var i, j int

	target = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	successResult = [][]int{
		{1, 4},
		{2, 5},
		{3, 6},
	}

	result := transpose(target)

	for i = 0; i <= len(successResult); i++ {
		for j = 0; j <= len(successResult[0]); j++ {
			if successResult[i][j] != result[i][j] {
				t.Fatal(result)
			}
		}
	}
}