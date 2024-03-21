package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	dtocomment "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/comment"
	dtopost "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/post"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (pr *PostRepository) Insert(ctx context.Context, data entity.Post, userId string) error {
	sql := `INSERT INTO posts (id, user_id, content, tags) VALUES ($1,$2,$3,$4)`
	if _, err := pr.db.Exec(ctx, sql, data.ID, userId, data.PostInHtml, data.Tags); err != nil {
		return err
	}

	return nil
}

func (pr *PostRepository) FindById(ctx context.Context, postId string) (entity.Post, error) {

	post := entity.Post{}
	sql := `SELECT id, user_id, content, tags, created_at FROM posts WHERE posts.id = $1`
	if err := pr.db.QueryRow(ctx, sql, postId).Scan(&post.ID, &post.UserId, &post.PostInHtml, &post.Tags, &post.CreatedAt); err != nil {
		return post, err
	}

	return post, nil
}

func (pr *PostRepository) GetPostWithFilter(ctx context.Context, filter dtopost.PostFilter, userId string) ([]dtopost.Post, int64, error) {

	where := fmt.Sprintf("WHERE (friends.friend_id = '%s' or posts.user_id = '%s')", userId, userId)
	if filter.Search != "" {
		where += " AND posts.content LIKE '%" + filter.Search + "%'"
	}

	if len(filter.SearchTag) > 0 {
		jsonTag, err := json.Marshal([]string(filter.SearchTag))
		if err == nil {
			replacer := strings.NewReplacer("[", "{", "]", "}")
			stringTag := replacer.Replace(string(jsonTag))
			where += fmt.Sprintf(" AND posts.tags && '%s'", stringTag)
		}
	}

	sql := fmt.Sprintf(`SELECT 
	posts.id, 
	posts.content,
	posts.tags, 
	posts.created_at, 
	users.id, 
	users.name,
	users.image_url, 
	users.friend_count, 
	users.created_at,
	array(SELECT (comments.comment || ',' || comments.created_at || ',' || users.id || ','  || users.name || ','  || users.image_url || ','  || users.friend_count || ','  || users.created_at) FROM comments JOIN users ON comments.user_id = users.id WHERE posts.id = comments.post_id) as comments
	FROM posts 
	LEFT JOIN friends ON posts.user_id = friends.user_id
	JOIN users ON posts.user_id = users.id
	%s ORDER BY posts.created_at desc LIMIT %d OFFSET %d`, where, *filter.Limit, *filter.Offset)

	rows, err := pr.db.Query(ctx, sql)
	if err != nil {
		return []dtopost.Post{}, 0, err
	}

	data := make([]dtopost.Post, 0)
	var count int64 = 0
	for rows.Next() {
		var post entity.Post
		var creator entity.User
		var commentString []string
		// m := pgtype.NewMap()
		err := rows.Scan(
			&post.ID,
			&post.PostInHtml,
			&post.Tags,
			&post.CreatedAt,
			&creator.ID,
			&creator.Name,
			&creator.ImageUrl,
			&creator.FriendCount,
			&creator.CreatedAt,
			&commentString)
		if err != nil {
			return []dtopost.Post{}, 0, err
		}

		comments := make([]dtocomment.Comment, 0)
		for i := 0; i < len(commentString); i++ {
			var comment dtocomment.Comment
			commentArray := strings.Split(commentString[i], ",")

			comment.Comment = commentArray[0]
			comment.CreatedAt = commentArray[1]
			comment.Creator.ID = commentArray[2]
			comment.Creator.Name = commentArray[3]
			comment.Creator.ImageUrl = commentArray[4]
			comment.Creator.FriendCount, _ = strconv.ParseInt(commentArray[5], 10, 64)
			comment.CreatedAt = commentArray[6]

			comments = append(comments, comment)
		}

		data = append(data, dtopost.Post{
			ID:       post.ID,
			Comments: comments,
			Post:     post,
			Creator:  creator,
		})
		count += 1
	}
	rows.Close()

	return data, count, nil
}
