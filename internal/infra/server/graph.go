package server

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ManuelLecaro/erpcore/internal/infra/graph"
	"github.com/gin-gonic/gin"
)

func (a *APIServer) initGraphDependencies() {
	a.Graph = &graph.Resolver{}
}

// /api/erpcore/v1/cms
func (a *APIServer) initAPIRoutes(g *gin.RouterGroup) {
	apiGroup := g.Group("/cms")
	apiGroup.Any("/", gin.WrapH(playground.Handler("GraphQL playground", fmt.Sprintf("/api/%s/v1/cms/query", a.AppName))))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: a.Graph}))
	apiGroup.
		POST("/query", gin.WrapH(srv))
}
