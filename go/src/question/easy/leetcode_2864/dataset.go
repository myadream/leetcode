package leetcode_2864

import "leetcode/src/common"

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "010",
			Assist: 0,
		}, TargetData: common.TDefault{
			Value: "001",
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "0101",
			Assist: 0,
		}, TargetData: common.TDefault{
			Value: "1001",
		},
	})

	return dataSets
}
