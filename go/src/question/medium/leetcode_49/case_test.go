package leetcode_49

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"leetcode/src/common"
	"testing"
)

func dataSet() []common.DataCarrier[common.DCDefault, common.TDefault] {
	var dataSets []common.DataCarrier[common.DCDefault, common.TDefault]

	//dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
	//	SourceData: common.DCDefault{
	//		Value:  []string{"eat", "tea", "tan", "ate", "nat", "bat"},
	//		Assist: nil,
	//	}, TargetData: common.TDefault{
	//		Value: [][]string{[]string{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}},
	//	},
	//})

	//dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
	//	SourceData: common.DCDefault{
	//		Value:  []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"},
	//		Assist: nil,
	//	}, TargetData: common.TDefault{
	//		Value: [][]string{{"max"}, {"buy"}, {"doc"}, {"may"}, {"ill"}, {"duh"}, {"tin"}, {"bar"}, {"pew"}, {"cab"}},
	//	},
	//})
	//
	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{""},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: [][]string{{""}},
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"a"},
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: [][]string{{"a"}},
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
		res := caseOne(data.SourceData.Value.([]string))
		assert.EqualValuesf(
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
