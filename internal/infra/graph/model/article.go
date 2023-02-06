package model

import (
	"fmt"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
)

func CreateNewArticle(id, name, ean, description string, category []*Category, images []*Image) *Article {
	return &Article{
		ID:          id,
		Name:        name,
		Ean:         ean,
		Description: description,
		Category:    category,
		Images:      images,
	}
}

func ToArticleDTO(article domain.Article) *Article {
	return &Article{
		ID:          fmt.Sprintf("%d", article.ID),
		Name:        article.Name,
		Ean:         article.EAN,
		Description: article.Description,
		Category:    ToModelCategories(article.Category),
		Images:      []*Image{},
	}
}
