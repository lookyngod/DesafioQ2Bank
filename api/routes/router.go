package routes

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()

	router := mux.NewRouter()
	r := router.PathPrefix("/api").Subrouter()

	//ROTAS DE INTEGRAÇÃO DO FRONT END COM BACK END

	// r.HandleFunc("/handle", (rest.handle))
	http.Handle("/", router)

	log.Printf("Padronizando para porta %d", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil); err != nil {
		log.Fatal(err)
	}

}
