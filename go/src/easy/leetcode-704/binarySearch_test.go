package leetcode704

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
			Value:  []int{-1, 0, 3, 5, 9, 12},
			Assist: 9,
		},
		Target: common.DataSetTarget{
			Value: 4,
		},
	})

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value:  []int{-1, 0, 3, 5, 9, 12},
			Assist: 2,
		},
		Target: common.DataSetTarget{
			Value: -1,
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget] {
	var handle []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]

	handle = append(handle, common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]{
		HandlerDataSet: func(dataSet common.DataSetParamAssist, target common.DataSetTarget) bool {

			res := defaultSearch(dataSet.Value.([]int), dataSet.Assist.(int))
			assert.Equal(t, res, target.Value.(int), fmt.Sprintf("result: dataSet: %v, target: %v, res: %v", dataSet, target, res))

			return true
		},
	})

	handle = append(handle, common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]{
		HandlerDataSet: func(dataSet common.DataSetParamAssist, target common.DataSetTarget) bool {

			res := binarySearch(dataSet.Value.([]int), dataSet.Assist.(int))
			assert.Equal(t, res, target.Value.(int), fmt.Sprintf("result: dataSet: %v, target: %v, res: %v", dataSet, target, res))

			return true
		},
	})

	return handle
}

func TestBinarySearch(t *testing.T) {
	test := common.DataSetControl[common.DataSetParamAssist, common.DataSetTarget]{}
	test.Run(dataSet(), handle(t))
}
