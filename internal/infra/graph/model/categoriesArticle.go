package model

import (
	"time"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
)

func (ca *CategoriesArticle) ToDomainCategory() *domain.Category {
	return &domain.Category{
		Children:  ToDomainCategories(ca.Children),
		CreatedAt: time.Now(),
		Name:      ca.Name,
	}
}

func ToDomainCategories(categories []*CategoriesArticle) []*domain.Category {
	domainCategories := []*domain.Category{}

	for _, category := range categories {
		domainCategories = append(domainCategories, category.ToDomainCategory())
	}

	return domainCategories
}
