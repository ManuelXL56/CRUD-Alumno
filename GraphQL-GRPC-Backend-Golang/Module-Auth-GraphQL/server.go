package main

import (
	"fmt"
	"net/http"
	"time"

	"Module.com/GraphQL-Server/graph"
	"Module.com/GraphQL-Server/graph/generated"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	cors "github.com/itsjamie/gin-cors"
)

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/graphql/auth")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "0.0.0.0"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "http://localhost:8080, http://localhost:3000, http://localhost:3001, http://localhost, https://localhost",
		Methods:         "POST",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: true,
	}))

	router.GET("/", playgroundHandler())
	router.POST("/graphql/auth", graphqlHandler())

	fmt.Println("Listen in port 8080")
	router.Run(":8080")
}
