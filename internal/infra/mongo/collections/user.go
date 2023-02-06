package collections

import (
	"context"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/repositories"
	"github.com/ManuelLecaro/erpcore/internal/infra/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

const UserCollectionName = "user"

type UserCollection struct {
	mongoConnection *mongo.MongoConnection
}

func NewUserCollection(connection *mongo.MongoConnection) repositories.IUserRepository {
	return &UserCollection{
		mongoConnection: connection,
	}
}

func (ac *UserCollection) Create(ctx context.Context, article domain.User) (*domain.User, error) {
	_, err := ac.mongoConnection.Connection.Collection(UserCollectionName).InsertOne(ctx, article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (ac *UserCollection) GetByID(ctx context.Context, id string) (*domain.User, error) {
	filter := bson.M{"id": id}
	user := &domain.User{}

	err := ac.mongoConnection.Connection.Collection(ArticleCollectionName).FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
