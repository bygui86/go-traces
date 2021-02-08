package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

func (db *InMemoryDb) GetProducts(start, count int, ctx context.Context) ([]*Product, error) {
	parentSpan := opentracing.SpanFromContext(ctx)
	var parentCtx opentracing.SpanContext
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}
	span := opentracing.StartSpan(
		"get-products-db",
		opentracing.ChildOf(parentCtx),
	)
	defer span.Finish()

	query := fmt.Sprintf("all-start[%d]-count[%d]", start, count)
	span.SetTag("query", query)
	span.SetTag("count", count)
	span.SetTag("start", start)
	span.LogKV(
		"query", query,
		"count", count,
		"start", start,
	)

	products := make([]*Product, 0)
	for _, prod := range db.products {
		products = append(products, prod)
	}

	span.SetTag("products-found", len(products))
	span.LogKV("products-found", len(products))

	return products, nil
}

func (db *InMemoryDb) GetProduct(product *Product, ctx context.Context) *Product {
	parentSpan := opentracing.SpanFromContext(ctx)
	var parentCtx opentracing.SpanContext
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}
	span := opentracing.StartSpan(
		"get-product-db",
		opentracing.ChildOf(parentCtx),
	)
	defer span.Finish()

	span.SetTag("product-id", product.ID)
	span.LogKV("product-id", product.ID)

	return db.products[product.ID]
}

func (db *InMemoryDb) CreateProduct(product *Product, ctx context.Context) error {
	parentSpan := opentracing.SpanFromContext(ctx)
	var parentCtx opentracing.SpanContext
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}
	span := opentracing.StartSpan(
		"create-product-db",
		opentracing.ChildOf(parentCtx),
	)
	defer span.Finish()

	span.SetTag("product", product.String())
	span.LogKV("product", product.String())

	newUuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	product.ID = newUuid.String()
	db.products[newUuid.String()] = product
	return nil
}

func (db *InMemoryDb) UpdateProduct(product *Product, ctx context.Context) error {
	parentSpan := opentracing.SpanFromContext(ctx)
	var parentCtx opentracing.SpanContext
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}
	span := opentracing.StartSpan(
		"update-product-db",
		opentracing.ChildOf(parentCtx),
	)
	defer span.Finish()

	span.SetTag("product", product.String())
	span.LogKV("product", product.String())

	if product != nil {
		if product.ID != "" {
			db.products[product.ID] = product
		} else {
			return errors.New("updated product ID cannot be empty")
		}
	} else {
		return errors.New("updated product cannot be empty")
	}
	return nil
}

func (db *InMemoryDb) DeleteProduct(productId string, ctx context.Context) error {
	parentSpan := opentracing.SpanFromContext(ctx)
	var parentCtx opentracing.SpanContext
	if parentSpan != nil {
		parentCtx = parentSpan.Context()
	}
	span := opentracing.StartSpan(
		"delete-product-db",
		opentracing.ChildOf(parentCtx),
	)
	defer span.Finish()

	span.SetTag("product-id", productId)
	span.LogKV("product-id", productId)

	if productId != "" {
		db.products[productId] = nil
	} else {
		return errors.New("product ID cannot be empty")
	}
	return nil
}
