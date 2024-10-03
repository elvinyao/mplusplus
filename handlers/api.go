package handlers

import (
	"matterplusplus/config"
	"matterplusplus/services"
	"net/http"

	"matterplusplus/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterAPIRoutes(router *gin.Engine, confluenceService *services.ConfluenceService) {
	for _, pageConfig := range config.AppConfig.Confluence {
		route := "/api/" + pageConfig.ID
		router.GET(route, func(c *gin.Context) {
			data, err := confluenceService.GetPageData(pageConfig)
			if err != nil {
				utils.Logger.Error("Failed to get page data", zap.Error(err))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}
			c.JSON(http.StatusOK, data)
		})
	}
}
