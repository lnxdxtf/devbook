package user_repository

import (
	"api/src/database"
	"api/src/models"
	"database/sql"
	"fmt"
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

func (r *Repository) Insert(user models.User) (uint, error) {
	query := "INSERT INTO devbook.users (name, nick, email, pswrd) VALUES (?, ?, ?, ?)"
	defer r.db.Close()
	statement, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Pswrd)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(lastId), nil
}

func (r *Repository) Update(query string, args ...interface{}) (sql.Result, error) {
	return r.db.Exec(query, args...)
}

func (r *Repository) Delete(query string, args ...interface{}) (sql.Result, error) {
	return r.db.Exec(query, args...)
}

func (r *Repository) GetAll(nickOrName string) ([]models.User, error) {
	nickOrName = fmt.Sprintf("%%%s%%", nickOrName)
	rows, err := r.db.Query("SELECT id, name, nick, email, created_at FROM devbook.users WHERE name LIKE ? OR nick LIKE ?", nickOrName, nickOrName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) Get(id uint) (models.User, error) {
	rows, err := r.db.Query("SELECT id, name, nick, email, created_at FROM devbook.users WHERE id LIKE ?", id)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()
	var user models.User
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
