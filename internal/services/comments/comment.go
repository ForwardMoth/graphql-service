package comments

//go:generate mockgen -source=comment.go -destination=./mocks/mock_comment.go

import (
	"database/sql"
	"errors"
	"github.com/ForwardMoth/graphql-service/graph/models"
	"github.com/ForwardMoth/graphql-service/internal/storage"
	"github.com/ForwardMoth/graphql-service/internal/util/consts"
	"github.com/ForwardMoth/graphql-service/internal/util/error_handler"
)

type CommentService struct {
	service    storage.Comments
	PostGetter PostGetter
}

type PostGetter interface {
	GetPostById(id int) (models.Post, error)
}

const (
	maxUsernameLength = 64
	maxTextLength     = 2000
)

func NewCommentService(service storage.Comments, getter PostGetter) *CommentService {
	return &CommentService{service: service, PostGetter: getter}
}

func (c CommentService) CreateComment(comment models.CommentDTO) (*models.Comment, error) {
	if len(comment.Username) == 0 {
		return &models.Comment{}, error_handler.ResponseError{
			Message: consts.EmptyAuthorError,
			Type:    consts.BadRequest,
		}
	}

	if len(comment.Username) > maxUsernameLength {
		return &models.Comment{}, error_handler.ResponseError{
			Message: consts.TooMuchLengthAuthorError,
			Type:    consts.BadRequest,
		}
	}

	if len(comment.Text) == 0 {
		return &models.Comment{}, error_handler.ResponseError{
			Message: consts.EmptyTextError,
			Type:    consts.BadRequest,
		}
	}

	if len(comment.Text) > maxTextLength {
		return &models.Comment{}, error_handler.ResponseError{
			Message: consts.TooMuchLengthCommentError,
			Type:    consts.BadRequest,
		}
	}

	post, err := c.PostGetter.GetPostById(comment.PostID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.Comment{}, error_handler.ResponseError{
				Message: consts.PostNotFountError,
				Type:    consts.NotFound,
			}
		}
	}

	if !post.IsCommented {
		return &models.Comment{}, error_handler.ResponseError{
			Message: consts.CommentsNotAllowedError,
			Type:    consts.BadRequest,
		}

	}

	newComment, err := c.service.CreateComment(comment.ToModel())
	if err != nil {
		return &models.Comment{}, error_handler.ResponseError{
			Message: consts.CreatingCommentError,
			Type:    consts.InternalError,
		}
	}

	return &newComment, nil
}
