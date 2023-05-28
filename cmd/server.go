package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AhmedEnnaime/GoQl/config"
	"github.com/AhmedEnnaime/GoQl/db"
	"github.com/AhmedEnnaime/GoQl/graph"
	"github.com/AhmedEnnaime/GoQl/graph/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func init() {
	config.LoadEnvVariables()
	db.ConnectDatabase()
}

func main() {
	type Question struct {
		model.Question
		gorm.Model
	}

	type Choice struct {
		model.Choice
		gorm.Model
	}

	config.LoadEnvVariables()
	db.ConnectDatabase()

	db.Model.AutoMigrate(&Question{})
	db.Model.AutoMigrate(&Choice{})

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
