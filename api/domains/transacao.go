package domains

import (
	"DESAFIOQ2BANK/api/models"
	"database/sql"
	"fmt"
)

//QUERY PARA BUSCAR TRANSACAO POR ID

func BuscaTransacaoID(db *sql.DB, id string) (models.Transacao, error) {
	var row models.Transacao
	err := db.QueryRow("SELECT * FROM transacao WHERE id=$1", id).Scan(&row.ID, &row.IDOrigem, &row.IDDestino, &row.Valor)
	if err != nil {
		return row, fmt.Errorf("falha na execução da busca de transação no postgres: %v", err)
	}
	return row, nil
}

//QUERY PARA BUSCAR TODAS AS TRANSACOES

func BuscaTodasTransacoes(db *sql.DB) ([]models.Transacao, error) {
	var registros []models.Transacao
	rows, err := db.Query("SELECT * FROM transacao")
	if err != nil {
		return nil, fmt.Errorf("falha na execução da busca de todas as transações no postgres: %v", err)
	}
	for rows.Next() {
		var row models.Transacao
		err = rows.Scan(&row.ID, &row.IDOrigem, &row.IDDestino, &row.Valor)
		if err != nil {
			return nil, fmt.Errorf("falha na execução da busca de todas as transações no postgres: %v", err)
		}
		registros = append(registros, row)
	}
	return registros, nil
}

//QUERY PARA GRAVAR UMA TRANSAÇÃO

func InserirTransacao(db *sql.DB, T models.Transacao) error {
	_, err := db.Exec("INSERT INTO transacao VALUES ($1, $2, $3, $4)", T.ID, T.IDOrigem, T.IDDestino, T.Valor)
	if err != nil {
		return fmt.Errorf("falha na execução da inserção de transação no postgres: %v", err)
	}
	return nil
}
