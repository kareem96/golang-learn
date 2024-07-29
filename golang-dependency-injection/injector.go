//go:build wireinject
// +build wireinject

package main

import (
	"golang-dependency-injection/app"
	"golang-dependency-injection/controller"
	"golang-dependency-injection/middleware"
	"golang-dependency-injection/repository"
	"golang-dependency-injection/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func ProviderValidator() *validator.Validate {
	return validator.New()
}

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		ProviderValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}