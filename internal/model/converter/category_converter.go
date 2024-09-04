package converter

import (
	"gofiber-hexagonal-architecture/internal/model/domain"
	"gofiber-hexagonal-architecture/internal/model/web"
)

func CategoryToResponse(category *domain.Category) *web.CategoryResponse {
	return &web.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}
