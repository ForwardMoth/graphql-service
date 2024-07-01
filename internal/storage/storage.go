package storage

import "github.com/ForwardMoth/graphql-service/graph/models"

type Storage struct {
	Posts
	Comments
}

func NewStorage(posts Posts, comments Comments) *Storage {
	return &Storage{
		Posts:    posts,
		Comments: comments,
	}
}

type Posts interface {
	CreatePost(post models.Post) (models.Post, error)
	GetPostById(id int) (models.Post, error)
	GetPosts(limit, offset int) ([]models.Post, error)
}

type Comments interface {
	CreateComment(comment models.Comment) (models.Comment, error)
}
