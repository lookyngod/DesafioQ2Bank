package models

type Usuario struct {
	ID      int64   `json:"id"`
	Nome    string  `json:"nome"`
	Email   string  `json:"email"`
	CPFCNPJ string  `json:"cpf_cnpj"`
	Saldo   float64 `json:"saldo"`
	Senha   string  `json:"senha"`
}

type Money int64

func (m Money) Float64() float64 {
	return float64(m) / 100
}

func (m Money) Int64() int64 {
	return int64(m)
}
