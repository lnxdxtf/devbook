package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct{}

type MySQLDBInterface interface {
	Connect() (*sql.DB, error)
}

func (db *MySQLDB) Connect() (*sql.DB, error) {
	db_conn, err := sql.Open("mysql", config.MySQLStrConn)
	if err != nil {
		return nil, err
	}
	if err = db_conn.Ping(); err != nil {
		db_conn.Close()
		return nil, err
	}

	return db_conn, nil
}
