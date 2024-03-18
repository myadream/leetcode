package common

import (
	"fmt"

	"github.com/jinzhu/copier"
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

type ProcessorMethod[P SourceData, T TargetData] func(P, T) bool

type Processor[P SourceData, T TargetData] struct {
	Fn ProcessorMethod[P, T]
}

type DataFlowProcessor[P SourceData, T TargetData] struct {
}

type DataFlowProcessorInterface[P SourceData, T TargetData] interface {
	ProcessData(processors []Processor[P, T], dataCarriers []DataCarrier[P, T]) interface{}
}

func (c DataFlowProcessor[P, T]) ProcessData(processors []Processor[P, T], dataCarriers []DataCarrier[P, T]) interface{} {

	channels := make(chan int)
	for index, processor := range processors {

		go func(processor Processor[P, T], carriers []DataCarrier[P, T], index int) {
			carrierCopy := DataCarrier[P, T]{}
			for _, carrier := range carriers {
				if err := copier.Copy(&carrierCopy, &carrier); err != nil {
					panic(err)
				}
				if !processor.Fn(carrierCopy.SourceData, carrierCopy.TargetData) {
					break
				}
			}

			channels <- index
		}(processor, dataCarriers, index)
	}

	for i := 0; i < len(processors); i++ {
		fmt.Println(<-channels)
	}

	return nil
}
