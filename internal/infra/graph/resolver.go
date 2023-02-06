package graph

import (
	"github.com/ManuelLecaro/erpcore/internal/core/ports/services"
	"github.com/ManuelLecaro/erpcore/internal/core/transaction"
)

// This file will not be regenerated automatically.

// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ArticleService     services.IArticleService
	CategoryService    services.ICategoryService
	TransactionService transaction.Transaction
}
