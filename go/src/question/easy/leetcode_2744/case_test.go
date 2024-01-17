package leetcode_2744

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"wk", "xf", "ot", "je", "hd", "kw", "fx", "to", "ej"},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 4,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"cd", "ac", "dc", "ca", "zz"},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 2,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"ab", "ba", "cc"},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 1,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"aa", "ab"},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 0,
		},
	})

	return dataSets
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		res := caseOne(data.SourceData.Value.([]string))
		assert.Equal(
			t,
			res,
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				res,
			),
		)
	}
}

func TestCaseTwo(t *testing.T) {
	for _, data := range dataSet() {
		res := caseTwo(data.SourceData.Value.([]string))
		assert.Equal(
			t,
			res,
			data.TargetData.Value,
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				res,
			),
		)
	}
}
