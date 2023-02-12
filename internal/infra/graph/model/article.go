package model

import (
	"fmt"
	"time"

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

func ToDomainArticleFromImput(input NewArticle) *domain.Article {
	return &domain.Article{
		Category:    ToDomainCategories(input.Categories),
		Images:      ToDomainImageFromInputArray(input.Images),
		CreatedAt:   time.Now(),
		Name:        input.Name,
		EAN:         input.Ean,
		Description: input.Description,
	}
}
