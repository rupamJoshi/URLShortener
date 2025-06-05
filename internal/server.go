package internal

import (
	"short_url.com/api"
	"short_url.com/config"
	"short_url.com/internal/service"
)

type Server struct {
	Service service.URLShortener
	Handler api.URLHandler
}

func InitServer() Server {
	config := config.NewConfig("./config/config.yaml")
	service := service.NewURLService(config)
	handler := api.NewURLHandler(service)

	return Server{
		Service: service,
		Handler: handler,
	}
}
