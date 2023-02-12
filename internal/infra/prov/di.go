package prov

import (
	"context"

	"github.com/ManuelLecaro/erpcore/internal/core/article"
	"github.com/ManuelLecaro/erpcore/internal/core/category"
	"github.com/ManuelLecaro/erpcore/internal/core/transaction"
	"github.com/ManuelLecaro/erpcore/internal/infra/configuration"
	"github.com/ManuelLecaro/erpcore/internal/infra/mongo"
	"github.com/ManuelLecaro/erpcore/internal/infra/mongo/collections"
	"github.com/ManuelLecaro/erpcore/internal/infra/server"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() (*dig.Container, error) {
	Container = dig.New()

	configs := configuration.LoadConfiguration()

	dbName := configs.GetString("DB_NAME")

	err := Container.Provide(func() *mongo.MongoConnection {
		conn, err := mongo.NewMongoConnection(context.Background(), dbName)
		if err == nil {
			return conn
		}

		return nil
	})

	err = Container.Provide(server.NewApp)
	err = Container.Provide(collections.NewArticleCollection)
	err = Container.Provide(article.NewArticleService)

	err = Container.Provide(collections.NewCategoryCollection)
	err = Container.Provide(category.NewCategoryService)

	err = Container.Provide(collections.NewTransactionCollection)
	err = Container.Provide(transaction.NewTransactionService)

	err = Container.Provide(server.NewGraphHandler)

	return Container, err
}
