package main

import (
	"DesafioQ2Bank/api/config"
	"DesafioQ2Bank/api/db"
	"DesafioQ2Bank/api/routes"
)

func main() {
	config.Carregar()
	db.ConectarDB()
	routes.Router()

}
