package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToPostgreSQL(dbString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL database")

	return db, nil
}
