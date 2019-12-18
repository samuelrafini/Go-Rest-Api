package util

import (
	"database/sql"
)



func ConnectDB(connectionUrl string) (*sql.DB, error) {
	var db *sql.DB

	db, err := sql.Open("postgres", connectionUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
