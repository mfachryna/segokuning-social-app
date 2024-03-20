package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

type CommentRepository struct {
	db *pgxpool.Pool
}

func NewCommentRepo(db *pgxpool.Pool) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (pr *CommentRepository) Insert(ctx context.Context, data entity.Comment) error {
	sql := `INSERT INTO comments (id, user_id, comment, post_id) VALUES ($1,$2,$3,$4)`
	if _, err := pr.db.Exec(ctx, sql, data.ID, data.UserId, data.Comment, data.PostId); err != nil {
		return err
	}

	return nil
}
