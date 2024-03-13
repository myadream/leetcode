package leetcode_2864

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		res := caseOne(data.SourceData.Value.(string))

		assert.Equal(
			t,
			res,
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				res,
			),
		)

	}
}

func TestCaseTwo(t *testing.T) {
	for _, data := range dataSet() {
		res := caseTwo(data.SourceData.Value.(string))

		assert.Equal(
			t,
			res,
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				res,
			),
		)

	}
}
