package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(user, password, host string, port int, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	// Проверка подключения к БД
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
