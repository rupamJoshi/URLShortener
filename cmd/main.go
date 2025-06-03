package main

import (
	"github.com/gin-gonic/gin"
	"short_url.com/api"
	"short_url.com/internal"
)

func main() {

	r := gin.Default()

	server := internal.InitServer()
	api.SetupRoutes(r, server.Handler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
