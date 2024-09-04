package config

import (
	"gofiber-hexagonal-architecture/internal/controller"
	"gofiber-hexagonal-architecture/internal/repository"
	"gofiber-hexagonal-architecture/internal/router"
	"gofiber-hexagonal-architecture/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	categoryRepository := repository.NewCategoryRepository(config.Log)

	// setup service
	categoryService := service.NewCategoryService(config.DB, config.Log, config.Validate, categoryRepository)

	// setup controller
	categoryController := controller.NewCategoryController(categoryService, config.Log)

	routeConfig := router.RouteConfig{
		App:                config.App,
		CategoryController: categoryController,
	}
	routeConfig.Setup()
}
