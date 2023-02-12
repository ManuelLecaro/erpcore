package collections

import (
	"context"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/repositories"
	"github.com/ManuelLecaro/erpcore/internal/infra/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ArticleCollectionName = "article"

type ArticleCollection struct {
	mongoConnection *mongo.MongoConnection
}

func NewArticleCollection(connection *mongo.MongoConnection) repositories.IArticleRepository {
	return &ArticleCollection{
		mongoConnection: connection,
	}
}

func (ac *ArticleCollection) Create(ctx context.Context, article domain.Article) (*domain.Article, error) {
	_, err := ac.mongoConnection.Connection.Collection(ArticleCollectionName).InsertOne(ctx, article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (ac *ArticleCollection) GetByID(ctx context.Context, id string) (*domain.Article, error) {
	filter := bson.M{"id": id}
	article := &domain.Article{}

	err := ac.mongoConnection.Connection.Collection(ArticleCollectionName).FindOne(ctx, filter).Decode(article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (ac *ArticleCollection) Search(ctx context.Context, fields map[string]string) ([]*domain.Article, error) {
	filter := bson.D{}
	articles := []*domain.Article{}

	for field, value := range fields {
		filter = append(filter, primitive.E{Key: field, Value: value})
	}

	cur, err := ac.mongoConnection.Connection.Collection(ArticleCollectionName).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	cur.All(ctx, &articles)

	return articles, nil
}
