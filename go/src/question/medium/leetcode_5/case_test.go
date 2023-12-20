package template

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
			Value: "babad",
		}, TargetData: common.TDefault{
			Value: "aba",
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value: "cbbd",
		}, TargetData: common.TDefault{
			Value: "bb",
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			res := CaseOne(dataSet.Value.(string))
			return assert.Equal(
				t,
				res,
				target.Value.(string),
				fmt.Sprintf(
					"case one: dataSet: %v, target: %v, res: %v",
					dataSet.Value.(string),
					target,
					res,
				),
			)
		},
	})

	return handle
}

func Test(t *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(t), dataSet())
}
