package domains

import (
	"DESAFIOQ2BANK/api/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	ErroUsuarioDuplicado = "pq: duplicate key value violates unique constraint \"usuarios_pkey\""

	ERROIDNAOENCONTRADO = "sql: no rows in result set"
)

func ConectarDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=15432 user=postgres password=master dbname=postgresql sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

//QUERY PARA BUSCA DE TODOS OS CPF/CNPJ NO POSTGRES

func BuscaTodosUsuarios(db *sql.DB) ([]models.Usuario, error) {
	var registros []models.Usuario
	rows, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		return nil, fmt.Errorf("falha na execução da busca de todos os usuários no postgres: %v", err)
	}
	for rows.Next() {
		var row models.Usuario
		err = rows.Scan(&row.ID, &row.Nome, &row.Tipo, &row.CPFCNPJ, &row.Email, &row.Senha, &row.Saldo)
		if err != nil {
			return nil, fmt.Errorf("falha na execução da busca de todos os usuários no postgres: %v", err)
		}
		registros = append(registros, row)
	}
	return registros, nil
}

func BuscaUsuarioID(db *sql.DB, id string) (models.Usuario, error) {
	var row models.Usuario
	err := db.QueryRow("SELECT * FROM usuarios WHERE id=$1", id).Scan(&row.ID, &row.Nome, &row.Tipo, &row.CPFCNPJ, &row.Email, &row.Senha, &row.Saldo)
	if err != nil {
		return row, fmt.Errorf("falha na execução da busca de usuário no postgres: %v", err)
	}
	return row, nil

}

func AtualizaSaldo(db *sql.DB, id string, saldo float64) error {
	_, err := db.Exec("UPDATE usuarios SET saldo=$1 WHERE id=$2", saldo, id)
	if err != nil {
		return fmt.Errorf("falha na execução da atualização de saldo no postgres: %v", err)
	}
	return nil
}
