package user_repository

import (
	"api/src/database"
	models_user "api/src/models/user"
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

func (r *Repository) Insert(user models_user.User) (uint, error) {
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

func (r *Repository) Update(id uint, user models_user.User) error {
	defer r.db.Close()
	statements, err := r.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statements.Close()
	if _, err := statements.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id uint) error {
	defer r.db.Close()
	statements, err := r.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statements.Close()
	if _, err := statements.Exec(id); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAll(nickOrName string) ([]models_user.User, error) {
	defer r.db.Close()
	nickOrName = fmt.Sprintf("%%%s%%", nickOrName)
	rows, err := r.db.Query("SELECT id, name, nick, email, created_at FROM devbook.users WHERE name LIKE ? OR nick LIKE ?", nickOrName, nickOrName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models_user.User
	for rows.Next() {
		var user models_user.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) GetRandom() ([]models_user.User, error) {
	rows, err := r.db.Query(`
	SELECT u.id, u.name, u.nick, u.email, u.created_at
	FROM devbook.users u
	WHERE u.id NOT IN (
		SELECT f.follower_id
		FROM devbook.followers f
		WHERE f.user_id = 1
	) AND u.id != 1
	ORDER BY RAND()
	LIMIT 5
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models_user.User
	for rows.Next() {
		var user models_user.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) GetById(id uint) (models_user.User, error) {
	defer r.db.Close()
	rows, err := r.db.Query("SELECT id, name, nick, email, created_at FROM devbook.users WHERE id = ? ", id)
	if err != nil {
		return models_user.User{}, err
	}
	defer rows.Close()
	var user models_user.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models_user.User{}, err
		}
	}
	return user, nil
}

func (r *Repository) FollowUser(userID, followerID uint) error {
	statement, err := r.db.Prepare("INSERT ignore into devbook.followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}
	return nil
}

func (r *Repository) UnfollowUser(userID, followerID uint) error {
	statement, err := r.db.Prepare("DELETE from devbook.followers WHERE  user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserFollowers(userID uint) ([]models_user.User, error) {
	rows, err := r.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM devbook.users u
		INNER JOIN devbook.followers f ON u.id = f.follower_id
		WHERE f.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models_user.User
	for rows.Next() {
		var user models_user.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) GetUserFollowing(userID uint) ([]models_user.User, error) {
	rows, err := r.db.Query(`
		SELECT u.id, u.name, u.nick, u.email, u.created_at
		FROM devbook.users u
		INNER JOIN devbook.followers f ON u.id = f.user_id
		WHERE f.follower_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models_user.User
	for rows.Next() {
		var user models_user.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) GetPswrd(userID uint) (string, error) {
	rows, err := r.db.Query("SELECT pswrd FROM devbook.users WHERE id = ?", userID)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var user models_user.User
	for rows.Next() {
		if err := rows.Scan(&user.Pswrd); err != nil {
			return "", err
		}
	}
	return user.Pswrd, nil
}

func (r *Repository) UpdatePswrd(userID uint, pswrd string) error {
	statement, err := r.db.Prepare("UPDATE devbook.users SET pswrd = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(pswrd, userID); err != nil {
		return err
	}
	return nil
}
