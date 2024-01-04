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
			Value:  []byte{'e', 'e', 'e', 'k', 'q', 'q', 'q', 'v', 'v', 'y'},
			Assist: byte('e'),
		}, TargetData: common.TDefault{
			Value: byte('k'),
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []byte{'c', 'f', 'j'},
			Assist: byte('g'),
		}, TargetData: common.TDefault{
			Value: byte('j'),
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []byte{'c', 'f', 'j'},
			Assist: byte('j'),
		}, TargetData: common.TDefault{
			Value: byte('c'),
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []byte{'c', 'f', 'j'},
			Assist: byte('d'),
		}, TargetData: common.TDefault{
			Value: byte('f'),
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []byte{'c', 'f', 'j'},
			Assist: byte('a'),
		}, TargetData: common.TDefault{
			Value: byte('c'),
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []byte{'c', 'f', 'j'},
			Assist: byte('c'),
		}, TargetData: common.TDefault{
			Value: byte('f'),
		},
	})

	dataSets = append(dataSets, common.DataCarrier[common.DCDefault, common.TDefault]{
		SourceData: common.DCDefault{
			Value:  []byte{'x', 'x', 'y', 'y'},
			Assist: byte('z'),
		}, TargetData: common.TDefault{
			Value: byte('x'),
		},
	})

	return dataSets
}

func TestCaseOne(t *testing.T) {
	for _, data := range dataSet() {
		res := CaseOne(data.SourceData.Value.([]byte), data.SourceData.Assist.(byte))
		assert.Equal(
			t,
			res,
			data.TargetData.Value.(byte),
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
		res := CaseTwo(data.SourceData.Value.([]byte), data.SourceData.Assist.(byte))
		assert.Equal(
			t,
			res,
			data.TargetData.Value.(byte),
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				res,
			),
		)
	}
}

func TestCaseThree(t *testing.T) {
	for _, data := range dataSet() {
		res := CaseThree(data.SourceData.Value.([]byte), data.SourceData.Assist.(byte))
		assert.Equal(
			t,
			res,
			data.TargetData.Value.(byte),
			fmt.Sprintf(
				"case one: dataSet: %v, target: %v, res: %v",
				data,
				data.TargetData.Value,
				res,
			),
		)
	}
}
