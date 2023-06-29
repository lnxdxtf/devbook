package repositories

import "database/sql"

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

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Insert(query string, args ...interface{}) (sql.Result, error) {
	return r.db.Exec(query, args...)
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
