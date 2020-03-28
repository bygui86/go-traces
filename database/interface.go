package database

import (
	"database/sql"
)

func GetProducts(db *sql.DB, start, count int) ([]Product, error) {
	rows, err := db.Query(getProductsQuery, count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func GetProduct(db *sql.DB, product *Product) error {
	return db.QueryRow(getProductQuery, product.ID).
		Scan(&product.Name, &product.Price)
}

func CreateProduct(db *sql.DB, product *Product) error {
	err := db.QueryRow(createProductQuery, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	_, err := db.Exec(updateProductQuery, product.Name, product.Price, product.ID)
	return err
}

func DeleteProduct(db *sql.DB, product *Product) error {
	_, err := db.Exec(deleteProductQuery, product.ID)
	return err
}
