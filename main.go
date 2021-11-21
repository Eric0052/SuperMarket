package main

import (
	"github.com/gin-gonic/gin"
)

var (
	Store = NewStoreHouse()
	Fac   = Factory{}
)

func main() {

	r := gin.Default()
	r.GET("/ping", PingHandle)
	r.GET("/product", GetProductsHandle)
	r.POST("/product", AddNewProductHandle)
	r.DELETE("/product", DeleteProductHandle)
	r.GET("/product/singleProduct", GetExactProduct)
	// r.GET("/good/:id", GetGoodHandle)
	// r.PUT("/good/:id", ModifyGoodHandle)
	// r.DELETE("good/:id")
	r.Run()
}
