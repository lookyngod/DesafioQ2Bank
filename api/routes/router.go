package routes

import (
	"DesafioQ2Bank/api/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() {
	router := mux.NewRouter()
	// r := router.PathPrefix("/api").Subrouter()

	//ROTAS DE INTEGRAÇÃO DO FRONT END COM BACK END

	// r.HandleFunc("/table", (rest.table.CreateTable))
	http.Handle("/", router)

	log.Printf("Padronizando para porta %d", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil); err != nil {
		log.Fatal(err)
	}
}
