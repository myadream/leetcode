package numMatrix

import (
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestNumMatrix(t *testing.T) {
	var obj = Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})


	assert.Equal(t, obj.SumRegion(2, 1, 4, 3), 8, "error")

	assert.Equal(t, 11, obj.SumRegion(1, 1, 2, 2), "error")

	assert.Equal(t, 12, obj.SumRegion(1, 2, 2, 4), "error")
}