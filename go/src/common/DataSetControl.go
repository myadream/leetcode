package common

import (
	"github.com/jinzhu/copier"
)

type HandlerDataSet[P DataSetParamInterface, T DataSetInterface] func(dataSet P, target T) bool

type DataSet[P DataSetParamInterface, T DataSetInterface] struct {
	Param  P
	Target T
}

type DataSetParamInterface interface {
	DataSetParamDefault | DataSetParamAssist
}

type DataSetParamDefault struct {
	Value interface{}
}

type DataSetParamAssist struct {
	Value  interface{}
	Assist interface{}
}
type DataSetInterface interface {
	DataSetTarget
}

type DataSetTarget struct {
	Value interface{}
}

type FunDataSet[P DataSetParamInterface, T DataSetInterface] struct {
	HandlerDataSet[P, T]
}

type DataSetControl[P DataSetParamInterface, T DataSetInterface] struct {
}

type DataSetControlInterface[P DataSetParamInterface, T DataSetInterface] interface {
	Run(dataSet []DataSet[P, T], funDataSet []FunDataSet[P, T]) interface{}
}

func (c DataSetControl[P, T]) Run(dataSet []DataSet[P, T], funDataSet []FunDataSet[P, T]) interface{} {

	for i := 0; i < len(funDataSet); i++ {
		for j := 0; j < len(dataSet); j++ {
			dataSetCopy := DataSet[P, T]{}
			copier.Copy(&dataSetCopy.Param, &dataSet[j].Param)
			copier.Copy(&dataSetCopy.Target, &dataSet[j].Target)

			funDataSet[i].HandlerDataSet(dataSetCopy.Param, dataSetCopy.Target)
		}
	}

	return nil
}
