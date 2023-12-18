package leetcode_989

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
			Value:  []int{1, 2, 0, 0},
			Assist: 34,
		}, TargetData: common.TDefault{
			Value: []int{1, 2, 3, 4},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{2, 7, 4},
			Assist: 181,
		}, TargetData: common.TDefault{
			Value: []int{4, 5, 5},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{2, 1, 5},
			Assist: 806,
		}, TargetData: common.TDefault{
			Value: []int{1, 0, 2, 1},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			res := CaseOne(dataSet.Value.([]int), dataSet.Assist.(int))
			return assert.Equal(
				t,
				res,
				target.Value.([]int),
				fmt.Sprintf(
					"case one: dataSet: %v, assist: %v, target: %v, res: %v",
					dataSet.Value.([]int),
					dataSet.Assist.(int),
					target,
					res,
				),
			)
		},
	})

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			res := CaseTwo(dataSet.Value.([]int), dataSet.Assist.(int))
			return assert.Equal(
				t,
				res,
				target.Value.([]int),
				fmt.Sprintf(
					"case one: dataSet: %v, assist: %v, target: %v, res: %v",
					dataSet.Value.([]int),
					dataSet.Assist.(int),
					target,
					res,
				),
			)
		},
	})

	return handle
}

func TestLeetCode989(t *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(t), dataSet())
}
