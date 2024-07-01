package services

//go:generate mockgen -source=service.go -destination=mocks/mock.go

import (
	"github.com/ForwardMoth/graphql-service/graph/models"
	"github.com/ForwardMoth/graphql-service/internal/services/comments"
	"github.com/ForwardMoth/graphql-service/internal/services/posts"
	"github.com/ForwardMoth/graphql-service/internal/storage"
)

type Service struct {
	Posts
	Comments
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		Posts:    posts.NewPostService(storage.Posts),
		Comments: comments.NewCommentService(storage.Comments, storage.Posts),
	}
}

type Posts interface {
	CreatePost(post models.PostDTO) (*models.Post, error)
	GetPostById(id int) (*models.Post, error)
	GetPosts(limit, offset *int) ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.CommentDTO) (*models.Comment, error)
}
