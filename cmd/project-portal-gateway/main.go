package main

import (
	"flag"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jeevangb/project-portal-gateway/internal/config"
	"github.com/jeevangb/project-portal-gateway/internal/graph"
	"github.com/jeevangb/project-portal-gateway/internal/server"
)

func main() {
	//Load env variables
	env := flag.String("env", "", "")
	flag.Parse()
	configdata, err := config.LoadConfig(env)
	if err != nil {
		return
	}
	//Initialize gin router
	router := gin.Default()
	//Initialize graphql server
	graph := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
	srv := handler.NewDefaultServer(graph)
	//This serves the GraphQL Playground interface
	router.GET("/", func(ctx *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(ctx.Writer, ctx.Request)
	})
	//This handles the GraphQL queries sent via POST requests.
	router.POST("/query", func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	})
	//start server
	server.SetUpServer(router, configdata.Port)
}
