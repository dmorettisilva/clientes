package main

import (
	"log"

	"github.com/dmorettisilva/clientes/config"
	"github.com/dmorettisilva/clientes/database"
	"github.com/dmorettisilva/clientes/services"
	_ "github.com/godror/godror"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("não pode iniciar servidor: %v", err)
	}

	dao, err := database.New(cfg.DBConfig)
	if err != nil {
		log.Fatalf("não pode criar DAO: %v", err)
	}

	defer dao.Close()

	log.Println("Connection with database has a successful!")
	log.Println("Web-service Alpha-Consulta-CNPJ - v1.2")

	service := services.New(cfg, dao)
	service.Listen()

}
