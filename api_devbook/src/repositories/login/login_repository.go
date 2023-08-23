package login_repository

import (
	"api/src/database"
	models_user "api/src/models/user"
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

func (r *Repository) GetUserByEmail(email string) (models_user.User, error) {
	row, err := r.db.Query("SELECT id, name, nick, pswrd, email, created_at FROM devbook.users WHERE email = ?", email)
	if err != nil {
		return models_user.User{}, err
	}
	defer row.Close()
	var user models_user.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Nick, &user.Pswrd, &user.Email, &user.CreatedAt); err != nil {
			return models_user.User{}, err
		}
	}
	return user, nil
}
