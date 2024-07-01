package posts

import (
	"database/sql"
	"errors"
	"github.com/ForwardMoth/graphql-service/graph/models"
	"github.com/ForwardMoth/graphql-service/internal/storage"
	"github.com/ForwardMoth/graphql-service/internal/util/consts"
	"github.com/ForwardMoth/graphql-service/internal/util/error_handler"
	"log"
)

type PostService struct {
	service storage.Posts
}

const (
	maxAuthorNameLength = 64
	maxTextLength       = 3000
)

func NewPostService(service storage.Posts) *PostService {
	return &PostService{service: service}
}

func (ps PostService) CreatePost(post models.PostDTO) (*models.Post, error) {
	if len(post.Author) == 0 {
		log.Print(consts.EmptyAuthorError)
		return &models.Post{}, error_handler.ResponseError{
			Message: consts.EmptyAuthorError,
			Type:    consts.BadRequest,
		}
	}

	if len(post.Author) > maxAuthorNameLength {
		log.Print(consts.TooMuchLengthAuthorError)
		return &models.Post{}, error_handler.ResponseError{
			Message: consts.TooMuchLengthAuthorError,
			Type:    consts.BadRequest,
		}
	}

	if len(post.Text) == 0 {
		log.Print(consts.EmptyTextError)
		return &models.Post{}, error_handler.ResponseError{
			Message: consts.EmptyTextError,
			Type:    consts.BadRequest,
		}
	}

	if len(post.Text) > maxTextLength {
		log.Print(consts.TooMuchLengthTextError)
		return &models.Post{}, error_handler.ResponseError{
			Message: consts.TooMuchLengthTextError,
			Type:    consts.BadRequest,
		}
	}

	newPost, err := ps.service.CreatePost(post.ToModel())
	if err != nil {
		log.Print(err)
		return &models.Post{}, error_handler.ResponseError{
			Message: consts.CreatingPostError,
			Type:    consts.InternalError,
		}
	}

	return &newPost, nil
}

func (ps PostService) GetPosts(limit, offset *int) ([]models.Post, error) {
	if limit != nil && *limit < 0 {
		return nil, error_handler.ResponseError{
			Message: consts.WrongLimitError,
			Type:    consts.BadRequest,
		}
	}

	if limit != nil && *offset < 0 {
		return nil, error_handler.ResponseError{
			Message: consts.WrongOffsetError,
			Type:    consts.BadRequest,
		}
	}

	posts, err := ps.service.GetPosts(*limit, *offset)
	if err != nil {
		return nil, error_handler.ResponseError{
			Message: consts.GettingPostError,
			Type:    consts.InternalError,
		}
	}

	return posts, nil
}

func (ps PostService) GetPostById(id int) (*models.Post, error) {
	if id <= 0 {
		return &models.Post{}, error_handler.ResponseError{
			Message: consts.WrongIdError,
			Type:    consts.BadRequest,
		}
	}

	post, err := ps.service.GetPostById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.Post{}, error_handler.ResponseError{
				Message: consts.PostNotFountError,
				Type:    consts.NotFound,
			}
		}
		return &models.Post{}, error_handler.ResponseError{
			Message: consts.GettingPostError,
			Type:    consts.InternalError,
		}
	}

	return &post, nil
}
