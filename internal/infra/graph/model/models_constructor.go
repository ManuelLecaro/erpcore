package model

import (
	"time"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
)

func (ca *CategoriesArticle) ToDomainCategory() *domain.Category {
	return &domain.Category{
		Name:      ca.Name,
		Children:  ToDomainCategories(ca.Children),
		CreatedAt: time.Now(),
	}
}

func ToDomainCategories(categories []*CategoriesArticle) []*domain.Category {
	domainCategories := []*domain.Category{}

	for _, category := range categories {
		domainCategories = append(domainCategories, category.ToDomainCategory())
	}

	return domainCategories
}
