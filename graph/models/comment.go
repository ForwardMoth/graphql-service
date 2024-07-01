package models

type CommentDTO struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	PostID    int    `json:"postID"`
	CommentID *int   `json:"commentID"`
}

func (input CommentInput) ToDto() CommentDTO {
	return CommentDTO{
		Username:  input.Username,
		Text:      input.Text,
		PostID:    input.PostID,
		CommentID: input.CommentID,
	}
}

func (dto CommentDTO) ToModel() Comment {
	return Comment{
		Username: dto.Username,
		Text:     dto.Text,
		PostID:   dto.PostID,
	}
}
