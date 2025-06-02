package api

import (
	"github.com/gin-gonic/gin"
	"short_url.com/internal/service"
)

type Handler struct {
	urlService service.URLShortener
}

type URLHandler interface {
	GetShortenedURL(c *gin.Context)
	GetOrignalURL(c *gin.Context)
}

func (h *Handler) GetShortenedURL(c *gin.Context) {

	_, err := h.urlService.Shorten("")
	if err != nil {
		return
	}

}

func (h *Handler) GetOrignalURL(c *gin.Context) {
	_, err := h.urlService.ResolveOrignalURL("")
	if err != nil {
		return
	}

}

func NewURLHandler(service service.URLShortener) URLHandler {
	return &Handler{
		urlService: service,
	}
}
