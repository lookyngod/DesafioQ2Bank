package transport

import (
	"DESAFIOQ2BANK/api/domains"
	"DESAFIOQ2BANK/api/models"
	"DESAFIOQ2BANK/api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	reqT = models.Transacao{}
)

func BuscaTodosUsuarios(w http.ResponseWriter, r *http.Request) {
	pg, err := domains.ConectarDB()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	dados, err := domains.BuscaTodosUsuarios(pg)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, dados)
}

func BuscaUsuarioID(w http.ResponseWriter, r *http.Request) {
	pg, err := domains.ConectarDB()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	id := r.FormValue("id")

	dados, err := domains.BuscaUsuarioID(pg, id)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, dados)
}

func BuscaTodasTransacoes(w http.ResponseWriter, r *http.Request) {
	pg, err := domains.ConectarDB()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	dados, err := domains.BuscaTodasTransacoes(pg)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, dados)
}

func BuscaTransacaoID(w http.ResponseWriter, r *http.Request) {
	pg, err := domains.ConectarDB()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	id := mux.Vars(r)["id"]

	dados, err := domains.BuscaTransacaoID(pg, id)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, dados)
}

func InserirTransacao(w http.ResponseWriter, r *http.Request) {
	pg, err := domains.ConectarDB()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	resp, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(resp, &reqT)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}

	err = domains.InserirTransacao(pg, reqT)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Transação inserida com sucesso")
}
