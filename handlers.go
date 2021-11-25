package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/Eric0052/SuperMarket/consts"
	"github.com/Eric0052/SuperMarket/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PingHandle(c *gin.Context) {
	// typ := c.Query("type")
	// c.JSON(http.StatusOK, gin.H{
	// 	"type": typ,
	// })
	// return
	product := ProductFactory.NewRealProduct("apple", 12, 15)
	productJson, err := json.Marshal(product)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			Message:    "failed to marshal product",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		Message:    "get product done",
		Data:       string(productJson),
	})
}

func GetProductsHandle(c *gin.Context) {
	products := Store.GetProducts()
	productsJson, err := json.Marshal(products)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			Message:    "failed to marshal product",
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		Message:    "get product done",
		Data:       string(productsJson),
	})
}

func AddNewProductHandle(c *gin.Context) {
	productJson, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			Message:    "failed to get product",
		})
		return
	}
	if typ := c.Request.Header.Get("type"); typ == consts.ProductType_Virtual {
		virProduct := VirtualProduct{}
		json.Unmarshal(productJson, virProduct)
		virProduct.Id = uuid.New().String()
		Store.AddProduct(&virProduct)
	} else {
		realProduct := RealProduct{}
		json.Unmarshal(productJson, &realProduct)
		realProduct.Id = uuid.New().String()
		fmt.Printf("add product: %+v\n", realProduct)
		Store.AddProduct(&realProduct)
	}
	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		Message:    "Add a product done!",
	})

	//添加商品，返回给用户添加成功

}

func DeleteProductHandle(c *gin.Context) {
	ProId := c.Request.Header.Get("id")
	if _, ok := Store.products[ProId]; !ok {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 2,
			Message:    "no such product",
		})
		return
	}
	Store.DeleteProduct(ProId)
	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		Message:    "Delete product done!",
	})

}

func GetExactProduct(c *gin.Context) {
	ProId := c.Request.Header.Get("id")
	if _, ok := Store.products[ProId]; !ok {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 2,
			Message:    "no such product",
		})
		return
	}
	theProduct := Store.products[ProId]
	productJson, err := json.Marshal(theProduct)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			Message:    "fail to marshal product",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		Message:    "get single product done!",
		Data:       string(productJson),
	})

}

func GetOrder(c *gin.Context) {
	orderJson, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			Message:    "fail to finish this order!",
		})
		return
	}
	orders := []Order{}
	json.Unmarshal(orderJson, orders)
	success := Store.DoOrder(orders)
	if !success {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			Message:    "fail to finish this order!",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		Message:    "order done!",
	})

}

// func CreatNewProduct(c *gin.Context) {
// 	jsonBytes,err :=  ioutil.ReadAll(c.Request.Body)
// 	var product
// 	json.Unmarshal()
// }
