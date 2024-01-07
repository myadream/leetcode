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
			Value:  []int{10, 6, 8, 5, 11, 9},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: []int{3, 1, 2, 1, 1, 0},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{5, 1, 2, 3, 10},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: []int{4, 1, 1, 1, 0},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			res := CaseOne(dataSet.Value.([]int))
			return assert.Equal(
				t,
				res,
				target.Value.([]int),
				fmt.Sprintf(
					"case one: dataSet: %v, assist: %v, target: %v, res: %v",
					dataSet.Value,
					dataSet.Assist,
					target,
					res,
				),
			)
		},
	})

	return handle
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		res := CaseOne(data.SourceData.Value.([]int))
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
