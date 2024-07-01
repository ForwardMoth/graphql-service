package posts

import (
	"github.com/ForwardMoth/graphql-service/graph/models"
	mock_services "github.com/ForwardMoth/graphql-service/internal/services/mocks"
	"github.com/ForwardMoth/graphql-service/internal/util/consts"
	"github.com/ForwardMoth/graphql-service/internal/util/error_handler"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func longText() string {
	var sb strings.Builder
	for i := 0; i <= 100; i++ {
		sb.WriteString("aaaaaaaaaa")
	}
	return sb.String()
}

func TestCommentService_CreateComment(t *testing.T) {
	type input struct {
		com models.PostDTO
	}

	type res struct {
		com    *models.Post
		errCom error
	}

	inputPosts := []models.PostDTO{
		{
			ID:          1,
			Author:      "Simon",
			Title:       "new post",
			Text:        "some text",
			IsCommented: true,
		},
		{
			ID:          1,
			Author:      "",
			Title:       "new post",
			Text:        "LoL",
			IsCommented: true,
		},
		{
			ID:          1,
			Author:      "Simon",
			Title:       "Title",
			Text:        "",
			IsCommented: true,
		},
		{
			ID:          1,
			Author:      "Simon",
			Title:       "Title",
			Text:        longText(),
			IsCommented: true,
		},
	}

	wantPosts := []*models.Post{
		{
			ID:          "1",
			Author:      "Simon",
			Title:       "new post",
			Text:        "some text",
			IsCommented: true,
		},
	}

	testTable := []struct {
		name  string
		input input
		want  res
	}{
		{
			name:  "OK",
			input: input{com: inputPosts[0]},
			want:  res{com: wantPosts[0], errCom: nil},
		},
		{
			name:  "No author",
			input: input{com: inputPosts[1]},
			want: res{
				com:    &models.Post{},
				errCom: error_handler.ResponseError{Message: consts.EmptyAuthorError}},
		},
		{
			name:  "No text",
			input: input{com: inputPosts[2]},
			want: res{
				com:    &models.Post{},
				errCom: error_handler.ResponseError{Message: consts.EmptyTextError}},
		},
		{
			name:  "long text",
			input: input{com: inputPosts[3]},
			want: res{
				com:    &models.Post{},
				errCom: error_handler.ResponseError{Message: consts.TooMuchLengthTextError}},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			defer ctl.Finish()

			c := mock_services.NewMockPosts(ctl)
			c.EXPECT().CreatePost(tt.input.com).Return(tt.want.com, tt.want.errCom)

			got, err := c.CreatePost(tt.input.com)

			if err != nil {
				if !assert.Equal(t, tt.want.errCom, err) {
					log.Fatal(err)
				}
			}

			assert.Equal(t, tt.want.com, got)
		})
	}
}
