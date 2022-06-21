package domain

import "DesafioQ2Bank/api/models"

func NovaTransacao(IDOrigem int64, IDDestino int64, Valor float64) models.Transacao {
	return models.Transacao{IDOrigem: models.Usuario{ID: IDOrigem},
		IDDestino: models.Usuario{ID: IDDestino},
		Valor:     Valor}

}
