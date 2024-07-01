package comment_cache

import (
	"github.com/ForwardMoth/graphql-service/graph/models"
	"sync"
)

type CommentCache struct {
	comments []models.Comment
	count    int
	mu       sync.Mutex
}

func NewCommentCache() *CommentCache {
	return &CommentCache{
		comments: make([]models.Comment, 0),
	}
}

func (c *CommentCache) CreateComment(comment models.Comment) (models.Comment, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.comments = append(c.comments, comment)
	c.count++

	return comment, nil
}
