package handlers

import (
	"matterplusplus/processors"
	"matterplusplus/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{}

func WebSocketHandler(processorManager *processors.ProcessorManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			utils.Logger.Error("Failed to set websocket upgrade", zap.Error(err))
			return
		}
		defer conn.Close()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				utils.Logger.Error("Error reading message", zap.Error(err))
				break
			}

			// 假设消息包含 action 字段，决定使用哪个处理器
			action := string(message) // 简化处理，实际应解析 JSON

			processor, found := processorManager.GetProcessor(action)
			if !found {
				utils.Logger.Warn("Processor not found", zap.String("action", action))
				continue
			}

			err = processor.Process(message)
			if err != nil {
				utils.Logger.Error("Processor failed", zap.Error(err))
			}
		}
	}
}
