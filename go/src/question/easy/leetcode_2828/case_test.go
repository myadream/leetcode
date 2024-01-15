package leetcode_2828

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
			Value:  []string{"afqcpzsx", "icenu"},
			Assist: "yi",
		}, TargetData: common.TDefault{
			Value: false,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"alice", "bob", "charlie"},
			Assist: "abc",
		}, TargetData: common.TDefault{
			Value: true,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"never", "gonna", "give", "up", "on", "you"},
			Assist: "ngguoy",
		}, TargetData: common.TDefault{
			Value: true,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []string{"an", "apple"},
			Assist: "a",
		}, TargetData: common.TDefault{
			Value: false,
		},
	})

	return dataSets
}

func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
	var handle []common.Processor[common.DCDefault, common.TDefault]

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			res := CaseOne(dataSet.Value.([]string), dataSet.Assist.(string))
			return assert.Equal(
				t,
				res,
				target.Value.(bool),
				fmt.Sprintf(
					"case one: dataSet: %v, assist: %v, target: %v, res: %v",
					dataSet.Value.([]string),
					dataSet.Assist.(string),
					target,
					res,
				),
			)
		},
	})

	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
			res := CaseTwo(dataSet.Value.([]string), dataSet.Assist.(string))
			return assert.Equal(
				t,
				res,
				target.Value.(bool),
				fmt.Sprintf(
					"case one: dataSet: %v, assist: %v, target: %v, res: %v",
					dataSet.Value.([]string),
					dataSet.Assist.(string),
					target,
					res,
				),
			)
		},
	})

	return handle
}

func TestLeetcode2828(t *testing.T) {
	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
	test.ProcessData(handle(t), dataSet())
}
