package domains

import (
	"DESAFIOQ2BANK/api/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	ErroCPFCNPJDuplicado = "pq: duplicate key value violates unique constraint \"usuarios_pkey\""
)

func NovoUsuario(db *sql.DB, usu models.Usuario) error {
	sqlStatement := fmt.Sprintf("INSERT INTO Usuarios(nome, email, cpfcnpj, senha) VALUES('%s', '%s','%s','%s');", usu.Nome, usu.Email, usu.CPFCNPJ, usu.Senha)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		if err.Error() == ErroCPFCNPJDuplicado {
			return fmt.Errorf("Usuário já cadastrado na tabela")
		}
		return fmt.Errorf("falha na execução do insert de CPF/CNPJ no postgres: %v", err)
	}
	return nil
}

//QUERY PARA BUSCA DE TODOS OS CPF/CNPJ NO POSTGRES

func BuscaTodosUsuarios(db *sql.DB) ([]models.Usuario, error) {
	var registros []models.Usuario
	rows, err := db.Query("SELECT * FROM Usuarios;")
	if err != nil {
		return nil, fmt.Errorf("falha na execução da busca de todos os CPSAPK no postgres: %v", err)
	}
	for rows.Next() {
		var row models.Usuario
		err = rows.Scan(&row.ID, &row.Nome, &row.Tipo, &row.Saldo)
		if err != nil {
			return nil, fmt.Errorf("falha na execução da busca de todos os cpf no postgres: %v", err)
		}
		registros = append(registros, row)
	}
	return registros, nil
}

func DeleteUsuarioPorID(db *sql.DB, id string) error {
	_, err := db.Query("DELETE FROM Usuarios WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("falha na execução da busca de cpf no postgres: %v", err)
	}
	return nil
}

func BuscaUsuarioPorID(db *sql.DB, id string) (models.Usuario, error) {
	var row models.Usuario
	err := db.QueryRow("SELECT ID FROM Usuarios WHERE ID=$1", id).Scan(&row.ID)
	if err != nil {
		return row, fmt.Errorf("falha na execução da busca de id no postgres: %v", err)
	}
	return row, nil

}

func InsereUsuario(db *sql.DB, usu models.Usuario) error {
	sqlStatement := fmt.Sprintf("INSERT INTO Usuario(Nome, Email, CPFCNPJ, Senha) VALUES('%s','%s','%s','%s');", usu.Nome, usu.Email, usu.CPFCNPJ, usu.Senha)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		if err.Error() == ErroCPFCNPJDuplicado {
			return fmt.Errorf("CPF/CNPJ já cadastrado")
		}
		return fmt.Errorf("falha na execução do insert de CPF/CNPJ no postgres: %v", err)
	}
	return nil
}
