package service

import (
	"github.com/nalyx1/marketplace/go/internal/database"
	"github.com/nalyx1/marketplace/go/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (s *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := s.CategoryDB.GetCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) GetCategoryByID(id string) (*entity.Category, error) {
	category, err := s.CategoryDB.GetCategoryByID(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := s.CategoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

