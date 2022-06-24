package services

import (
	"DESAFIOQ2BANK/api/domains"
	"DESAFIOQ2BANK/api/models"
	"fmt"
)

var (
	loj, com models.Usuario
	err      error
)

func NovaTransacao(T models.Transacao) (string, error) {
	db, _ := domains.ConectarDB()
	mock, err := domains.ValidaMock()
	if err != nil {
		return "Erro na Validação do Mock", err
	}

	if mock != true {
		return "Mock não autorizado", err
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
	if valid != true {
		valid := fmt.Sprintf("Não autorizado, saldo insuficiente:%", c)
		return valid, err
	}

	err = domains.InserirTransacao(T)
	if err != nil {
		return "Erro ao inserir transação", err
	}

	loj.Saldo -= T.Valor
	com.Saldo += T.Valor

	err = domains.AtualizaSaldo(loj)
	if err != nil {
		return "Erro ao atualizar saldo", err
	}

	err = domains.AtualizaSaldo(com)
	if err != nil {
		return "Erro ao atualizar saldo", err
	}

	return "Transação realizada com sucesso", nil

}

func ValidarSaldo(usu models.Usuario, Saldo float64) (bool, string) {

	if usu.Tipo != "Comum" {
		return false, "Usuário não é comum"

	}

	if usu.Saldo < Saldo {
		return false, "Saldo insuficiente"
	}

	return true, ""

}
