package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"short_url.com/internal/service"
)

type Handler struct {
	urlService service.URLShortener
}

type URLHandler interface {
	GetShortenedURL(c *gin.Context)
	GetOrignalURL(c *gin.Context)
	GetAnalytics(c *gin.Context)
}

func (h *Handler) GetShortenedURL(c *gin.Context) {

	req := &GetShortenedURLRequest{}

	err := c.BindJSON(req)
	if err != nil {
		return
	}

	shortURL, err := h.urlService.Shorten(req.OrignalURL)
	if err != nil {
		return
	}
	c.JSON(200, &GetShortenedURLResponse{
		ShortURL: shortURL,
	})

}

func (h *Handler) GetOrignalURL(c *gin.Context) {

	orignalURL, err := h.urlService.ResolveOrignalURL(c.Param("shortURL"))
	if err != nil {
		return
	}
	fmt.Println(orignalURL)
	c.Redirect(302, orignalURL)

}

func (h *Handler) GetAnalytics(c *gin.Context) {
	m, err := h.urlService.Analytics(c.Param("shortURL"))
	if err != nil {
		return
	}

	c.JSON(200, m)
}

func NewURLHandler(service service.URLShortener) URLHandler {
	return &Handler{
		urlService: service,
	}
}
