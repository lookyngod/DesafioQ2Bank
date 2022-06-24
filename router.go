package main

import (
	"DESAFIOQ2BANK/api/rest"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// config.Carregar()
	router := mux.NewRouter()
	r := router.PathPrefix("/api").Subrouter()

	//ROTAS DE INTEGRAÇÃO COM BACK END

	r.HandleFunc("/buscausuarios", (rest.UsuariosHandler))
	r.HandleFunc("/buscausuario/{id}", (rest.UsuarioHandler))

	r.HandleFunc("/buscatransacoes", (rest.TransacoesHandler))
	r.HandleFunc("/buscatransacao/{id}", (rest.TransacaoHandler))
	r.HandleFunc("/inserirtransacao", (rest.TransacaoHandler))

	http.Handle("/", router)

	log.Printf("Padronizando para porta %d", 5000)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 5000), nil); err != nil {
		log.Fatal(err)
	}

}

// 	router := mux.NewRouter()
// 	// r := router.PathPrefix("/api").Subrouter()

// 	//ROTAS DE INTEGRAÇÃO DO FRONT END COM BACK END

// 	// r.HandleFunc("/table", (rest.table.CreateTable))
// 	http.Handle("/", router)

// 	log.Printf("Padronizando para porta %d", config.Port)
// 	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
