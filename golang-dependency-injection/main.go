package main

import (
	"golang-dependency-injection/helper"
	"golang-dependency-injection/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	// db := app.NewDB()
	// validate := validator.New()
	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)
	// router := app.NewRouter(categoryController)
	// authMiddleware := middleware.NewAuthMiddleware(router)

	// server := NewServer(authMiddleware)
	server := InitializedServer()
	
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
