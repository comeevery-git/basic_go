package service

// Network, repository 의 다리 역할
// (API Reqeust) Network => Service => Repository

import (
	"errors"
	"example.com/m/model"
	"example.com/m/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(r repository.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) GetProduct(id int) (*model.Product, error) {
	Product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if Product == nil {
		return nil, errors.New("Product not found")
	}
	return Product, nil
}

func (s *ProductService) GetAllProducts() ([]*model.Product, error) {
    Products, err := s.repo.GetAllProducts()
    if err != nil {
        return nil, err
    }
    if Products == nil {
        return nil, errors.New("No Products found")
    }
	return Products, nil
}

