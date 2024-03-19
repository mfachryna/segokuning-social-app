package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Get(ctx context.Context, data entity.User) error {
	return nil
}

func (ur *UserRepository) FindById(ctx context.Context, userId string) error {
	return nil
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	res := &entity.User{}
	sql := `SELECT id, name, email, COALESCE(phone, ''), password FROM users WHERE email = $1`

	err := ur.db.QueryRow(ctx, sql, email).Scan(&res.ID, &res.Name, &res.Email, &res.Phone, &res.Password)
	if err != nil {
		fmt.Println("AWAW", err)
		return nil, err
	}

	return res, nil
}

func (ur *UserRepository) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	res := &entity.User{}
	sql := `SELECT id, name, COALESCE(email, ''), phone, password FROM users WHERE phone = $1`

	err := ur.db.QueryRow(ctx, sql, phone).Scan(&res.ID, &res.Name, &res.Email, &res.Phone, &res.Password)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (ur *UserRepository) FindByField(ctx context.Context, userId string) error {
	return nil
}

func (ur *UserRepository) Insert(ctx context.Context, data entity.User, credType string) error {
	var sql string

	switch credType {
	case "phone":
		sql = `INSERT INTO users (id,phone,name,password,friend_count,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`
		if _, err := ur.db.Exec(ctx, sql, data.ID, data.Phone, data.Name, data.Password, 0, data.CreatedAt, data.CreatedAt); err != nil {
			return err
		}
	case "email":
		sql = `INSERT INTO users (id,email,name,password,friend_count,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`
		if _, err := ur.db.Exec(ctx, sql, data.ID, data.Email, data.Name, data.Password, 0, data.CreatedAt, data.CreatedAt); err != nil {
			return err
		}
	}

	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, userId string) error {
	return nil
}

func (ur *UserRepository) Update(ctx context.Context, data entity.User) error {
	return nil
}

func (ur *UserRepository) EmailCheck(ctx context.Context, email string) (int64, error) {
	var count int64

	if err := ur.db.QueryRow(ctx, "SELECT COUNT(id) FROM users WHERE email = $1", email).Scan(&count); err != nil {
		return 0, nil
	}

	return count, nil
}

func (ur *UserRepository) PhoneCheck(ctx context.Context, phone string) (int64, error) {
	var count int64

	if err := ur.db.QueryRow(ctx, "SELECT COUNT(id) FROM users WHERE phone = $1", phone).Scan(&count); err != nil {
		return 0, nil
	}

	return count, nil
}
