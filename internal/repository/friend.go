package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/friend"
)

type FriendRepository struct {
	db *pgxpool.Pool
}

func NewFriendRepo(db *pgxpool.Pool) *FriendRepository {
	return &FriendRepository{
		db: db,
	}
}

func (ur *FriendRepository) Get(ctx context.Context, userId string) (dto.Friend, error) {
	return dto.Friend{}, nil
}

func (ur *FriendRepository) FindById(ctx context.Context, userId string, FriendId string) error {
	return nil
}

func (ur *FriendRepository) FindByField(ctx context.Context, FriendId string) error {
	return nil
}

func (ur *FriendRepository) FindByRelation(ctx context.Context, userId string, friendId string) (int, error) {
	var count int
	sql := `SELECT count(*) password FROM friends WHERE user_id = $1 and friend_id = $2`
	err := ur.db.QueryRow(ctx, sql, userId, friendId).Scan(&count)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return count, nil
}

func (ur *FriendRepository) Insert(ctx context.Context, userId string, friendId string) error {
	var sql string

	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

	sql = `INSERT INTO friends (user_id, friend_id, created_at) VALUES ($1,$2,$3),($4,$5,$6)`
	if _, err := ur.db.Exec(ctx, sql, userId, friendId, dt, friendId, userId, dt); err != nil {
		return err
	}

	return nil
}

func (ur *FriendRepository) Delete(ctx context.Context, FriendId string) error {
	return nil
}
