package service

import (
	"github.com/nalyx1/marketplace/go/internal/database"
	"github.com/nalyx1/marketplace/go/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: productDB,
	}
}

func (s *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := s.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := s.ProductDB.GetProductsByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductByID(id string) (*entity.Product, error) {
	product, err := s.ProductDB.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}


func (s *ProductService) CreateProduct(name string, description string, price float64, categoryID string, imageURL string) (*entity.Product, error) {
	product := entity.NewProduct(name, description, price, categoryID, imageURL)
	_, err := s.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

