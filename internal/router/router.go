package router

import (
	"gofiber-hexagonal-architecture/internal/controller"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	CategoryController controller.CategoryController
}

func (c *RouteConfig) Setup() {
	c.SetupRoute()
}

func (c *RouteConfig) SetupRoute() {
	c.App.Post("/api/categories", c.CategoryController.Create)
	c.App.Put("/api/categories/:categoryId", c.CategoryController.Update)
	c.App.Delete("/api/categories/:categoryId", c.CategoryController.Delete)
	c.App.Get("/api/categories/:categoryId", c.CategoryController.Get)
	c.App.Get("/api/categories", c.CategoryController.List)
}
