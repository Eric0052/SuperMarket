package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type StoreHouse struct {
	products  map[string]Product
	needStore bool
}

//创建StoreHouse对象，同时完成products的创建
func NewStoreHouse() *StoreHouse {
	store := &StoreHouse{
		products:  make(map[string]Product),
		needStore: false,
	}

	store.loadRealProducts()
	store.loadVirtualProducts()

	go func() {
		for {
			if store.needStore {
				store.SaveProducts()
				time.Sleep(5 * time.Second)
				store.needStore = false
			}
		}
	}()

	return store
}

func (store *StoreHouse) GetProducts() []Product {
	products := []Product{}
	for _, product := range store.products {
		products = append(products, product)
	}

	return products
}

//商品的增加
func (store *StoreHouse) AddProduct(product Product) {
	store.products[product.GetId()] = product
	store.needStore = true
}

//商品的删除
func (store *StoreHouse) DeleteProduct(id string) {
	delete(store.products, id)
	store.needStore = true
}

func (store *StoreHouse) SaveProducts() {
	products := store.GetProducts()

	jsonBytes, _ := json.Marshal(products)

	ioutil.WriteFile("real_products.json", jsonBytes, 0666)
}

//加载实物商品的信息
func (store *StoreHouse) loadRealProducts() {
	jsonBytes, _ := ioutil.ReadFile("real_products.json")
	products := []RealProduct{}
	json.Unmarshal(jsonBytes, &products)

	for _, product := range products {
		store.products[product.Id] = &product
	}
}

//加载虚拟商品的信息
func (store *StoreHouse) loadVirtualProducts() {
	jsonBytes, _ := ioutil.ReadFile("virtual_products.json")
	products := []VirtualProduct{}
	json.Unmarshal(jsonBytes, &products)

	for _, product := range products {
		store.products[product.Id] = &product
	}

}
