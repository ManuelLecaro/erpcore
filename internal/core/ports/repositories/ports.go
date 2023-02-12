package repositories

import (
	"context"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
)

//go:generate mockery --name IArticleRepository
//go:generate mockery --name ICategoryRepository
//go:generate mockery --name ITransactionRepository
//go:generate mockery --name IUserRepository
//go:generate mockery --name IFilesRepository

type IArticleRepository interface {
	Create(ctx context.Context, article domain.Article) (*domain.Article, error)
	GetByID(ctx context.Context, id string) (*domain.Article, error)
	Search(ctx context.Context, fields map[string]string) ([]*domain.Article, error)
}

type ICategoryRepository interface {
	Create(ctx context.Context, category domain.Category) (*domain.Category, error)
	GetByID(ctx context.Context, id string) (*domain.Category, error)
	Get(ctx context.Context) ([]*domain.Category, error)
	CreateMany(ctx context.Context, categories []domain.Category) ([]domain.Category, error)
}

type ITransactionRepository interface {
	Create(ctx context.Context, transaction domain.Transaction) (*domain.Transaction, error)
}

type IUserRepository interface {
	Create(ctx context.Context, user domain.User) (*domain.User, error)
	GetByID(ctx context.Context, id string) (*domain.User, error)
}

type IFilesRepository interface {
	Save(ctx context.Context, file domain.User) error
	GetByUUID(ctx context.Context, uuid string) (string, error)
}
