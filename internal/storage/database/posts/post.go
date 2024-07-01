package posts

import (
	"github.com/ForwardMoth/graphql-service/graph/models"
	"github.com/jmoiron/sqlx"
)

type PostDatabase struct {
	db *sqlx.DB
}

func NewPostDatabase(db *sqlx.DB) *PostDatabase {
	return &PostDatabase{db: db}
}

func (p *PostDatabase) CreatePost(post models.Post) (models.Post, error) {
	query := `INSERT INTO Posts (author, title, text, isCommented) 
				VALUES ($1, $2, $3, $4)
				RETURNING id`

	tx, err := p.db.Begin()
	if err != nil {
		return models.Post{}, err
	}

	row := tx.QueryRow(query, post.Author, post.Title, post.Text, post.IsCommented)
	if err := row.Scan(&post.ID); err != nil {
		err := tx.Rollback()
		if err != nil {
			return models.Post{}, err
		}
		return models.Post{}, err
	}

	return post, tx.Commit()
}

func (p *PostDatabase) GetPosts(limit, offset int) ([]models.Post, error) {
	query := "SELECT * FROM posts ORDER BY id OFFSET $1"
	args := []interface{}{offset}

	if limit > 0 {
		query += " LIMIT $2"
		args = append(args, limit)
	}

	var posts []models.Post

	if err := p.db.Select(&posts, query, args...); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostDatabase) GetPostById(id int) (models.Post, error) {
	query := `SELECT * FROM posts WHERE id = $1`

	var post models.Post

	if err := p.db.Get(&post, query, id); err != nil {
		return models.Post{}, err
	}

	return post, nil
}
