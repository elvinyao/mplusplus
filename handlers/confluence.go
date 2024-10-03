package handlers

import (
	"net/http"

	"matterplusplus/services"
	"matterplusplus/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetConfluenceData(service *services.ConfluenceService) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := service.GetPageData()
		if err != nil {
			utils.Logger.Error("获取 Confluence 数据失败", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": gin.H{
					"code":    "INTERNAL_ERROR",
					"message": err.Error(),
				},
			})
			return
		}
		utils.Logger.Info("成功返回数据", zap.String("PageID", service.Config.ID))
		c.JSON(http.StatusOK, data)
	}
}
