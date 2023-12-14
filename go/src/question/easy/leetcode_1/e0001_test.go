package leetcode_1

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
			Value:  []int{2, 7, 11, 15},
			Assist: 9,
		},
		TargetData: common.TDefault{
			Value: []int{0, 1},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{3, 2, 4},
			Assist: 6,
		},
		TargetData: common.TDefault{Value: []int{1, 2}},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{3, 3},
			Assist: 6,
		},
		TargetData: common.TDefault{
			Value: []int{0, 1},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(
		handle,
		common.Processor[common.DCDefault, common.TDefault]{
			Fn: func(sourceData common.DCDefault, targetData common.TDefault) bool {

				res := twoSum(sourceData.Value.([]int), sourceData.Assist.(int))
				if assert.Equal(t, res, targetData.Value.([]int), fmt.Sprintf("two sum: dataSet: %v, target: %v, res: %v", sourceData, targetData, res)) {
					return true
				}

				return false
			},
		},
	)

	return handle
}

func TestTwoSum(Test *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(Test), dataSet())
}
