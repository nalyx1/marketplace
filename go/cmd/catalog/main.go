package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nalyx1/marketplace/go/internal/controller"
	"github.com/nalyx1/marketplace/go/internal/database"
	"github.com/nalyx1/marketplace/go/internal/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/catalog")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)
	categoryController := webserver.NewCategoryController(categoryService)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)
	productController := webserver.NewProductController(productService)

	c := chi.NewRouter()

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/categories", categoryController.GetCategories)
	c.Get("/categories/{id}", categoryController.GetCategoryByID)
	c.Post("/categories", categoryController.CreateCategory)

	c.Get("/products", productController.GetProducts)
	c.Get("/products/category/{id}", productController.GetProductsByCategoryID)
	c.Get("/products/{id}", productController.GetProductByID)
	c.Post("/products", productController.CreateProduct)
	
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}