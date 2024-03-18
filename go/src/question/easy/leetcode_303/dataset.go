package leetcode_303

import "leetcode/src/common"

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"NumArray", "sumRange", "sumRange", "sumRange"},
			Assist: [][]int{{-2, 0, 3, -5, 2, -1}, {0, 2}, {2, 5}, {0, 5}},
		}, TargetData: common.TDefault{
			Value: []int{0, 1, -1, -3},
		},
	})

	return dataSets
}
