package main

import "github.com/google/uuid"

var (
	ProductFactory *Factory = &Factory{}
)

type Factory struct{}

func (fac *Factory) NewRealProduct(name string, price, weight float32) *RealProduct {

	product := &RealProduct{
		Id:     uuid.New().String(),
		Name:   name,
		Price:  price,
		Weight: weight,
	}
	return product
}

func (fac *Factory) NewVirtualProduct(name, content string, price float32) *VirtualProduct {
	product := &VirtualProduct{
		Id:      uuid.New().String(),
		Name:    name,
		Price:   price,
		Content: content,
	}
	return product

}
