package common

import (
	"github.com/jinzhu/copier"
	"sync"
)

type DataCarrier[P SourceData, T TargetData] struct {
	SourceData P
	TargetData T
}

type SourceData interface {
	DCDefault | interface{}
}

type DCDefault struct {
	Value  interface{}
	Assist interface{}
}

type TargetData interface {
	TDefault | interface{}
}

type TDefault struct {
	Value interface{}
}

type Processor[P SourceData, T TargetData] func(P, T) bool

type DataFlowProcessor[P SourceData, T TargetData] struct {
}

type DataSetControlInterface[P SourceData, T TargetData] interface {
	ProcessData(processors []Processor[P, T], dataCarriers []DataCarrier[P, T]) interface{}
}

func (c DataFlowProcessor[P, T]) ProcessData(processors []Processor[P, T], dataCarriers []DataCarrier[P, T]) interface{} {
	var wg sync.WaitGroup

	chs := make(chan DataCarrier[P, T], len(dataCarriers))
	for _, carrier := range dataCarriers {
		carrierCopy := DataCarrier[P, T]{}
		if err := copier.Copy(&carrierCopy, &carrier); err == nil {
			chs <- carrierCopy
		}
	}

	for _, processor := range processors {
		go func(processor Processor[P, T]) {
			for carrier := range chs {
				if !processor(carrier.SourceData, carrier.TargetData) {
					return
				}
			}
			wg.Done()
		}(processor)
	}

	wg.Wait()
	close(chs)

	return nil
}
