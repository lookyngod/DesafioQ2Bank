package models

import "time"

type Usuario struct {
	ID      int64   `json:"id"`
	Nome    string  `json:"nome"`
	Email   string  `json:"email"`
	CPFCNPJ string  `json:"cpf_cnpj"`
	Saldo   float64 `json:"saldo"`
	Senha   string  `json:"senha"`
	isCFP   bool    `json:"is_cpf"`
	Tipo    string  `json:"tipo"`
}

type Transacao struct {
	ID        int64     `json:"id"`
	IDOrigem  Usuario   `json:"usuario"`
	IDDestino Usuario   `json:"usuario_destino"`
	Valor     float64   `json:"valor"`
	Data      time.Time `json:"data"`
}
