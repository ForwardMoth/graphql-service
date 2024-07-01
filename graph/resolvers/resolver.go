package resolvers

import "github.com/ForwardMoth/graphql-service/internal/services"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	PostService     services.Posts
	CommentsService services.Comments
}
