package rest

import (
	"DESAFIOQ2BANK/api/domains/transport"
	"DESAFIOQ2BANK/api/models"
	"DESAFIOQ2BANK/api/services"
	"DESAFIOQ2BANK/api/utils"
	"net/http"
)

var (
	tran = models.Transacao{}
)

//Handler para buscar todos os usuários

func UsuariosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		transport.BuscaTodosUsuarios(w, r)
		return
	}

	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")

}

//Handler para buscar usuário por ID

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		transport.BuscaUsuarioID(w, r)
		return
	}
	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")
}

//Handler para buscar todas as transações

func TransacoesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		transport.BuscaTodasTransacoes(w, r)
		return
	}
	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")

}

//Handler para buscar transação por ID e fazer uma nova transação

func TransacaoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		transport.BuscaTransacaoID(w, r)
		return
	}

	if r.Method == http.MethodPost {
		services.NovaTransacao(tran)
		return
	}
	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")
}
