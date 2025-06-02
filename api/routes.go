package api

import (
	"github.com/gin-gonic/gin"
	"short_url.com/internal/service"
)

func SetupRoutes(router *gin.Engine) {

	service := service.NewURLService()
	handler := NewURLHandler(service)
	router.POST("/shorten", handler.GetShortenedURL)
	router.GET("/:shortURL", handler.GetOrignalURL)

}
