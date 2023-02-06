package article

import (
	"context"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/repositories"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/services"
)

type Article struct {
	collection repositories.IArticleRepository
}

// NewArticleService creates a service use to manage the article logic.
func NewArticleService(repo repositories.IArticleRepository) services.IArticleService {
	return &Article{
		collection: repo,
	}
}

// Creates a new Article.
func (c *Article) Create(ctx context.Context, article domain.Article) (*domain.Article, error) {
	return c.collection.Create(ctx, article)
}

func (c *Article) Search(ctx context.Context, fields ...string) ([]*domain.Article, error) {
	return c.collection.Search(ctx, fields...)
}

func (c *Article) GetByID(ctx context.Context, id string) (*domain.Article, error) {
	return c.collection.GetByID(ctx, id)
}
