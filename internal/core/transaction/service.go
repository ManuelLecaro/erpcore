package transaction

import (
	"context"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/repositories"
)

// https://www.udemy.com/course/golang-how-to-build-a-blockchain-in-go/

type Transaction struct {
	collection repositories.ITransactionRepository
}

// NewTransactionService creates a service use to manage the transaction logic.
func NewTransactionService(repo repositories.ITransactionRepository) *Transaction {
	return &Transaction{
		collection: repo,
	}
}

// Creates a new Transaction.
func (c *Transaction) Create(ctx context.Context, transaction domain.Transaction) (*domain.Transaction, error) {
	return c.collection.Create(ctx, transaction)
}
