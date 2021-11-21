package main

var (
	_       Product = (*RealProduct)(nil)
)

type Product interface {
	GetId() string
	GetName() string
	GetPrice() float32
}

//实物商品类
type RealProduct struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Weight float32 `json:"weight"`
}

func (realPro *RealProduct) GetId() string {
	return realPro.Id
}

func (realPro *RealProduct) GetName() string {
	return realPro.Name
}

func (realPro *RealProduct) GetPrice() float32 {
	return realPro.Price
}

func (realPro *RealProduct) GetWeight() float32 {
	return realPro.Weight
}

type VirtualProduct struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Content string  `json:"content"`
}

func (virPro *VirtualProduct) GetId() string {
	return virPro.Id
}

func (virPro *VirtualProduct) GetName() string {
	return virPro.Name
}

func (virPro *VirtualProduct) GetPrice() float32 {
	return virPro.Price
}

func (virPro *VirtualProduct) GetContent() string {
	return virPro.Content
}
