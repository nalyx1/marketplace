package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nalyx1/marketplace/go/internal/entity"
	"github.com/nalyx1/marketplace/go/internal/service"
)

type CategoryController struct {
	CategoryService *service.CategoryService
}

func NewCategoryController(categoryService *service.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (c *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.CategoryService.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (c *CategoryController) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	category, err := c.CategoryService.GetCategoryByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (c *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdCategory, err := c.CategoryService.CreateCategory(category.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdCategory)
}