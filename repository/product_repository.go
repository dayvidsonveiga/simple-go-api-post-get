package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{connection: connection}
}

func (p *ProductRepository) GetAll() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"

	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	products := []model.Product{}
	var product model.Product
	for rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		products = append(products, product)
	}

	rows.Close()

	return products, nil
}

func (p *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	query, err := p.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&product.Id)
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	query.Close()

	return product, nil
}
