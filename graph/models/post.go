package models

type PostDTO struct {
	ID          int    `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	IsCommented bool   `json:"isCommented"`
}

func (input PostInput) ToDto() PostDTO {
	return PostDTO{
		Author:      input.Author,
		Title:       input.Title,
		Text:        input.Text,
		IsCommented: input.IsCommented,
	}
}

func (dto PostDTO) ToModel() Post {
	return Post{
		Author:      dto.Author,
		Title:       dto.Title,
		Text:        dto.Text,
		IsCommented: dto.IsCommented,
	}
}

func ToArray(posts []Post) []*Post {
	newPosts := make([]*Post, len(posts))

	for i, post := range posts {
		newPosts[i] = &post
	}

	return newPosts
}
