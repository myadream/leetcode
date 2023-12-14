package leetcode704

import (
	"fmt"
	"leetcode/src/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{-1, 0, 3, 5, 9, 12},
			Assist: 9,
		}, TargetData: common.TDefault{
			Value: 4,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []int{-1, 0, 3, 5, 9, 12},
			Assist: 2,
		}, TargetData: common.TDefault{
			Value: -1,
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {

			res := defaultSearch(dataSet.Value.([]int), dataSet.Assist.(int))
			assert.Equal(t, res, target.Value.(int), fmt.Sprintf("result: dataSet: %v, target: %v, res: %v", dataSet, target, res))

			return true
		},
	})

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {

			res := binarySearch(dataSet.Value.([]int), dataSet.Assist.(int))
			assert.Equal(t, res, target.Value.(int), fmt.Sprintf("result: dataSet: %v, target: %v, res: %v", dataSet, target, res))

			return true
		},
	})

	return handle
}

func TestBinarySearch(t *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(t), dataSet())
}
