package repositories

import (
	"database/sql"
)

type RepositoryDataLayer interface {
	Insert(db *sql.DB, query string, args ...interface{}) (sql.Result, error)
	Update(db *sql.DB, query string, args ...interface{}) (sql.Result, error)
	Delete(db *sql.DB, query string, args ...interface{}) (sql.Result, error)
	GetAll(db *sql.DB, dest interface{}, query string, args ...interface{}) error
	GetById(db *sql.DB, dest interface{}, query string, args ...interface{}) error
}