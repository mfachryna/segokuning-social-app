package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

type UserRepository struct {
	*pgxpool.Pool
}

func NewUserRepo(pg *pgxpool.Pool) *UserRepository {
	return &UserRepository{pg}
}

func (ur *UserRepository) Get(ctx context.Context, data entity.User) error {
	return nil
}

func (ur *UserRepository) FindById(ctx context.Context, userId string) error {
	return nil
}

func (ur *UserRepository) Insert(ctx context.Context, data entity.User) error {
	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, userId string) error {
	return nil
}

func (ur *UserRepository) Update(ctx context.Context, data entity.User) error {
	return nil
}
