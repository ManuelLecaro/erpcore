package collections

import (
	"context"
	"errors"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/repositories"
	"github.com/ManuelLecaro/erpcore/internal/infra/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CategoryCollectionName = "category"

var ErrAlreadyCreatedCategory = errors.New("error: category already created")

type CategoryCollection struct {
	mongoConnection *mongo.MongoConnection
}

func NewCategoryCollection(connection *mongo.MongoConnection) repositories.ICategoryRepository {
	connection.AddIndex(context.Background(), "name", CategoryCollectionName, options.Index().SetUnique(true)) // hacer indice de unico valor para la categoria

	return &CategoryCollection{
		mongoConnection: connection,
	}
}

func (ac *CategoryCollection) Create(ctx context.Context, category domain.Category) (*domain.Category, error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "name", Value: category.Name}}
	update := bson.D{
		{
			Key: "$setOnInsert", Value: bson.D{
				{Key: "id", Value: category.ID},
				{Key: "name", Value: category.Name},
				{Key: "description", Value: category.Description},
				{Key: "children", Value: category.Children},
				{Key: "created_at", Value: category.CreatedAt},
				{Key: "path", Value: category.Path},
			},
		},
		{
			Key: "$set", Value: bson.D{},
		},
	}

	res, err := ac.mongoConnection.Connection.Collection(CategoryCollectionName).
		UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}

	if res.UpsertedID == nil {
		return nil, ErrAlreadyCreatedCategory
	}

	return &category, nil
}

func (ac *CategoryCollection) UpdateOne(ctx context.Context, category domain.Category) (*domain.Category, error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "id", Value: category.ID}}
	update := bson.D{
		{
			Key: "$setOnInsert", Value: bson.D{},
		},
		{
			Key: "$set", Value: bson.D{
				{Key: "name", Value: category.Name},
				{Key: "path", Value: category.Path},
				{Key: "children", Value: category.Children},
				{Key: "created_at", Value: category.CreatedAt},
			},
		},
	}

	_, err := ac.mongoConnection.Connection.Collection(CategoryCollectionName).
		UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (ac *CategoryCollection) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	filter := bson.M{"id": id}
	category := &domain.Category{}

	err := ac.mongoConnection.Connection.Collection(CategoryCollectionName).FindOne(ctx, filter).Decode(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (ac *CategoryCollection) CreateMany(ctx context.Context, categories []domain.Category) ([]domain.Category, error) {
	categoriesData := []interface{}{}

	for _, data := range categories {
		categoriesData = append(categoriesData, data)
	}

	options := options.InsertMany().SetOrdered(true)

	_, err := ac.mongoConnection.Connection.Collection(CategoryCollectionName).InsertMany(ctx, categoriesData, options)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (ac *CategoryCollection) Get(ctx context.Context) ([]*domain.Category, error) {
	categories := []*domain.Category{}

	cur, err := ac.mongoConnection.Connection.Collection(CategoryCollectionName).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	cur.All(ctx, &categories)

	return categories, nil
}
