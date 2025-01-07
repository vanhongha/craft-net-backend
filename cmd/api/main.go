package main

import (
	"craftnet/config"
	"craftnet/graph"
	"craftnet/internal/app/directives"
	"craftnet/internal/app/middleware"
	"craftnet/internal/db"
	"craftnet/internal/util"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	// init logger
	util.InitLogger("../../logs/app.log")
	defer util.GetLogger().Close()

	// load config file
	config.LoadConfig("../../")

	// connect to DB
	db.ConnectDatabase()
	defer db.Instance.Close()

	// run graphQL
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)

	gc := graph.Config{Resolvers: &graph.Resolver{}}
	gc.Directives.Auth = directives.Auth

	srv := handler.New(graph.NewExecutableSchema(gc))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.AuthMiddleware(c.Handler(srv)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
