package server

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ManuelLecaro/erpcore/internal/core/ports/services"
	"github.com/ManuelLecaro/erpcore/internal/infra/graph"
	"github.com/gin-gonic/gin"
)

type GraphHandler struct {
	graph           *graph.Resolver
	articleService  services.IArticleService
	categoryService services.ICategoryService
	server          *APIServer
}

func NewGraphHandler(
	server *APIServer,
	articleService services.IArticleService,
	categoryService services.ICategoryService,
) *GraphHandler {
	return &GraphHandler{
		graph: &graph.Resolver{
			ArticleService:  articleService,
			CategoryService: categoryService,
		},
		articleService:  articleService,
		categoryService: categoryService,
		server:          server,
	}
}

func (gh *GraphHandler) Init(appName string) {
	gh.graph = &graph.Resolver{
		ArticleService:  gh.articleService,
		CategoryService: gh.categoryService,
	}

	gh.server.Server.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{
			SkipPaths: []string{
				fmt.Sprintf("/api/%s/health-check", appName),
			},
		}),
		gin.Recovery(),
	)

	baseGroup := gh.server.Server.Group(fmt.Sprintf("/api/%s", appName))
	apiGroup := baseGroup.Group("/v1")

	gh.initAPIRoutes(apiGroup, appName)
}

// /api/erpcore/v1/cms
func (gh *GraphHandler) initAPIRoutes(g *gin.RouterGroup, appName string) {
	apiGroup := g.Group("/cms")

	apiGroup.Any("/",
		gin.WrapH(playground.Handler("GraphQL playground", fmt.Sprintf("/api/%s/v1/cms/query", appName))),
	)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: gh.graph}))

	apiGroup.POST("/query", gin.WrapH(srv))
}
