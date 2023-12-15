package leetcode_67

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
			Value:  "11",
			Assist: "1",
		}, TargetData: common.TDefault{
			Value: "100",
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "1010",
			Assist: "1011",
		}, TargetData: common.TDefault{
			Value: "10101",
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  "110",
			Assist: "1",
		}, TargetData: common.TDefault{
			Value: "111",
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			res := CaseOne(dataSet.Value.(string), dataSet.Assist.(string))
			return assert.Equal(
				t,
				res,
				target.Value.(string),
				fmt.Sprintf(
					"case one: dataSet: %s, assist: %s, target: %s, res: %s",
					dataSet.Value.(string),
					dataSet.Assist.(string),
					target,
					res,
				),
			)
		},
	})

	return handle
}

func TestLeetcode67(t *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(t), dataSet())
}
