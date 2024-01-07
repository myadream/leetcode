package leetcode_383

import (
	"leetcode/src/common"
)

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "aa",
			Assist: "aab",
		}, TargetData: common.TDefault{
			Value: true,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "aa",
			Assist: "ab",
		}, TargetData: common.TDefault{
			Value: false,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "a",
			Assist: "b",
		}, TargetData: common.TDefault{
			Value: false,
		},
	})

	return dataSets
}
