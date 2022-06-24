package domains

import (
	"DESAFIOQ2BANK/api/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JSONMock struct {
	Authorization bool `json:"authorization"`
}

func ValidaMock(JSONMock JSONMock) bool {
	valida, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
	if err != nil {
		return false
	}

	val, _ := ioutil.ReadAll(valida.Body)

	err = json.Unmarshal(val, &JSONMock)
	if err != nil {
		return false
	}

	return JSONMock.Authorization

}

func BuscaTransacaoID(db *sql.DB, id string) (models.Transacao, error) {
	var row models.Transacao
	err := db.QueryRow("SELECT * FROM transacoes WHERE id=$1", id).Scan(&row.ID, &row.IDOrigem, &row.IDDestino, &row.Valor, &row.Data)
	if err != nil {
		return row, fmt.Errorf("falha na execução da busca de transação no postgres: %v", err)
	}
	return row, nil
}

func BuscaTodasTransacoes(db *sql.DB) ([]models.Transacao, error) {
	var registros []models.Transacao
	rows, err := db.Query("SELECT * FROM transacoes")
	if err != nil {
		return nil, fmt.Errorf("falha na execução da busca de todas as transações no postgres: %v", err)
	}
	for rows.Next() {
		var row models.Transacao
		err = rows.Scan(&row.ID, &row.IDOrigem, &row.IDDestino, &row.Valor, &row.Data)
		if err != nil {
			return nil, fmt.Errorf("falha na execução da busca de todas as transações no postgres: %v", err)
		}
		registros = append(registros, row)
	}
	return registros, nil
}
