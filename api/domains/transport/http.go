package transport

import (
	"DESAFIOQ2BANK/api/db"
	"DESAFIOQ2BANK/api/domains"
	"DESAFIOQ2BANK/api/models"
	"DESAFIOQ2BANK/api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BuscaUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	var dados models.Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao ler corpo da requisição")
		return
	}

	if err := json.Unmarshal(body, &dados); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao realizar unmarshal")
		return
	}
	pg, err := db.ConectarDB()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar com banco de dados")
		return
	}

	v, err := domains.BuscaUsuarioPorID(pg, dados.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, v)
}

func BuscaTodosUsuarios(w http.ResponseWriter, r *http.Request) {
	pg, err := db.ConectarDB()
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

func DeleteDados(w http.ResponseWriter, r *http.Request) {
	pg, err := db.ConectarDB()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	err = domains.DeleteUsuarioPorID(pg, r.FormValue("ID"))
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Usuário deletado com sucesso!")

}

func InsereUsuario(w http.ResponseWriter, r *http.Request) {
	var dados models.Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao ler o corpo da requisição")
		return
	}

	if err := json.Unmarshal(body, &dados); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao realizar unmarshal do cliente")
		return
	}

	pg, err := db.ConectarDB()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Erro ao conectar ao banco de dados")
		return
	}

	if len(dados.CPFCNPJ) == 11 {
		dados.Tipo = "comuns"
	}

	if len(dados.CPFCNPJ) == 14 {

		dados.Tipo = "lojista"
	}

	if len(dados.CPFCNPJ) != 14 && len(dados.CPFCNPJ) != 11 {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "CPF/CNPJ inválido!")
		return
	}

	err = domains.InsereUsuario(pg, dados)
	if err != nil {
		utils.RespondWithError(w, http.StatusOK, 0, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "CPF/CNPJ inserido com sucesso!")
}
