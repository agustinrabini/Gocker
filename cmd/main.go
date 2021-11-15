package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "arranco") })
	}

	r.Run(":8000")
}
