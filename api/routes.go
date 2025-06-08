package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handler URLHandler) {
	router.POST("/shorten", handler.GetShortenedURL)
	router.GET("/:shortURL", handler.GetOrignalURL)
	router.GET("/count", handler.GetAnalytics)
}
