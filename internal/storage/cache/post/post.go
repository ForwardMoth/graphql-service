package post_cache

import (
	"github.com/ForwardMoth/graphql-service/graph/models"
	"github.com/ForwardMoth/graphql-service/internal/util/consts"
	"github.com/ForwardMoth/graphql-service/internal/util/error_handler"
	"sync"
)

type PostCache struct {
	posts []models.Post
	count int
	mu    sync.Mutex
}

func NewPostCache() *PostCache {
	return &PostCache{
		posts: make([]models.Post, 0),
	}
}

func (p *PostCache) CreatePost(post models.Post) (models.Post, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.posts = append(p.posts, post)
	p.count++

	return post, nil
}

func (p *PostCache) GetPostById(id int) (models.Post, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if id > p.count {
		return models.Post{}, error_handler.ResponseError{
			Message: consts.GettingPostError,
			Type:    consts.NotFound,
		}
	}

	return p.posts[id-1], nil
}

func (p *PostCache) GetPosts(limit, offset int) ([]models.Post, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if offset >= p.count {
		return nil, nil
	}

	if offset < 0 {
		return nil, error_handler.ResponseError{
			Message: consts.WrongOffsetError,
			Type:    consts.BadRequest,
		}
	}

	if limit != -1 && limit < 0 {
		return nil, error_handler.ResponseError{
			Message: consts.WrongLimitError,
			Type:    consts.BadRequest,
		}
	}

	if limit == -1 || offset+limit > p.count {
		return p.posts[offset:], nil
	}

	return p.posts[offset : offset+limit], nil
}
