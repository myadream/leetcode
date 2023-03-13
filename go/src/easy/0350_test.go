package easy

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func dataSet() []common.DataSet[common.DataSetParamAssist, common.DataSetTarget] {
	var dataSets []common.DataSet[common.DataSetParamAssist, common.DataSetTarget]

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value:  []int{1, 2, 2, 1},
			Assist: []int{2, 2},
		},
		Target: common.DataSetTarget{
			Value: []int{2, 2},
		},
	})

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value:  []int{1, 2, 2},
			Assist: []int{9, 4, 9, 8, 4},
		},
		Target: common.DataSetTarget{
			Value: []int{9, 4},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget] {
	var handle []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]

	handle = append(handle, common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]{
		HandlerDataSet: func(dataSet common.DataSetParamAssist, target common.DataSetTarget) bool {

			res := intersect(dataSet.Value.([]int), dataSet.Assist.([]int))
			assert.Equal(t, res, target.Value.([]int), fmt.Sprintf("two sum: dataSet: %v, target: %v, res: %v", dataSet, target, res))

			return true
		},
	})

	return handle
}

func TestIntersect(t *testing.T) {
	test := common.DataSetControl[common.DataSetParamAssist, common.DataSetTarget]{}
	test.Run(dataSet(), handle(t))
}
