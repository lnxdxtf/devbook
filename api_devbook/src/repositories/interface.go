package repositories

import (
	"api/src/database"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

type DataLayer interface {
	Insert(db *sql.DB, query string, args ...interface{}) (sql.Result, error)
	Update(db *sql.DB, query string, args ...interface{}) (sql.Result, error)
	Delete(db *sql.DB, query string, args ...interface{}) (sql.Result, error)
	GetAll(db *sql.DB, dest interface{}, query string, args ...interface{}) error
	GetById(db *sql.DB, dest interface{}, query string, args ...interface{}) error
}

func NewRepository() (*Repository, error) {
	db_mysql := database.MySQLDB{}
	db, err := db_mysql.Connect()
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}

func (r *Repository) Insert(query string, args ...interface{}) (uint, error) {
	defer r.db.Close()
	statement, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(args...)
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

func (r *Repository) GetAll(dest interface{}, query string, args ...interface{}) error {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Scan rows into dest
	for rows.Next() {
		if err := rows.Scan(dest); err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) GetById(dest interface{}, query string, args ...interface{}) error {
	row := r.db.QueryRow(query, args...)
	if err := row.Scan(dest); err != nil {
		return err
	}

	return nil
}
