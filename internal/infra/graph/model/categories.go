package model

import (
	"time"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
)

func CreateNewCategory(category domain.Category) *Category {
	return &Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Category:    ToModelCategories(category.Children),
	}
}

func ToModelCategories(categories []*domain.Category) []*Category {
	domainCategories := []*Category{}

	for _, category := range categories {
		domainCategories = append(domainCategories,
			CreateNewCategory(*category),
		)
	}

	return domainCategories
}

func (ca *Category) ToDomainCategory() *domain.Category {
	return &domain.Category{
		Name:        ca.Name,
		Children:    ToDomainCategoryArray(ca.Category),
		Description: ca.Description,
		CreatedAt:   time.Now(),
	}
}

func ToDomainCategoryArray(input []*Category) []*domain.Category {
	domainCategories := []*domain.Category{}

	for _, category := range input {
		domainCategories = append(domainCategories, category.ToDomainCategory())
	}

	return domainCategories
}

func (ca *NewCategory) ToDomainCategoryFromInput() *domain.Category {
	return &domain.Category{
		Name:        ca.Name,
		Children:    ToDomainCategoryFromImput(ca.Category),
		Description: ca.Description,
		CreatedAt:   time.Now(),
	}
}

func ToDomainCategoryFromImput(input []*NewCategory) []*domain.Category {
	domainCategories := []*domain.Category{}

	for _, category := range input {
		domainCategories = append(domainCategories, category.ToDomainCategoryFromInput())
	}

	return domainCategories
}
