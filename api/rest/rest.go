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