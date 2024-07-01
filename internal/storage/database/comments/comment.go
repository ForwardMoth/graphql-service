package comments

import (
	"github.com/ForwardMoth/graphql-service/graph/models"
	"github.com/jmoiron/sqlx"
)

type CommentDatabase struct {
	db *sqlx.DB
}

func NewCommentDatabase(db *sqlx.DB) *CommentDatabase {
	return &CommentDatabase{db: db}
}

func (c *CommentDatabase) CreateComment(comment models.Comment) (models.Comment, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return models.Comment{}, err
	}

	query := `INSERT INTO comments (username, text, postId, commentId) 
				VALUES ($1, $2, $3, $4) RETURNING id`

	row := tx.QueryRow(query, comment.Username, comment.Text, comment.PostID, comment.CommentID)
	if err := row.Scan(&comment.ID); err != nil {
		err := tx.Rollback()
		if err != nil {
			return models.Comment{}, err
		}
		return models.Comment{}, err
	}

	return comment, tx.Commit()

}
