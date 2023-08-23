package post_repository

import (
	"api/src/database"
	models_post "api/src/models/post"

	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository() (*Repository, error) {
	db_mysql := database.MySQLDB{}
	db, err := db_mysql.Connect()
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

func (r *Repository) Insert(post models_post.Post) (uint, error) {
	statement, err := r.db.Prepare("INSERT INTO devbook.posts (title, content, author_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(postID), nil
}

func (r *Repository) GetById(id uint) (models_post.Post, error) {
	rows, err := r.db.Query(`
		SELECT p.id, p.title, p.content, p.author_id, p.likes, p.created_at, u.nick FROM devbook.posts p
		INNER JOIN devbook.users u ON u.id = p.author_id
		WHERE p.id = ?
	`, id)
	if err != nil {
		return models_post.Post{}, err
	}
	defer rows.Close()

	var post models_post.Post
	if rows.Next() {
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); err != nil {
			return models_post.Post{}, err
		}
	}
	return post, nil
}

func (r *Repository) GetAll(userID uint) ([]models_post.Post, error) {
	rows, err := r.db.Query(`
	SELECT distinct p.id, p.title, p.content, p.author_id, p.likes, p.created_at, u.nick from devbook.posts p
	INNER JOIN devbook.users u ON u.id = p.author_id
	LEFT JOIN devbook.followers f ON u.id = f.user_id
	WHERE u.id = ? OR f.follower_id = ?
	order by 1 desc
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models_post.Post
	for rows.Next() {
		var post models_post.Post
		if err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
