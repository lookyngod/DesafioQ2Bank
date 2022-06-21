package db

import (
	"DesafioQ2Bank/api/config"
	"database/sql"
)

func ConectarDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringConexaoPostgres)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
