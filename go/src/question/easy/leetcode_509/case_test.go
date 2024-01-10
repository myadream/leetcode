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
			Value:  4,
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 3,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  3,
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 2,
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  2,
			Assist: nil,
		}, TargetData: common.TDefault{
			Value: 1,
		},
	})

	return dataSets
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		res := caseOne(data.SourceData.Value.(int))
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

//func handle(t *testing.T) []common.Processor[common.DCDefault, common.TDefault] {
//	var handle []common.Processor[common.DCDefault, common.TDefault]
//
//	handle = append(handle, common.Processor[common.DCDefault, common.TDefault]{
//		Fn: func(dataSet common.DCDefault, target common.TDefault) bool {
//			res := CaseOne(dataSet.Value.(string), dataSet.Assist.(string))
//			return assert.Equal(
//				t,
//				res,
//				target.Value.(string),
//				fmt.Sprintf(
//					"case one: dataSet: %v, assist: %v, target: %v, res: %v",
//					dataSet.Value.(string),
//					dataSet.Assist.(string),
//					target,
//					res,
//				),
//			)
//		},
//	})
//
//	return handle
//}
//
//func Test(t *testing.T) {
//	test := common.DataFlowProcessor[common.DCDefault, common.TDefault]{}
//	test.ProcessData(handle(t), dataSet())
//}
