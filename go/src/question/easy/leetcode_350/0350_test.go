package leetcode_350

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
			Value:  []int{1, 2, 2, 1},
			Assist: []int{2, 2},
		}, TargetData: common.TDefault{
			Value: []int{2, 2},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{4, 9, 5},
			Assist: []int{9, 4, 9, 8, 4},
		}, TargetData: common.TDefault{
			Value: []int{9, 4},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {

			res := intersect(dataSet.Value.([]int), dataSet.Assist.([]int))
			assert.Equal(t, res, target.Value.([]int), fmt.Sprintf("two sum: dataSet: %v, target: %v, res: %v", dataSet, target, res))

			return true
		},
	})

	return handle
}

func TestIntersect(t *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(t), dataSet())
}
