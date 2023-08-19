package main

import "github.com/aimbot1526/test-go/db"

type Inject struct {
	ShopService    *db.ShopRepository
	ProductService *db.ProductRepository
}

func NewInject() *Inject {
	inj := &Inject{}

	inj.ShopService = &db.ShopRepository{}
	inj.ProductService = &db.ProductRepository{}

	return inj
}
