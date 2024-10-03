package processors

import "sync"

type ProcessorManager struct {
	processors sync.Map
}

func NewProcessorManager() *ProcessorManager {
	return &ProcessorManager{}
}

func (m *ProcessorManager) Register(processor Processor) {
	m.processors.Store(processor.GetID(), processor)
}

func (m *ProcessorManager) GetProcessor(id string) (Processor, bool) {
	processor, ok := m.processors.Load(id)
	if !ok {
		return nil, false
	}
	return processor.(Processor), true
}
