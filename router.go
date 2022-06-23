package main

import (
	"DESAFIOQ2BANK/api/db"
	"DESAFIOQ2BANK/api/domains/transport"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// config.Carregar()

	db.ConectarDB()

	router := mux.NewRouter()
	r := router.PathPrefix("/api").Subrouter()

	//ROTAS DE INTEGRAÇÃO COM BACK END

	r.HandleFunc("/buscausuarios", (transport.BuscaTodosUsuarios))
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
