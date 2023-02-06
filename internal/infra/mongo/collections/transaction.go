package collections

import (
	"context"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/repositories"
	"github.com/ManuelLecaro/erpcore/internal/infra/mongo"
)

const TransactionCollectionName = "transaction"

type TransactionCollection struct {
	mongoConnection *mongo.MongoConnection
}

func NewTransactionCollection(connection *mongo.MongoConnection) repositories.ITransactionRepository {
	return &TransactionCollection{
		mongoConnection: connection,
	}
}

func (ac *TransactionCollection) Create(ctx context.Context, article domain.Transaction) (*domain.Transaction, error) {
	_, err := ac.mongoConnection.Connection.Collection(TransactionCollectionName).InsertOne(ctx, article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}
