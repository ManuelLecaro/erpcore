package services

import (
	"context"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
)

//go:generate mockery --name IArticleService
//go:generate mockery --name ICategoryService
//go:generate mockery --name ITransactionService

type IArticleService interface {
	Create(ctx context.Context, article domain.Article) (*domain.Article, error)
	Search(ctx context.Context, fields ...string) ([]*domain.Article, error)
	GetByID(ctx context.Context, id string) (*domain.Article, error)
}

type ICategoryService interface {
	Create(ctx context.Context, article domain.Category) (*domain.Category, error)
}

type ITransactionService interface {
	Create(ctx context.Context, category domain.Transaction) (*domain.Transaction, error)
}
