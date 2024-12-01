package service

import (
	"go-api/model"
	"go-api/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return ProductService{repository: repository}
}

func (p *ProductService) GetAll() ([]model.Product, error) {
	return p.repository.GetAll()
}

func (p *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	return p.repository.CreateProduct(product)
}
