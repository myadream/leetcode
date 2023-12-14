package leetcode_1

import (
	"fmt"
	"leetcode/src/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dataSet() []common.DataSet[common.DataSetParamAssist, common.DataSetTarget] {
	var dataSets []common.DataSet[common.DataSetParamAssist, common.DataSetTarget]

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value:  []int{2, 7, 11, 15},
			Assist: 9,
		},
		Target: common.DataSetTarget{
			Value: []int{0, 1},
		},
	})

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value:  []int{3, 2, 4},
			Assist: 6,
		},
		Target: common.DataSetTarget{
			Value: []int{1, 2},
		},
	})

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value:  []int{3, 3},
			Assist: 6,
		},
		Target: common.DataSetTarget{
			Value: []int{0, 1},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget] {
	var handle []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]

	handle = append(handle, common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]{
		HandlerDataSet: func(dataSet common.DataSetParamAssist, target common.DataSetTarget) bool {

			res := twoSum(dataSet.Value.([]int), dataSet.Assist.(int))
			assert.Equal(t, res, target.Value.([]int), fmt.Sprintf("two sum: dataSet: %v, target: %v, res: %v", dataSet, target, res))

			return true
		},
	})

	return handle
}

func TestTwoSum(Test *testing.T) {
	test := common.DataSetControl[common.DataSetParamAssist, common.DataSetTarget]{}
	test.Run(dataSet(), handle(Test))
}
