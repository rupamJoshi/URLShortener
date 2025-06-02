package main

import (
	"github.com/gin-gonic/gin"
	"short_url.com/api"
)

func main() {

	r := gin.Default()
	api.SetupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
