package processors

import (
	"matterplusplus/services"
	"matterplusplus/utils"

	"go.uber.org/zap"
)

type ConfluenceProcessor struct {
	ID                string
	ConfluenceService *services.ConfluenceService
}

func NewConfluenceProcessor(id string, service *services.ConfluenceService) *ConfluenceProcessor {
	return &ConfluenceProcessor{
		ID:                id,
		ConfluenceService: service,
	}
}

func (p *ConfluenceProcessor) Process(message interface{}) error {
	utils.Logger.Info("Processing message with ConfluenceProcessor", zap.String("processor_id", p.ID))
	// 处理逻辑（此处使用假方法）
	return nil
}

func (p *ConfluenceProcessor) GetID() string {
	return p.ID
}
