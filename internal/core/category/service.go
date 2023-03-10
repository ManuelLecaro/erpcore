package category

import (
	"context"
	"sync"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/repositories"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/services"
	"github.com/google/uuid"
)

var wg sync.WaitGroup

type Category struct {
	collection repositories.ICategoryRepository
}

// NewCategoryService creates a service use to manage the category logic.
func NewCategoryService(repo repositories.ICategoryRepository) services.ICategoryService {
	return &Category{
		collection: repo,
	}
}

// Creates a new Category.
func (c *Category) Create(ctx context.Context, category domain.Category) (*domain.Category, error) {
	category.ID = uuid.New().String()
	category.Path = ","

	createMaterializedPath(ctx, &category)

	categories := []domain.Category{}

	flatCategories(ctx, category, &categories)

	if _, err := c.collection.CreateMany(ctx, categories); err != nil {
		return nil, err
	}

	return &category, nil
}

func createMaterializedPath(ctx context.Context, category *domain.Category) bool {
	currentPath := category.Path

	for _, cat := range category.Children {
		cat.ID = uuid.New().String()
		cat.Path = currentPath + category.Name + ","

		createMaterializedPath(ctx, cat)
	}

	return true
}

func flatCategories(ctx context.Context, category domain.Category, categoriesArray *[]domain.Category) bool {
	*categoriesArray = append(*categoriesArray, category)

	for _, cat := range category.Children {
		flatCategories(ctx, *cat, categoriesArray)
	}

	return true
}

func (c *Category) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	return c.collection.GetByID(ctx, id)
}

func (c *Category) GetAll(ctx context.Context) ([]*domain.Category, error) {
	return c.collection.Get(ctx)
}
