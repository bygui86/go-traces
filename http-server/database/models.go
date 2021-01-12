package database

import (
	"context"
	"fmt"
)

type InMemoryDbInt interface {
	GetProducts(start, count int, ctx context.Context) ([]*Product, error)
	GetProduct(product *Product, ctx context.Context) *Product
	CreateProduct(product *Product, ctx context.Context) error
	UpdateProduct(product *Product, ctx context.Context) error
	DeleteProduct(productId string, ctx context.Context) error
}

// InMemoryDb implements InMemoryDbInt
type InMemoryDb struct {
	products map[string]*Product // [id] product
}

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) String() string {
	return fmt.Sprintf("ID[%s], Name[%s], Price[%f]",
		p.ID, p.Name, p.Price)
}
