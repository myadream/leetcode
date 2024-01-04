package template

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}},
			Assist: 2,
		}, TargetData: common.TDefault{
			Value: 3,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  [][]int{{1}, {0}},
			Assist: 1,
		}, TargetData: common.TDefault{
			Value: 2,
		},
	})

	return dataSets
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		res := CaseOne(data.SourceData.Value.([][]int), data.SourceData.Assist.(int))
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
