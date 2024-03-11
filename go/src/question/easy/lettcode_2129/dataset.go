package lettcode_2129

import "leetcode/src/common"

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "capiTalIze tHe titLe",
			Assist: 0,
		}, TargetData: common.TDefault{
			Value: "Capitalize The Title",
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "First leTTeR of EACH Word",
			Assist: 0,
		}, TargetData: common.TDefault{
			Value: "First Letter of Each Word",
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "i lOve leetcode",
			Assist: 0,
		}, TargetData: common.TDefault{
			Value: "i Love Leetcode",
		},
	})

	return dataSets
}
