package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/user"
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

func (ur *UserRepository) FindById(ctx context.Context, userId string) (*entity.User, error) {
	res := &entity.User{}
	sql := `SELECT id, name, COALESCE(email, ''), COALESCE(phone, ''), password FROM users WHERE id = $1`

	err := ur.db.QueryRow(ctx, sql, userId).Scan(&res.ID, &res.Name, &res.Email, &res.Phone, &res.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	res := &entity.User{}
	sql := `SELECT id, name, email, COALESCE(phone, ''), password FROM users WHERE email = $1`

	err := ur.db.QueryRow(ctx, sql, email).Scan(&res.ID, &res.Name, &res.Email, &res.Phone, &res.Password)
	if err != nil {
		fmt.Println(err)
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

func (ur *UserRepository) GetUserWithFilter(ctx context.Context, userId string, filter dto.UserFilter) ([]entity.User, int64, error) {
	sort := "users.created_at"
	if !(filter.SortBy == "") && (filter.SortBy != "createdAt") {
		sort = "users.friend_count"
	}
	order := "asc"
	if !(filter.OrderBy == "") {
		order = filter.OrderBy
	}
	where := fmt.Sprintf(" WHERE users.id <> '%s'", userId)
	join := ""
	if filter.OnlyFriend {
		where += fmt.Sprintf(" AND friends.user_id = '%s'", userId)
		join = " JOIN friends ON users.id = friends.friend_id"
	}

	if filter.Search != "" {
		where += " AND users.name LIKE '%" + filter.Search + "%'"
	}

	rows, err := ur.db.Query(ctx, fmt.Sprintf("SELECT users.id, users.name, users.image_url, users.friend_count, users.created_at FROM users %s %s ORDER BY %s %s LIMIT %d OFFSET %d", join, where, sort, order, *filter.Limit, *filter.Offset))
	if err != nil {
		return []entity.User{}, 0, err
	}

	data := make([]entity.User, 0)
	var count int64 = 0
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.ImageUrl, &user.FriendCount, &user.CreatedAt)
		if err != nil {
			return []entity.User{}, 0, err
		}
		data = append(data, user)
		count += 1
	}
	rows.Close()

	return data, count, nil
}

func (ur *UserRepository) Insert(ctx context.Context, data entity.User, credType string) error {
	var sql string

	switch credType {
	case "phone":
		sql = `INSERT INTO users (id,phone,name,password,friend_count,created_at) VALUES ($1,$2,$3,$4,$5,$6)`
		if _, err := ur.db.Exec(ctx, sql, data.ID, data.Phone, data.Name, data.Password, 0, data.CreatedAt); err != nil {
			return err
		}
	case "email":
		sql = `INSERT INTO users (id,email,name,password,friend_count,created_at) VALUES ($1,$2,$3,$4,$5,$6)`
		if _, err := ur.db.Exec(ctx, sql, data.ID, data.Email, data.Name, data.Password, 0, data.CreatedAt); err != nil {
			return err
		}
	}

	return nil
}

func (ur *UserRepository) Delete(ctx context.Context, userId string) error {
	return nil
}

func (ur *UserRepository) Update(ctx context.Context, data entity.User) error {
	sql := `UPDATE users SET name = $1, email = $2, phone = $3, password = $4, image_url = $5, friend_count = $6 WHERE id = $7`

	_, err := ur.db.Exec(ctx, sql, data.Name, data.Email, data.Phone, data.Password, data.ImageUrl, data.FriendCount, data.ID)
	if err != nil {
		return err
	}
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
