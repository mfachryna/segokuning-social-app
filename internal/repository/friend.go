package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type FriendRepository struct {
	db  *pgxpool.Pool
	log *zap.Logger
}

func NewFriendRepo(db *pgxpool.Pool, log *zap.Logger) *FriendRepository {
	return &FriendRepository{
		db:  db,
		log: log,
	}
}

func (ur *FriendRepository) FindByRelation(ctx context.Context, userId, friendId string) (int, error) {
	var count int
	sql := `SELECT count(user_id) password FROM friends WHERE user_id = $1 and friend_id = $2`
	err := ur.db.QueryRow(ctx, sql, userId, friendId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (ur *FriendRepository) Insert(ctx context.Context, userId, friendId string) error {
	var sql string

	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

	sql = `INSERT INTO friends (user_id, friend_id, created_at) VALUES ($1,$2,$3),($4,$5,$6)`
	if _, err := ur.db.Exec(ctx, sql, userId, friendId, dt, friendId, userId, dt); err != nil {
		return err
	}

	userSql := `UPDATE users SET friend_count = friend_count + 1 WHERE (id = $1 or id = $2)`
	if _, err := ur.db.Exec(ctx, userSql, userId, friendId); err != nil {
		return err
	}

	return nil
}

func (ur *FriendRepository) Delete(ctx context.Context, userId, friendId string) error {
	sql := `DELETE from friends where (user_id = $2 and friend_id = $1) or (user_id = $1 and friend_id = $2)`
	if _, err := ur.db.Exec(ctx, sql, userId, friendId); err != nil {
		return err
	}
	userSql := `UPDATE users SET friend_count = friend_count - 1 WHERE (id = $1 or id = $2)`
	if _, err := ur.db.Exec(ctx, userSql, userId, friendId); err != nil {
		return err
	}

	return nil
}
