package domain

import (
	"DesafioQ2Bank/api/models"
)

func NovoUsuario(ID int64, Nome string, Tipo string, CPFCNPJ string, Email string, Senha string, Saldo float64) models.Usuario {
	return models.Usuario{ID: ID, Nome: Nome, Tipo: Tipo, CPFCNPJ: CPFCNPJ, Email: Email, Senha: Senha, Saldo: Saldo}
}
