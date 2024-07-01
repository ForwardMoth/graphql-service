package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ForwardMoth/graphql-service/graph/generated"
	"github.com/ForwardMoth/graphql-service/graph/resolvers"
	"github.com/ForwardMoth/graphql-service/internal/config"
	"github.com/ForwardMoth/graphql-service/internal/services"
	"github.com/ForwardMoth/graphql-service/internal/storage"
	"github.com/ForwardMoth/graphql-service/internal/storage/cache/comment"
	"github.com/ForwardMoth/graphql-service/internal/storage/cache/post"
	"github.com/ForwardMoth/graphql-service/internal/storage/database"
	"github.com/ForwardMoth/graphql-service/internal/storage/database/comments"
	"github.com/ForwardMoth/graphql-service/internal/storage/database/posts"
	"log"
	"net/http"
)

func main() {
	appConfig := config.LoadConfig()
	port := appConfig.ServerPort

	var storageService *storage.Storage
	if appConfig.InMemoryMode {
		post := post_cache.NewPostCache()
		comment := comment_cache.NewCommentCache()
		storageService = storage.NewStorage(post, comment)
	} else {
		db, err := database.Setup(appConfig.DBConfig)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connection to Postgres database is successful")
		defer db.DB.Close()

		post := posts.NewPostDatabase(db.DB)
		comment := comments.NewCommentDatabase(db.DB)
		storageService = storage.NewStorage(post, comment)
	}
	dataServices := services.NewService(storageService)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		PostService:     dataServices.Posts,
		CommentsService: dataServices.Comments,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
