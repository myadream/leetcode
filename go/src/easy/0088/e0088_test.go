package _088

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func dataSet() []common.DataSet[common.DataSetParamAssist, common.DataSetTarget] {
	var dataSets []common.DataSet[common.DataSetParamAssist, common.DataSetTarget]

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value: [][]int{
				{0, 0, 0, 0, 0},
				{1, 2, 3, 4, 5},
			},
			Assist: []int{0, 5},
		},
		Target: common.DataSetTarget{
			Value: []int{1, 2, 3, 4, 5},
		},
	})

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value: [][]int{
				{1, 2, 3, 0, 0, 0},
				{2, 5, 6},
			},
			Assist: []int{3, 3},
		},
		Target: common.DataSetTarget{
			Value: []int{1, 2, 2, 3, 5, 6},
		},
	})

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value: [][]int{
				{1},
				{0},
			},
			Assist: []int{1, 0},
		},
		Target: common.DataSetTarget{
			Value: []int{1},
		},
	})

	dataSets = append(dataSets, common.DataSet[common.DataSetParamAssist, common.DataSetTarget]{
		Param: common.DataSetParamAssist{
			Value: [][]int{
				{0},
				{1},
			},
			Assist: []int{0, 1},
		},
		Target: common.DataSetTarget{
			Value: []int{1},
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget] {
	var handle []common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]

	handle = append(handle, common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]{
		HandlerDataSet: func(dataSet common.DataSetParamAssist, target common.DataSetTarget) bool {
			var p1, p2 []int
			copier.Copy(&p1, &dataSet.Value.([][]int)[0])
			copier.Copy(&p2, &dataSet.Value.([][]int)[1])

			merge2(p1, dataSet.Assist.([]int)[0], p2, dataSet.Assist.([]int)[1])
			assert.Equal(t, p1, target.Value.([]int), fmt.Sprintf("merge: dataSet: %v, target: %v, res: %v", dataSet, target, p1))

			return true
		},
	})

	handle = append(handle, common.FunDataSet[common.DataSetParamAssist, common.DataSetTarget]{
		HandlerDataSet: func(dataSet common.DataSetParamAssist, target common.DataSetTarget) bool {
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
	test := common.DataSetControl[common.DataSetParamAssist, common.DataSetTarget]{}
	test.Run(dataSet(), handle(Test))
}
