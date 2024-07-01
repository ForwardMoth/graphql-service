package comments

import (
	"github.com/ForwardMoth/graphql-service/graph/models"
	mock_services "github.com/ForwardMoth/graphql-service/internal/services/mocks"
	"github.com/ForwardMoth/graphql-service/internal/util/consts"
	"github.com/ForwardMoth/graphql-service/internal/util/error_handler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCommentService_CreateComment(t *testing.T) {
	type input struct {
		com models.CommentDTO
	}

	type res struct {
		com    *models.Comment
		errCom error
	}

	inputComments := []models.CommentDTO{
		{
			ID:       1,
			Username: "Simon",
			Text:     "LoL",
			PostID:   1,
		},
		{
			ID:       1,
			Username: "",
			Text:     "LoL",
			PostID:   1,
		},
		{
			ID:       1,
			Username: "Simon",
			Text:     "",
			PostID:   1,
		},
	}

	wantComments := []*models.Comment{
		{
			ID:       "1",
			Username: "Simon",
			Text:     "LoL",
			PostID:   1,
		},
	}

	testTable := []struct {
		name  string
		input input
		want  res
	}{
		{
			name:  "OK",
			input: input{com: inputComments[0]},
			want:  res{com: wantComments[0], errCom: nil},
		},
		{
			name:  "No name",
			input: input{com: inputComments[1]},
			want: res{
				com:    &models.Comment{},
				errCom: error_handler.ResponseError{Message: consts.EmptyAuthorError}},
		},
		{
			name:  "No text",
			input: input{com: inputComments[2]},
			want: res{
				com:    &models.Comment{},
				errCom: error_handler.ResponseError{Message: consts.EmptyTextError}},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			c := mock_services.NewMockComments(ctl)
			c.EXPECT().CreateComment(tt.input.com).Return(tt.want.com, tt.want.errCom)

			got, err := c.CreateComment(tt.input.com)

			if err != nil {
				if !assert.Equal(t, tt.want.errCom, err) {
					log.Fatal(err)
				}
			}

			assert.Equal(t, tt.want.com, got)
		})
	}
}
