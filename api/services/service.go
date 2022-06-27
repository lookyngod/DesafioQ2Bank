package services

import (
	"DESAFIOQ2BANK/api/domains"
	"DESAFIOQ2BANK/api/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	loj, com models.Usuario
)

type JSONMock struct {
	Authorization bool `json:"authorization"`
}

//Validações para uma transação aprovada

func NovaTransacao(T models.Transacao) (string, error) {
	db, _ := domains.ConectarDB()
	mock, err := ValidaMock()
	if err != nil {
		return "Erro na Validação do Mock", err
	}

	if !mock {
		return "Erro na Inserção da Transação", err
	}

	loj, err = domains.BuscaUsuarioID(db, T.IDOrigem)
	if err != nil {
		return "Erro ao buscar usuário", err
	}

	com, err = domains.BuscaUsuarioID(db, T.IDDestino)
	if err != nil {
		return "Erro ao buscar usuário", err
	}

	valid, c := ValidarSaldo(loj, T.Valor)

	if !valid {
		valid := fmt.Sprintf("Não autorizado, saldo insuficiente: %v", c)
		return valid, err
	}

	err = domains.InserirTransacao(db, T)
	if err != nil {
		return "Erro ao inserir transação", err
	}

	loj.Saldo -= T.Valor
	com.Saldo += T.Valor

	err = domains.AtualizaSaldo(db, loj)
	if err != nil {
		return "Erro ao atualizar saldo", err
	}

	err = domains.AtualizaSaldo(db, com)
	if err != nil {
		return "Erro ao atualizar saldo", err
	}

	return "Transação realizada com sucesso", nil

}

// Validação de saldo e qual tipo de usuário

func ValidarSaldo(usu models.Usuario, Saldo float64) (bool, string) {

	if usu.Tipo != "Comum" {
		return false, "Usuário não é comum"

	}

	if usu.Saldo < Saldo {
		return false, "Saldo insuficiente"
	}

	return true, ""

}

// Validação da autorização mock

func ValidaMock() (bool, error) {
	var jsonMock JSONMock

	valida, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
	if err != nil {
		return false, err

	}

	err = json.NewDecoder(valida.Body).Decode(&jsonMock)
	if err != nil {
		return false, err
	}

	return jsonMock.Authorization, nil
}
