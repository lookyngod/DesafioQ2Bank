package models

import (
	"time"
)

type Usuario struct {
	ID      int64  `json:"id"`
	Nome    string `json:"nome"`
	Email   string `json:"email"`
	CPFCNPJ string `json:"cpf_cnpj"`
	Saldo   string `json:"saldo"`
	Senha   string `json:"senha"`
	Tipo    string `json:"tipo"`
}

type Transacao struct {
	ID        int64     `json:"id"`
	IDOrigem  string    `json:"usuario"`
	IDDestino string    `json:"usuario_destino"`
	Valor     string    `json:"valor"`
	Data      time.Time `json:"data"`
}

// func QueryInserir(transacao Transacao) string {
// 	fmtquery := fmt.Sprintf("INSERT INTO TRANSACTION (Valor, IDOrigem, IDDestino) VALUES (%s, %s, %s)", transacao.Valor, transacao.IDOrigem, transacao.IDDestino)
// 	return fmtquery
// }
