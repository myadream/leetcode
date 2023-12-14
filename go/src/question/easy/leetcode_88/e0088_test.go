package leetcode_88

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value: [][]int{
				{0, 0, 0, 0, 0},
				{1, 2, 3, 4, 5},
			},
			Assist: []int{0, 5},
		}, TargetData: common.TDefault{
			Value: []int{1, 2, 3, 4, 5},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value: [][]int{
				{1, 2, 3, 0, 0, 0},
				{2, 5, 6},
			},
			Assist: []int{3, 3},
		}, TargetData: common.TDefault{
			Value: []int{1, 2, 2, 3, 5, 6},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value: [][]int{
				{1},
				{0},
			},
			Assist: []int{1, 0},
		}, TargetData: common.TDefault{
			Value: []int{1},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value: [][]int{
				{0},
				{1},
			},
			Assist: []int{0, 1},
		}, TargetData: common.TDefault{
			Value: []int{1},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			var p1, p2 []int
			copier.Copy(&p1, &dataSet.Value.([][]int)[0])
			copier.Copy(&p2, &dataSet.Value.([][]int)[1])

			merge2(p1, dataSet.Assist.([]int)[0], p2, dataSet.Assist.([]int)[1])
			assert.Equal(t, p1, target.Value.([]int), fmt.Sprintf("merge: dataSet: %v, target: %v, res: %v", dataSet, target, p1))

			return true
		},
	})

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			var p1, p2 []int
			copier.Copy(&p1, &dataSet.Value.([][]int)[0])
			copier.Copy(&p2, &dataSet.Value.([][]int)[1])

			merge(p1, dataSet.Assist.([]int)[0], p2, dataSet.Assist.([]int)[1])
			assert.Equal(t, p1, target.Value.([]int), fmt.Sprintf("merge: dataSet: %v, target: %v, res: %v", dataSet, target, p1))

			return true
		},
	})

	return handle
}

func TestMerge(Test *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(Test), dataSet())
}
