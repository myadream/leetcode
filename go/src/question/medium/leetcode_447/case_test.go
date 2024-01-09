package leetcode_447

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
			Value:  [][]int{{1, 1}},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 0,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  [][]int{{1, 1}, {2, 2}, {3, 3}},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 2,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  [][]int{{0, 0}, {1, 0}, {2, 0}},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 2,
		},
	})

	return dataSets
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		val := data.SourceData.Value.([][]int)
		res := CaseOne(val)
		assert.Equal(
			t,
			res,
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				val,
				data.TargetData.Value,
				res,
			),
		)
	}
}
