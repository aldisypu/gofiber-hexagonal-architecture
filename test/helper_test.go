package test

import (
	"gofiber-hexagonal-architecture/internal/model/domain"
	"testing"

	"github.com/google/uuid"
)

func ClearCategory() {
	err := db.Where("id is not null").Delete(&domain.Category{}).Error
	if err != nil {
		log.Fatalf("Failed clear category data : %+v", err)
	}
}

func CreateCategories(t *testing.T, total int) {
	for i := 0; i < total; i++ {
		category := &domain.Category{
			ID:   uuid.NewString(),
			Name: "Category",
		}
		err := db.Create(category).Error
		if err != nil {
			log.Fatalf("Failed create category data : %+v", err)
		}
	}
}
