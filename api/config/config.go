package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoPostgres = ""
	Port                  = 0
	SecretKey             []byte
)

// Carregar() vai inicializar as variáveis de ambiente
func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	StringConexaoPostgres = fmt.Sprintf("host=localhost port=15432 user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

}

func ConectarDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", StringConexaoPostgres)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// 	//Criar usuario

// func CriarUsuarioDB() {
// 	db := ConectarDB()
// 	defer db.Close()
// }
