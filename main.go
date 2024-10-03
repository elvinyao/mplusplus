package main

import (
	"flag"
	"matterplusplus/config"
	"matterplusplus/handlers"
	"matterplusplus/processors"
	"matterplusplus/services"
	"matterplusplus/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 使用命令行参数指定配置文件路径
	configPath := flag.String("config", "config/server.yaml", "Path to configuration file")
	flag.Parse()

	// 加载配置
	err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	// 初始化日志
	err = utils.InitLogger(config.AppConfig.Server.LogLevel)
	if err != nil {
		panic(err)
	}

	// 初始化缓存服务和 Confluence 服务
	cacheService := services.NewCacheService()
	confluenceService := services.NewConfluenceService(cacheService)

	// 初始化处理器管理器并注册处理器
	processorManager := processors.NewProcessorManager()
	confluenceProcessor := processors.NewConfluenceProcessor("confluence", confluenceService)
	processorManager.Register(confluenceProcessor)

	// 设置 Gin 路由
	router := gin.Default()

	// 注册 API 路由
	handlers.RegisterAPIRoutes(router, confluenceService)

	// 注册 WebSocket 路由
	router.GET("/ws", handlers.WebSocketHandler(processorManager))

	// 启动服务器
	utils.Logger.Info("Starting server", zap.Int("port", config.AppConfig.Server.Port))
	router.Run(":" + strconv.Itoa(config.AppConfig.Server.Port))
}
