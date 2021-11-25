package main

var (
	_ Product = (*RealProduct)(nil)
)

type Product interface {
	GetId() string
	GetName() string
	GetPrice() float32
	GetStock() int
	ChangeStock(add int)
}

//实物商品类
type RealProduct struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Weight float32 `json:"weight"`
	Stock  int     `json:"stock"`
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

func (realPro *RealProduct) GetStock() int {
	return realPro.Stock
}

func (realPro *RealProduct) ChangeStock(add int) {
	realPro.Stock = realPro.Stock + add
}

type VirtualProduct struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
	Content string  `json:"content"`
	Stock   int     `json:"Stock"`
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

func (virPro *VirtualProduct) GetStock() int {
	return virPro.Stock
}

func (virPro *VirtualProduct) ChangeStock(add int) {
	virPro.Stock = virPro.Stock + add
}
