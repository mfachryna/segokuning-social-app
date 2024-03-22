package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
	"go.uber.org/zap"
)

type CommentRepository struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewCommentRepo(db *pgxpool.Pool, log *zap.Logger) *CommentRepository {
	return &CommentRepository{
		db:  db,
		log: log,
	}
}

func (pr *CommentRepository) Insert(ctx context.Context, data entity.Comment) error {
	sql := `INSERT INTO comments (id, user_id, comment, post_id) VALUES ($1,$2,$3,$4)`
	if _, err := pr.db.Exec(ctx, sql, data.ID, data.UserId, data.Comment, data.PostId); err != nil {
		return err
	}

	return nil
}
