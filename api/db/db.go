package db

import "database/sql"

func ConectarDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=15432 user=postgres password=master dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
