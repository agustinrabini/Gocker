package main

import (
	"context"
	"strconv"

	"github.com/agustinrabini/Gocker/internal/domain"
	"github.com/agustinrabini/Gocker/internal/product"
	"github.com/gin-gonic/gin"
)

type Product struct {
	productService product.Service
}

func NewProduct(w product.Service) *Product {
	return &Product{
		productService: w,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {

	type response struct {
		Data []domain.Product `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		products, err := p.productService.GetAll(ctx)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}

		c.JSON(200, &response{products})
	}
}

func (p *Product) Get() gin.HandlerFunc {

	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}
		ctx := context.Background()
		prod, err := p.productService.Get(ctx, int(id))
		if err != nil {
			c.JSON(404, err.Error())
			return
		}

		c.JSON(200, &response{prod})
	}
}

func (p *Product) Store() gin.HandlerFunc {

	type request struct {
		Name        string `json:"name"`
		Brand       string `json:"brand"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Price       int    `json:"price"`
		Stock       int    `json:"stock"`
	}

	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		if err := c.Bind(&req); err != nil {
			c.JSON(422, "json decoding: "+err.Error())
			return
		}
		if req.Name == "" {
			c.JSON(422, "name can not be empty")
			return
		}
		if req.Brand == "" {
			c.JSON(422, "brand can not be empty")
			return
		}
		if req.Description == "" {
			c.JSON(422, "description can not be empty")
			return
		}

		if req.Image == "" {
			c.JSON(422, "image can not be empty")
			return
		}
		if req.Price == 0 {
			c.JSON(422, "price can not be empty")
			return
		}

		if req.Stock < 0 {
			c.JSON(422, "stock can not be empty")
			return
		}

		product := domain.Product{
			Name:        req.Name,
			Brand:       req.Brand,
			Description: req.Description,
			Image:       req.Image,
			Price:       req.Price,
			Stock:       req.Stock,
		}

		ctx := context.Background()
		prod, err := p.productService.Store(ctx, product)
		if err != nil {

			c.JSON(500, err.Error())
			return
		}

		c.JSON(201, &response{prod})
	}
}
