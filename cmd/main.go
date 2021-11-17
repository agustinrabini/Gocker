package main

import (
	"net/http"

	"github.com/agustinrabini/Gocker/cmd/server/handler"
	"github.com/agustinrabini/Gocker/internal/database"
	"github.com/agustinrabini/Gocker/internal/product"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	productRepository := product.NewRepository(database.DbConecction())
	productService := product.NewService(productRepository)
	productHandler := handler.NewProduct(productService)
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "arranco") })
		api.GET("/", productHandler.GetAll())
	}

	r.Run(":8080")
}
