package rest

import (
	"DESAFIOQ2BANK/api/domains/transport"
	"DESAFIOQ2BANK/api/utils"
	"net/http"
)

func UsuariosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		transport.BuscaTodosUsuarios(w, r)
		return
	}

	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")

}

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		transport.BuscaUsuarioID(w, r)
		return
	}
	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")
}

func TransacoesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		transport.BuscaTodasTransacoes(w, r)
		return
	}
