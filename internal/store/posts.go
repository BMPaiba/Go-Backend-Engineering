package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserID    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreateAt  string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSTER INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.UserID,
		post.Tags,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreateAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
